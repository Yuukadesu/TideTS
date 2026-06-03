package segment

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/model"
	"github.com/hanami/tidets/core/storageengine/utils"
)

// CompactOptions 压缩策略（对齐 IoTDB LSM 子集：合并多层小文件）。
type CompactOptions struct {
	// Threshold 封存文件数达到该值时尝试压缩；0 使用 DefaultCompactThreshold。
	Threshold int
	// MergeCount 每次合并最老的多少个 .seg；0 使用 DefaultCompactMergeCount。
	MergeCount int
}

func (o CompactOptions) withDefaults() CompactOptions {
	if o.Threshold <= 0 {
		o.Threshold = DefaultCompactThreshold
	}
	if o.MergeCount <= 0 {
		o.MergeCount = DefaultCompactMergeCount
	}
	if o.MergeCount > o.Threshold {
		o.MergeCount = o.Threshold
	}
	return o
}

// MaybeCompact 封存文件数达到阈值时合并最老的一批。
func (mgr *Manager) MaybeCompact() error {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	if len(mgr.segments) < mgr.compactThreshold {
		return nil
	}
	return mgr.compactLocked()
}

// Compact 手动压缩：至少 2 个封存文件即合并最老的一批（忽略阈值）。
func (mgr *Manager) Compact() error {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	if len(mgr.segments) < 2 {
		return nil
	}
	return mgr.compactLocked()
}

func (mgr *Manager) compactLocked() error {
	n := len(mgr.segments)
	if n < 2 {
		return nil
	}
	mergeN := mgr.compactMerge
	if mergeN > n {
		mergeN = n
	}

	toMerge := mgr.segments[:mergeN]
	merged := mergeFileSeries(toMerge)

	mgr.nextID++
	path := filepath.Join(mgr.dir, fmt.Sprintf("%06d.seg", mgr.nextID))
	if err := writeFile(path, merged); err != nil {
		return commons.Wrap("segment", commons.CodeCorrupt, "compact write", err)
	}
	sf, err := openFile(path)
	if err != nil {
		_ = os.Remove(path)
		return commons.Wrap("segment", commons.CodeCorrupt, "compact open", err)
	}

	for _, old := range toMerge {
		_ = old.close()
		_ = os.Remove(old.path)
	}
	rest := mgr.segments[mergeN:]
	mgr.segments = append([]*file{sf}, rest...)
	return nil
}

func mergeFileSeries(files []*file) map[string][]model.Point {
	if len(files) == 0 {
		return nil
	}
	out := make(map[string][]model.Point)
	for _, sf := range files {
		for keyStr, pts := range sf.exportSeries() {
			if len(pts) == 0 {
				continue
			}
			if existing, ok := out[keyStr]; ok {
				out[keyStr] = utils.MergeSortedPreferNewer(existing, pts)
			} else {
				out[keyStr] = append([]model.Point(nil), pts...)
			}
		}
	}
	return out
}
