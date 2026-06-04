package wal

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel"
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

func (w *WAL) AppendInsert(key tsmodel.SeriesKey, p tsmodel.Point) error {
	if err := writeInsert(w.w, key, p); err != nil {
		return err
	}
	return w.syncIfNeeded()
}

func (w *WAL) AppendInsertBatch(records []tsmodel.BatchRecord) error {
	if len(records) == 0 {
		return nil
	}
	if err := writeInsertBatch(w.w, records); err != nil {
		return err
	}
	return w.syncIfNeeded()
}

func (w *WAL) AppendDelete(key tsmodel.SeriesKey, ts int64) error {
	if err := writeDelete(w.w, key, ts); err != nil {
		return err
	}
	return w.syncIfNeeded()
}

func (w *WAL) AppendDeleteRange(key tsmodel.SeriesKey, start, end int64) error {
	if err := writeDeleteRange(w.w, key, start, end); err != nil {
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
	_ = os.Remove(filepath.Join(dir, checkpointFilename))
	return OpenWithSync(dir, mode)
}

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

// Replay 将 WAL 记录回放（兼容仅 Insert 的旧回调）。
func Replay(dir string, apply InsertFn) error {
	return ReplayWithOps(dir, ReplayOps{Insert: apply})
}

// ReplayWithOps 回放 WAL 写入与删除操作。
func ReplayWithOps(dir string, ops ReplayOps) error {
	if ops.Insert == nil {
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
			op, err := readHeader(r)
			if err == nil && isKnownOp(op) {
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
			if err := ops.Insert(key, p); err != nil {
				return err
			}
		case opInsertBatch:
			if err := readInsertBatch(r, ops.Insert); err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}
		case opDelete:
			key, ts, err := readDelete(r)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}
			if ops.Delete != nil {
				if err := ops.Delete(key, ts); err != nil {
					return err
				}
			}
		case opDeleteRange:
			key, start, end, err := readDeleteRange(r)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}
			if ops.DeleteRange != nil {
				if err := ops.DeleteRange(key, start, end); err != nil {
					return err
				}
			}
		default:
			return commons.ErrWALCorruptRecord
		}
	}
}
