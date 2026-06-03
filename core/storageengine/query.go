package storageengine

import (
	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/utils"
)

// QueryOptions 范围查询参数。
type QueryOptions struct {
	Limit int // 0 表示不限制
}

// Query 按时间范围查询；limit 为 0 时不限制条数。
func (e *Engine) Query(key SeriesKey, start, end int64, limit int) ([]Point, error) {
	return e.QueryWithOptions(key, start, end, QueryOptions{Limit: limit})
}

func (e *Engine) QueryWithOptions(key SeriesKey, start, end int64, opts QueryOptions) ([]Point, error) {
	if start > end {
		return nil, commons.ErrStorageInvalidRange
	}

	e.mu.RLock()
	mem := e.queryMemLocked(key, start, end)
	disk := e.segments.Query(key, start, end)
	e.mu.RUnlock()

	merged := utils.MergeSorted(mem, disk)
	return utils.TruncatePoints(merged, opts.Limit), nil
}

func (e *Engine) queryMemLocked(key SeriesKey, start, end int64) []Point {
	mem := e.ws.Query(key, start, end)
	for _, snap := range e.pending {
		part := utils.QueryPointsInRange(snap.snap[key.String()], start, end)
		if len(part) > 0 {
			mem = utils.MergeSorted(mem, part)
		}
	}
	return mem
}
