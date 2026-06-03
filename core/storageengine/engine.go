package storageengine

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/flush"
	"github.com/hanami/tidets/core/storageengine/model"
	"github.com/hanami/tidets/core/storageengine/segment"
	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/storageengine/wal"
)

// Engine 存储引擎：WAL + WorkingMemTable（normal/delayed）+ Segment（.seg）。
type Engine struct {
	dir        string
	flushAt    int
	asyncFlush bool
	walSync    wal.SyncMode
	walTrunc   bool
	mu         sync.RWMutex
	ws         *workingSet
	wal        *wal.WAL
	segments   *segment.Manager
	pending    []pendingFlush
	flushMgr   *flush.Manager
	lastFlush  time.Time
	closing    bool
}

type pendingFlush struct {
	snap      map[string][]model.Point
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

	if err := wal.Replay(opts.DataDir, func(key model.SeriesKey, p model.Point) error {
		return ws.Insert(key, p)
	}); err != nil {
		return nil, err
	}

	w, err := wal.OpenWithSync(opts.DataDir, walSync)
	if err != nil {
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
	}
	if asyncFlush {
		e.flushMgr = flush.NewManager(64)
	}
	return e, nil
}

func (e *Engine) Insert(key SeriesKey, p Point) error {
	return e.InsertBatch([]Record{{Key: key, Point: p}})
}

// Record 单条写入记录。
type Record struct {
	Key   SeriesKey
	Point Point
}

// InsertBatch 批量写入。
func (e *Engine) InsertBatch(records []Record) error {
	if len(records) == 0 {
		return nil
	}
	batch := make([]model.BatchRecord, 0, len(records))
	for _, rec := range records {
		if err := utils.ValidatePoint(rec.Key, rec.Point); err != nil {
			return err
		}
		batch = append(batch, model.BatchRecord{Key: rec.Key, Point: rec.Point})
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if e.closing {
		return commons.ErrStorageEngineClosing
	}

	if err := e.wal.AppendInsertBatch(batch); err != nil {
		return err
	}
	for _, rec := range records {
		if err := e.ws.Insert(rec.Key, rec.Point); err != nil {
			return err
		}
	}
	if e.ws.PointCount() >= e.flushAt {
		return e.scheduleFlushWhileLocked()
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
func (e *Engine) Compact() error {
	return e.segments.Compact()
}

func (e *Engine) doFlushAndFinalize(job pendingFlush) error {
	if err := e.segments.Flush(job.snap); err != nil {
		return err
	}
	// Flush 内 seal 后可能已 MaybeCompact；此处再尝试一次（仅 append 未 seal 时）
	if err := e.segments.MaybeCompact(); err != nil {
		return err
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.pending) > 0 {
		e.pending = e.pending[1:]
	}
	e.ws.applyFlush(job.snap)
	e.lastFlush = time.Now()
	// SyncOnFlush 模式：仅在 flush 成功后再 fsync WAL（更高吞吐，但 flush 边界提供更强持久性保证）。
	if e.walSync == wal.SyncOnFlush && e.wal != nil {
		if err := e.wal.Sync(); err != nil {
			return err
		}
	}
	_ = wal.WriteCheckpoint(e.dir, job.walOffset)
	e.tryWALResetLocked()
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

func (e *Engine) tryWALResetLocked() {
	if len(e.pending) > 0 || e.ws.PointCount() > 0 {
		return
	}
	// 空闲且所有 flush 已完成：此时 wal 里不应再有需要回放的数据。
	if e.walTrunc && e.wal != nil {
		if err := e.wal.TruncateToZero(); err == nil {
			_ = wal.WriteCheckpoint(e.dir, 0)
			if e.walSync == wal.SyncOnFlush {
				_ = e.wal.Sync()
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
