package segment

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/hanami/tidets/core/storageengine/model"
	"github.com/hanami/tidets/core/storageengine/utils"
)

// Manager 管理 dataDir/segments：封存 .seg + 可追加 active.seg（阶段 3）。
type Manager struct {
	dir              string
	mu               sync.RWMutex
	segments         []*file
	activeMem        *file
	activePath       string
	activeFlushes    int
	sealAfter        int
	compactThreshold int
	compactMerge     int
	nextID           uint64
}

func OpenManager(dataDir string, sealAfter int) (*Manager, error) {
	return OpenManagerWithCompact(dataDir, sealAfter, CompactOptions{})
}

func OpenManagerWithCompact(dataDir string, sealAfter int, compact CompactOptions) (*Manager, error) {
	if sealAfter <= 0 {
		sealAfter = DefaultSealAfterFlushes
	}
	compact = compact.withDefaults()
	segDir := filepath.Join(dataDir, SubDir)
	if err := os.MkdirAll(segDir, 0o755); err != nil {
		return nil, err
	}

	mgr := &Manager{
		dir:              segDir,
		sealAfter:        sealAfter,
		compactThreshold: compact.Threshold,
		compactMerge:     compact.MergeCount,
		activePath:       filepath.Join(segDir, ActiveFileName),
	}

	entries, err := os.ReadDir(segDir)
	if err != nil {
		return nil, err
	}

	var ids []uint64
	for _, ent := range entries {
		if ent.IsDir() {
			continue
		}
		name := ent.Name()
		if name == ActiveFileName {
			sf, err := openFile(filepath.Join(segDir, name))
			if err != nil {
				return nil, fmt.Errorf("segment: open active: %w", err)
			}
			mgr.activeMem = sf
			mgr.rebuildFileIndex(sf)
			continue
		}
		if !strings.HasSuffix(name, ".seg") {
			continue
		}
		id, err := strconv.ParseUint(strings.TrimSuffix(name, ".seg"), 10, 64)
		if err != nil {
			continue
		}
		sf, err := openFile(filepath.Join(segDir, name))
		if err != nil {
			return nil, fmt.Errorf("segment: open %s: %w", name, err)
		}
		mgr.segments = append(mgr.segments, sf)
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	if len(ids) > 0 {
		mgr.nextID = ids[len(ids)-1]
	}
	return mgr, nil
}

func (mgr *Manager) Flush(series map[string][]model.Point) error {
	if len(series) == 0 {
		return nil
	}
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	if err := mgr.appendToActive(series); err != nil {
		return err
	}
	if mgr.activeFlushes >= mgr.sealAfter {
		if err := mgr.sealActiveLocked(); err != nil {
			return err
		}
		if len(mgr.segments) >= mgr.compactThreshold {
			return mgr.compactLocked()
		}
	}
	return nil
}

func (mgr *Manager) Query(key model.SeriesKey, start, end int64) []model.Point {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	var merged []model.Point
	for _, sf := range mgr.segments {
		if sf.index.canSkip(key, start, end) {
			continue
		}
		part := sf.query(key, start, end)
		if len(part) > 0 {
			merged = utils.MergeSorted(merged, part)
		}
	}
	if mgr.activeMem != nil && !mgr.activeMem.index.canSkip(key, start, end) {
		part := mgr.activeMem.query(key, start, end)
		if len(part) > 0 {
			merged = utils.MergeSorted(merged, part)
		}
	}
	return merged
}

func (mgr *Manager) FileCount() int {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	n := len(mgr.segments)
	if mgr.activeMem != nil {
		n++
	}
	return n
}

// SealedFileCount 仅统计已封存的 .seg 数量（不含 active.seg）。
func (mgr *Manager) SealedFileCount() int {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	return len(mgr.segments)
}

func (mgr *Manager) ActiveFileBytes() int64 {
	fi, err := os.Stat(mgr.activePath)
	if err != nil {
		return 0
	}
	return fi.Size()
}
