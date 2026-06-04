package storageengine

import (
	"sort"
	"sync"

	"github.com/hanami/tidets/core/tsmodel"
)

type timeRange struct {
	start int64
	end   int64
}

type tombstoneIndex struct {
	mu     sync.RWMutex
	log    *tombstoneLog
	ranges map[string][]timeRange
}

func newTombstoneIndex(log *tombstoneLog) *tombstoneIndex {
	return &tombstoneIndex{log: log, ranges: make(map[string][]timeRange)}
}

func (t *tombstoneIndex) Mark(keyStr string, start, end int64) error {
	if start > end {
		return nil
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.log != nil {
		if err := t.log.AppendMark(keyStr, start, end); err != nil {
			return err
		}
	}
	t.ranges[keyStr] = mergeTimeRanges(append(t.ranges[keyStr], timeRange{start: start, end: end}))
	return nil
}

func (t *tombstoneIndex) Restore(keyStr string, start, end int64) {
	if start > end {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	t.ranges[keyStr] = mergeTimeRanges(append(t.ranges[keyStr], timeRange{start: start, end: end}))
}

func (t *tombstoneIndex) Filter(keyStr string, pts []tsmodel.Point) []tsmodel.Point {
	if len(pts) == 0 {
		return pts
	}
	t.mu.RLock()
	ranges := t.ranges[keyStr]
	t.mu.RUnlock()
	if len(ranges) == 0 {
		return pts
	}
	out := make([]tsmodel.Point, 0, len(pts))
	for _, p := range pts {
		if !inTimeRanges(p.Timestamp, ranges) {
			out = append(out, p)
		}
	}
	return out
}

func (t *tombstoneIndex) IsEmpty() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.ranges) == 0
}

func (t *tombstoneIndex) Snapshot() map[string][]timeRange {
	t.mu.RLock()
	defer t.mu.RUnlock()

	out := make(map[string][]timeRange, len(t.ranges))
	for keyStr, ranges := range t.ranges {
		out[keyStr] = append([]timeRange(nil), ranges...)
	}
	return out
}

func (t *tombstoneIndex) Prune(drop map[string]map[timeRange]struct{}) error {
	if len(drop) == 0 {
		return nil
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	next := make(map[string][]timeRange, len(t.ranges))
	changed := false
	for keyStr, ranges := range t.ranges {
		drops := drop[keyStr]
		if len(drops) == 0 {
			next[keyStr] = append([]timeRange(nil), ranges...)
			continue
		}
		kept := make([]timeRange, 0, len(ranges))
		for _, r := range ranges {
			if _, ok := drops[r]; ok {
				changed = true
				continue
			}
			kept = append(kept, r)
		}
		if len(kept) > 0 {
			next[keyStr] = kept
		} else if len(ranges) > 0 {
			changed = true
		}
	}
	if !changed {
		return nil
	}
	if t.log != nil {
		if err := t.log.RewriteAll(next); err != nil {
			return err
		}
	}
	t.ranges = next
	return nil
}

func (t *tombstoneIndex) Sync() error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if t.log == nil {
		return nil
	}
	return t.log.Sync()
}

func (t *tombstoneIndex) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.log == nil {
		return nil
	}
	err := t.log.Close()
	t.log = nil
	return err
}

func mergeTimeRanges(ranges []timeRange) []timeRange {
	if len(ranges) == 0 {
		return nil
	}
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start == ranges[j].start {
			return ranges[i].end < ranges[j].end
		}
		return ranges[i].start < ranges[j].start
	})
	out := []timeRange{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &out[len(out)-1]
		if ranges[i].start <= last.end+1 {
			if ranges[i].end > last.end {
				last.end = ranges[i].end
			}
			continue
		}
		out = append(out, ranges[i])
	}
	return out
}

func inTimeRanges(ts int64, ranges []timeRange) bool {
	for _, r := range ranges {
		if ts < r.start {
			return false
		}
		if ts <= r.end {
			return true
		}
	}
	return false
}
