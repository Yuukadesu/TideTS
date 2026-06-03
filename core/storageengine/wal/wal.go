package wal

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/model"
)

// WAL 追加写日志（对齐 IoTDB：先 WAL 再写 MemTable）。
type WAL struct {
	file     *os.File
	w        *bufio.Writer
	syncMode SyncMode
}

func Open(dir string) (*WAL, error) {
	return OpenWithSync(dir, SyncAlways)
}

func OpenWithSync(dir string, mode SyncMode) (*WAL, error) {
	path := filepath.Join(dir, "wal.log")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, err
	}
	return &WAL{file: f, w: bufio.NewWriter(f), syncMode: mode}, nil
}

// Offset 返回 wal.log 当前字节偏移（用于写 checkpoint）。
// 注意：它会先 Flush bufio 缓冲，但不一定 fsync。
func (w *WAL) Offset() (int64, error) {
	if w == nil || w.file == nil {
		return 0, nil
	}
	if w.w != nil {
		if err := w.w.Flush(); err != nil {
			return 0, err
		}
	}
	fi, err := w.file.Stat()
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

// FileSize 返回 wal.log 字节数。
func FileSize(dir string) (int64, error) {
	fi, err := os.Stat(filepath.Join(dir, "wal.log"))
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}
	return fi.Size(), nil
}

func (w *WAL) AppendInsert(key model.SeriesKey, p model.Point) error {
	if err := writeInsert(w.w, key, p); err != nil {
		return err
	}
	return w.syncIfNeeded()
}

func (w *WAL) AppendInsertBatch(records []model.BatchRecord) error {
	if len(records) == 0 {
		return nil
	}
	if err := writeInsertBatch(w.w, records); err != nil {
		return err
	}
	return w.syncIfNeeded()
}

func (w *WAL) syncIfNeeded() error {
	if err := w.w.Flush(); err != nil {
		return err
	}
	if w.syncMode == SyncAlways {
		return w.file.Sync()
	}
	return nil
}

func (w *WAL) Close() error {
	if w.w != nil {
		if err := w.w.Flush(); err != nil {
			return err
		}
	}
	if w.file != nil {
		return w.file.Close()
	}
	return nil
}

// TruncateToZero 清空 wal.log 并将写指针回到起点。
// 这是 “轻量裁剪” 的最简实现：只在引擎确认没有需要回放的数据时调用。
func (w *WAL) TruncateToZero() error {
	if w == nil || w.file == nil {
		return nil
	}
	if w.w != nil {
		if err := w.w.Flush(); err != nil {
			return err
		}
	}
	if err := w.file.Truncate(0); err != nil {
		return err
	}
	if _, err := w.file.Seek(0, io.SeekStart); err != nil {
		return err
	}
	w.w = bufio.NewWriter(w.file)
	return nil
}

func Reset(dir string) (*WAL, error) {
	return ResetWithSync(dir, SyncAlways)
}

func ResetWithSync(dir string, mode SyncMode) (*WAL, error) {
	_ = os.Remove(filepath.Join(dir, "wal.log"))
	// wal.log 被重置后，旧 checkpoint 可能导致跳过新 WAL，因此同步清理。
	_ = os.Remove(filepath.Join(dir, checkpointFilename))
	return OpenWithSync(dir, mode)
}

// Sync 将缓冲刷盘（flush 完成后调用）。
func (w *WAL) Sync() error {
	if w.w != nil {
		if err := w.w.Flush(); err != nil {
			return err
		}
	}
	if w.file != nil {
		return w.file.Sync()
	}
	return nil
}

// Replay 将 WAL 记录回放给 apply（用于恢复 MemTable，并走与在线写入相同的路由）。
func Replay(dir string, apply InsertFn) error {
	if apply == nil {
		return commons.ErrWALApplyRequired
	}
	path := filepath.Join(dir, "wal.log")
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	var (
		r          *bufio.Reader
		firstOp    uint8
		hasFirstOp bool
	)

	if off, err := ReadCheckpoint(dir); err == nil && off > 0 {
		if _, err := f.Seek(off, io.SeekStart); err == nil {
			r = bufio.NewReader(f)
			// 探测 checkpoint 是否落在 record 边界；否则回退到全量回放。
			op, err := readHeader(r)
			if err == nil && (op == opInsert || op == opInsertBatch) {
				firstOp = op
				hasFirstOp = true
				goto startReplay
			}
		}
		_, _ = f.Seek(0, io.SeekStart)
	}

	r = bufio.NewReader(f)
startReplay:
	for {
		var (
			op  uint8
			err error
		)
		if hasFirstOp {
			op = firstOp
			hasFirstOp = false
		} else {
			op, err = readHeader(r)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}
		}
		switch op {
		case opInsert:
			key, p, err := readPointFields(r)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}
			if err := apply(key, p); err != nil {
				return err
			}
		case opInsertBatch:
			if err := readInsertBatch(r, apply); err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}
		default:
			return commons.ErrWALCorruptRecord
		}
	}
}
