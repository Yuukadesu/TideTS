package storageengine

import (
	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel"
)

// Delete 删除单时间戳测点。
func (e *Engine) Delete(key SeriesKey, ts int64) error {
	_, err := e.DeleteRange(key, ts, ts)
	return err
}

// DeleteRange 删除闭区间 [start, end] 内的测点，返回删除点数。
func (e *Engine) DeleteRange(key SeriesKey, start, end int64) (int, error) {
	if key.DevicePath == "" || key.Measurement == "" {
		return 0, commons.ErrStorageDeviceMeasurementRequired
	}
	if start > end {
		return 0, commons.ErrStorageInvalidRange
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if e.closing {
		return 0, commons.ErrStorageEngineClosing
	}

	before := e.countUnlocked(key, start, end)
	if before == 0 {
		return 0, nil
	}

	if start == end {
		if err := e.wal.AppendDelete(key, start); err != nil {
			return 0, err
		}
	} else {
		if err := e.wal.AppendDeleteRange(key, start, end); err != nil {
			return 0, err
		}
	}
	if err := e.applyDeleteLocked(key, start, end); err != nil {
		return 0, err
	}
	return before, nil
}

func (e *Engine) applyDeleteLocked(key SeriesKey, start, end int64) error {
	e.ws.DeleteRange(key, start, end)
	e.segments.DeleteRange(key, start, end)
	return e.tombstones.Mark(key.String(), start, end)
}

func (e *Engine) tombstoneFilter() func(string, []tsmodel.Point) []tsmodel.Point {
	return e.tombstones.Filter
}
