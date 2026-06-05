package storageengine

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/flush"
	"github.com/hanami/tidets/core/storageengine/segment"
	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/storageengine/wal"
	"github.com/hanami/tidets/core/tsmodel"
)

// Engine 存储引擎：WAL + WorkingMemTable（normal/delayed）+ Segment（.seg）。
type Engine struct {
	dir        string
	flushAt    int
	asyncFlush bool
	walSync    wal.SyncMode
	walTrunc   bool
	hooks      Hooks
	schemaHas  func(SeriesKey) bool
	mu         sync.RWMutex
	ws         *workingSet
	wal        *wal.WAL
	segments   *segment.Manager
	tombstones *tombstoneIndex
	pending    []pendingFlush
	flushMgr   *flush.Manager
	lastFlush  time.Time
	closing    bool
}

type pendingFlush struct {
	snap      map[string][]tsmodel.Point
	walOffset int64
}

// Open 在 dataDir 下打开存储。
func Open(dataDir string) (*Engine, error) {
	return OpenWithOptions(Options{DataDir: dataDir})
}

// Options 引擎配置。
type Options struct {
	DataDir          string
	FlushAt          int
	AsyncFlush       *bool // nil = true
	WALSync          wal.SyncMode
	WALTruncate      bool // 空闲时将 wal.log truncate 到 0（不删除文件）
	SealAfterFlushes int  // active.seg 追加多少次后封存，0 = 默认
	CompactThreshold int  // 封存 .seg 数量达到后触发压缩，0 = 默认 4
	CompactMerge     int  // 每次合并最老文件数，0 = 默认 2
}

// OpenWithOptions 按配置打开引擎。
func OpenWithOptions(opts Options) (*Engine, error) {
	if opts.DataDir == "" {
		return nil, commons.ErrStorageDataDirRequired
	}
	flushAt := opts.FlushAt
	if flushAt <= 0 {
		flushAt = segment.DefaultFlushPoints
	}
	asyncFlush := true
	if opts.AsyncFlush != nil {
		asyncFlush = *opts.AsyncFlush
	}
	walSync := opts.WALSync
	if walSync != wal.SyncOnFlush && walSync != wal.SyncAlways {
		walSync = wal.SyncAlways
	}
	sealAfter := opts.SealAfterFlushes
	if sealAfter <= 0 {
		sealAfter = segment.DefaultSealAfterFlushes
	}

	if err := os.MkdirAll(opts.DataDir, 0o755); err != nil {
		return nil, err
	}

	segMgr, err := segment.OpenManagerWithCompact(opts.DataDir, sealAfter, segment.CompactOptions{
		Threshold:  opts.CompactThreshold,
		MergeCount: opts.CompactMerge,
	})
	if err != nil {
		return nil, err
	}
	stable := segMgr.StableTimeByDevice()
	ws := newWorkingSet(stable)
	tombstoneLog, err := openTombstoneLog(opts.DataDir, walSync)
	if err != nil {
		return nil, err
	}
	tombstones := newTombstoneIndex(tombstoneLog)
	if err := loadTombstoneLog(opts.DataDir, tombstones.Restore); err != nil {
		_ = tombstones.Close()
		return nil, err
	}

	if err := wal.ReplayWithOps(opts.DataDir, wal.ReplayOps{
		Insert: func(key tsmodel.SeriesKey, p tsmodel.Point) error {
			return ws.Insert(key, p)
		},
		Delete: func(key tsmodel.SeriesKey, ts int64) error {
			tombstones.Restore(key.String(), ts, ts)
			ws.Delete(key, ts)
			return nil
		},
		DeleteRange: func(key tsmodel.SeriesKey, start, end int64) error {
			tombstones.Restore(key.String(), start, end)
			ws.DeleteRange(key, start, end)
			return nil
		},
	}); err != nil {
		_ = tombstones.Close()
		return nil, err
	}

	w, err := wal.OpenWithSync(opts.DataDir, walSync)
	if err != nil {
		_ = tombstones.Close()
		return nil, err
	}

	e := &Engine{
		dir:        filepath.Clean(opts.DataDir),
		flushAt:    flushAt,
		asyncFlush: asyncFlush,
		walSync:    walSync,
		walTrunc:   opts.WALTruncate,
		ws:         ws,
		wal:        w,
		segments:   segMgr,
		tombstones: tombstones,
	}
	if asyncFlush {
		e.flushMgr = flush.NewManager(64)
	}
	return e, nil
}

func (e *Engine) Insert(key SeriesKey, p Point) error {
	return e.insertRecords("insert", []Record{{Key: key, Point: p}})
}

// BindSchemaGuard 绑定 schema 守卫；绑定后裸写必须先有 catalog。
func (e *Engine) BindSchemaGuard(has func(SeriesKey) bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.schemaHas = has
}

// Record 单条写入记录。
type Record struct {
	Key   SeriesKey
	Point Point
}

// InsertBatch 批量写入。
func (e *Engine) InsertBatch(records []Record) error {
	return e.insertRecords("insert_batch", records)
}

func (e *Engine) insertRecords(op string, records []Record) error {
	if len(records) == 0 {
		return nil
	}
	start := time.Now()
	batch := make([]tsmodel.BatchRecord, 0, len(records))
	for _, rec := range records {
		if err := utils.ValidatePoint(rec.Key, rec.Point); err != nil {
			return err
		}
		if e.schemaHas != nil && !e.schemaHas(rec.Key) {
			return commons.ErrStorageSchemaRequired
		}
		batch = append(batch, tsmodel.BatchRecord{Key: rec.Key, Point: rec.Point})
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if e.closing {
		return commons.ErrStorageEngineClosing
	}

	if err := e.wal.AppendInsertBatch(batch); err != nil {
		return err
	}
	if e.hooks.OnWAL != nil {
		e.hooks.OnWAL("append_"+op, len(batch))
	}
	for _, rec := range records {
		if err := e.ws.Insert(rec.Key, rec.Point); err != nil {
			return err
		}
	}
	if e.ws.PointCount() >= e.flushAt {
		if err := e.scheduleFlushWhileLocked(); err != nil {
			return err
		}
	}
	if e.hooks.OnWrite != nil {
		e.hooks.OnWrite(op, len(records), time.Since(start))
	}
	return nil
}

func (e *Engine) Flush() error {
	for {
		e.mu.Lock()
		if e.ws.IsEmpty() {
			pending := len(e.pending)
			e.mu.Unlock()
			if pending == 0 {
				return nil
			}
			e.WaitFlushed()
			continue
		}
		err := e.scheduleFlushWhileLocked()
		e.mu.Unlock()
		if err != nil {
			return err
		}
		e.WaitFlushed()
	}
}

// WaitFlushed 等待异步落盘队列清空。
func (e *Engine) WaitFlushed() {
	if e.flushMgr == nil {
		return
	}
	done := make(chan struct{})
	_ = e.flushMgr.Submit(func() { close(done) })
	<-done
}

func (e *Engine) scheduleFlushWhileLocked() error {
	if e.ws.IsEmpty() {
		return nil
	}
	snap := e.ws.Snapshot()
	walOffset := e.walEndOffsetLocked()
	e.ws.Reset()
	e.pending = append(e.pending, pendingFlush{snap: snap, walOffset: walOffset})

	if e.asyncFlush && e.flushMgr != nil {
		job := pendingFlush{snap: snap, walOffset: walOffset}
		return e.flushMgr.Submit(func() { _ = e.doFlushAndFinalize(job) })
	}
	e.mu.Unlock()
	err := e.doFlushAndFinalize(pendingFlush{snap: snap, walOffset: walOffset})
	e.mu.Lock()
	return err
}

// Compact 手动触发 segment 压缩（合并最老的若干 .seg）。
func (e *Engine) Compact() (err error) {
	start := time.Now()
	e.mu.Lock()
	defer e.mu.Unlock()
	stats, err := e.segments.CompactWithStats(e.tombstoneFilter())
	if err != nil {
		return err
	}
	if err := e.pruneTombstonesLocked(); err != nil {
		return err
	}
	if e.walSync == wal.SyncOnFlush && e.tombstones != nil {
		if err := e.tombstones.Sync(); err != nil {
			return err
		}
	}
	e.tryWALResetLocked()
	if e.hooks.OnCompact != nil && stats.InputFiles > 0 {
		e.hooks.OnCompact(time.Since(start), stats.InputFiles, stats.OutputFiles)
	}
	return nil
}

func (e *Engine) doFlushAndFinalize(job pendingFlush) error {
	start := time.Now()
	flushPoints := pointsInSeries(job.snap)
	compactStats, err := e.segments.FlushWithStats(job.snap)
	if err != nil {
		return err
	}
	if err := e.observeCompactLocked(start, compactStats); err != nil {
		return err
	}
	maybeCompactStats, err := e.segments.MaybeCompactWithStats(e.tombstoneFilter())
	if err != nil {
		return err
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.pending) > 0 {
		e.pending = e.pending[1:]
	}
	e.ws.applyFlush(job.snap)
	e.lastFlush = time.Now()
	if err := e.pruneTombstonesLocked(); err != nil {
		return err
	}
	// SyncOnFlush 模式：仅在 flush 成功后再 fsync WAL（更高吞吐，但 flush 边界提供更强持久性保证）。
	if e.walSync == wal.SyncOnFlush && e.wal != nil {
		if err := e.wal.Sync(); err != nil {
			return err
		}
	}
	if e.walSync == wal.SyncOnFlush && e.tombstones != nil {
		if err := e.tombstones.Sync(); err != nil {
			return err
		}
	}
	_ = wal.WriteCheckpoint(e.dir, e.checkpointOffsetLocked(job.walOffset))
	e.tryWALResetLocked()
	if e.hooks.OnFlush != nil {
		e.hooks.OnFlush(flushPoints, time.Since(start))
	}
	if e.hooks.OnCompact != nil && maybeCompactStats.InputFiles > 0 {
		e.hooks.OnCompact(time.Since(start), maybeCompactStats.InputFiles, maybeCompactStats.OutputFiles)
	}
	return nil
}

// walEndOffsetLocked 返回 wal.log 的“逻辑末尾偏移”（尽量在 flush 快照边界记录）。
// 失败时返回 0（表示不使用 checkpoint）。
func (e *Engine) walEndOffsetLocked() int64 {
	if e.wal == nil {
		return 0
	}
	off, err := e.wal.Offset()
	if err != nil {
		return 0
	}
	return off
}

func (e *Engine) checkpointOffsetLocked(offset int64) int64 {
	if e.tombstones != nil && !e.tombstones.IsEmpty() {
		return 0
	}
	return offset
}

func (e *Engine) pruneTombstonesLocked() error {
	if e.tombstones == nil || e.tombstones.IsEmpty() {
		return nil
	}
	snapshot := e.tombstones.Snapshot()
	drop := make(map[string]map[timeRange]struct{})
	dropCount := 0
	for keyStr, ranges := range snapshot {
		device, measurement := tsmodel.SplitSeriesKey(keyStr)
		if measurement == "" {
			continue
		}
		key := SeriesKey{DevicePath: device, Measurement: measurement}
		for _, r := range ranges {
			if e.hasRawPointsLocked(key, r.start, r.end) {
				continue
			}
			if drop[keyStr] == nil {
				drop[keyStr] = make(map[timeRange]struct{})
			}
			drop[keyStr][r] = struct{}{}
			dropCount++
		}
	}
	if err := e.tombstones.Prune(drop); err != nil {
		return err
	}
	if dropCount > 0 && e.hooks.OnTombstone != nil {
		e.hooks.OnTombstone("prune", dropCount)
	}
	return nil
}

func (e *Engine) hasRawPointsLocked(key SeriesKey, start, end int64) bool {
	mem := e.queryMemLocked(key, start, end)
	if len(mem) > 0 {
		return true
	}
	disk := e.segments.QueryTimestamps(key, start, end)
	return len(disk) > 0
}

func (e *Engine) tryWALResetLocked() {
	if len(e.pending) > 0 || e.ws.PointCount() > 0 {
		return
	}
	if e.tombstones != nil && !e.tombstones.IsEmpty() {
		return
	}
	if e.walSync == wal.SyncOnFlush && e.tombstones != nil {
		if err := e.tombstones.Sync(); err != nil {
			return
		}
	}
	// 空闲且所有 flush 已完成：此时 wal 里不应再有需要回放的数据。
	if e.walTrunc && e.wal != nil {
		if err := e.wal.TruncateToZero(); err == nil {
			_ = wal.WriteCheckpoint(e.dir, 0)
			if e.walSync == wal.SyncOnFlush {
				_ = e.wal.Sync()
			}
			if e.hooks.OnWAL != nil {
				e.hooks.OnWAL("truncate", 1)
			}
			return
		}
	}

	if e.wal != nil {
		_ = e.wal.Close()
	}
	w, err := wal.ResetWithSync(e.dir, e.walSync)
	if err != nil {
		return
	}
	e.wal = w
	if e.walSync == wal.SyncOnFlush {
		_ = e.wal.Sync()
	}
	if e.hooks.OnWAL != nil {
		e.hooks.OnWAL("reset", 1)
	}
}

func (e *Engine) DataDir() string { return e.dir }

func (e *Engine) Stats() Stats {
	e.mu.RLock()
	defer e.mu.RUnlock()

	st := Stats{
		DataDir:               e.dir,
		MemTablePoints:        e.ws.normal.PointCount(),
		DelayedMemTablePoints: e.ws.delayed.PointCount(),
		PendingFlushBatches:   len(e.pending),
		LastFlushAt:           e.lastFlush,
		AsyncFlushEnabled:     e.asyncFlush,
	}
	if e.segments != nil {
		st.SegmentCount = e.segments.FileCount()
		st.SealedSegmentCount = e.segments.SealedFileCount()
		st.ActiveSegmentBytes = e.segments.ActiveFileBytes()
	}
	if b, err := wal.FileSize(e.dir); err == nil {
		st.WALBytes = b
	}
	return st
}

func (e *Engine) StableTime(devicePath string) int64 {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.ws.stable[devicePath]
}

func (e *Engine) Close() error {
	if err := e.Flush(); err != nil {
		return err
	}
	if e.flushMgr != nil {
		e.flushMgr.Close()
	}
	e.mu.Lock()
	e.closing = true
	if err := e.segments.SealActive(); err != nil {
		e.mu.Unlock()
		return err
	}
	if e.wal != nil {
		_ = e.wal.Sync()
		_ = e.wal.Close()
		e.wal = nil
	}
	if e.tombstones != nil {
		_ = e.tombstones.Sync()
		_ = e.tombstones.Close()
		e.tombstones = nil
	}
	e.mu.Unlock()
	return nil
}

func (e *Engine) MemTablePointCount() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.ws.PointCount()
}

func (e *Engine) DelayedMemTablePointCount() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.ws.delayed.PointCount()
}

func (e *Engine) PendingFlushBatches() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return len(e.pending)
}

func (e *Engine) observeCompactLocked(start time.Time, stats segment.CompactStats) error {
	if e.hooks.OnCompact != nil && stats.InputFiles > 0 {
		e.hooks.OnCompact(time.Since(start), stats.InputFiles, stats.OutputFiles)
	}
	return nil
}

func pointsInSeries(series map[string][]tsmodel.Point) int {
	total := 0
	for _, pts := range series {
		total += len(pts)
	}
	return total
}
