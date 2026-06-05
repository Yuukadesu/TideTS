package segment

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/tsmodel"
)

// CompactOptions 压缩策略（对齐 IoTDB LSM 子集：合并多层小文件）。
type CompactOptions struct {
	// Threshold 封存文件数达到该值时尝试压缩；0 使用 DefaultCompactThreshold。
	Threshold int
	// MergeCount 每次合并最老的多少个 .seg；0 使用 DefaultCompactMergeCount。
	MergeCount int
}

// CompactStats 描述一次 compact 的输入输出文件数。
type CompactStats struct {
	InputFiles  int
	OutputFiles int
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
	return mgr.MaybeCompactWithFilter(nil)
}

func (mgr *Manager) MaybeCompactWithFilter(filter func(string, []tsmodel.Point) []tsmodel.Point) error {
	_, err := mgr.MaybeCompactWithStats(filter)
	return err
}

// MaybeCompactWithStats 在达到阈值时 compact，并返回输入输出文件统计。
func (mgr *Manager) MaybeCompactWithStats(filter func(string, []tsmodel.Point) []tsmodel.Point) (CompactStats, error) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	if len(mgr.segments) < mgr.compactThreshold {
		return CompactStats{}, nil
	}
	return mgr.compactLocked(filter)
}

// Compact 手动压缩：至少 2 个封存文件即合并最老的一批（忽略阈值）。
func (mgr *Manager) Compact() error {
	return mgr.CompactWithFilter(nil)
}

// CompactWithFilter 压缩并在合并时应用 tombstone 过滤。
func (mgr *Manager) CompactWithFilter(filter func(string, []tsmodel.Point) []tsmodel.Point) error {
	_, err := mgr.CompactWithStats(filter)
	return err
}

// CompactWithStats 手动压缩并返回输入输出文件统计。
func (mgr *Manager) CompactWithStats(filter func(string, []tsmodel.Point) []tsmodel.Point) (CompactStats, error) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	if len(mgr.segments) < 2 {
		return CompactStats{}, nil
	}
	return mgr.compactLocked(filter)
}

func (mgr *Manager) compactLocked(filter func(string, []tsmodel.Point) []tsmodel.Point) (CompactStats, error) {
	n := len(mgr.segments)
	if n < 2 {
		return CompactStats{}, nil
	}
	mergeN := mgr.compactMerge
	if mergeN > n {
		mergeN = n
	}

	toMerge := mgr.segments[:mergeN]
	merged := mergeFileSeries(toMerge, filter)

	for _, old := range toMerge {
		_ = old.close()
		_ = os.Remove(old.path)
	}
	rest := mgr.segments[mergeN:]
	if len(merged) == 0 {
		mgr.segments = rest
		return CompactStats{InputFiles: len(toMerge), OutputFiles: 0}, nil
	}

	mgr.nextID++
	path := filepath.Join(mgr.dir, fmt.Sprintf("%06d.seg", mgr.nextID))
	if err := writeFile(path, merged); err != nil {
		return CompactStats{}, commons.Wrap("segment", commons.CodeCorrupt, "compact write", err)
	}
	sf, err := openFile(path)
	if err != nil {
		_ = os.Remove(path)
		return CompactStats{}, commons.Wrap("segment", commons.CodeCorrupt, "compact open", err)
	}
	mgr.segments = append([]*file{sf}, rest...)
	return CompactStats{InputFiles: len(toMerge), OutputFiles: 1}, nil
}

func mergeFileSeries(files []*file, filter func(string, []tsmodel.Point) []tsmodel.Point) map[string][]tsmodel.Point {
	if len(files) == 0 {
		return nil
	}
	out := make(map[string][]tsmodel.Point)
	for _, sf := range files {
		for keyStr, pts := range sf.exportSeries() {
			if len(pts) == 0 {
				continue
			}
			if filter != nil {
				pts = filter(keyStr, pts)
			}
			if len(pts) == 0 {
				continue
			}
			if existing, ok := out[keyStr]; ok {
				out[keyStr] = utils.MergeSortedPreferNewer(existing, pts)
			} else {
				out[keyStr] = append([]tsmodel.Point(nil), pts...)
			}
		}
	}
	return out
}
