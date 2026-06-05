package storageengine

import (
	"time"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/utils"
)

// Count 统计时间范围内测点数（存储层 pushdown：segment 只读时间戳列，不加载 value）。
func (e *Engine) Count(key SeriesKey, start, end int64) (int, error) {
	begin := time.Now()
	if start > end {
		return 0, commons.ErrStorageInvalidRange
	}

	e.mu.RLock()
	defer e.mu.RUnlock()
	n := e.countUnlocked(key, start, end)
	if e.hooks.OnRead != nil {
		e.hooks.OnRead("count", n, time.Since(begin))
	}
	return n, nil
}

func (e *Engine) countUnlocked(key SeriesKey, start, end int64) int {
	mem := e.queryMemLocked(key, start, end)
	disk := e.segments.QueryTimestamps(key, start, end)
	merged := utils.MergeSorted(mem, disk)
	merged = e.tombstones.Filter(key.String(), merged)
	return len(merged)
}
