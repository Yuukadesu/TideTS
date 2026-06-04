package segment

import (
	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/tsmodel"
)

// StableTimeByDevice 扫描已落盘 segment，重建每设备 stable time（对齐 IoTDB 启动恢复）。
func (mgr *Manager) StableTimeByDevice() map[string]int64 {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	stable := make(map[string]int64)
	scan := append([]*file{}, mgr.segments...)
	if mgr.activeMem != nil {
		scan = append(scan, mgr.activeMem)
	}
	for _, sf := range scan {
		sf.forEachSeriesMaxTs(func(keyStr string, maxTs int64) {
			device, _ := utils.SplitSeriesKey(keyStr)
			if maxTs > stable[device] {
				stable[device] = maxTs
			}
		})
	}
	return stable
}

// MaxTimestamp 返回某序列在磁盘上的最大时间戳（不存在则 ok=false）。
func (mgr *Manager) MaxTimestamp(key tsmodel.SeriesKey) (int64, bool) {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	var max int64
	found := false
	scan := append([]*file{}, mgr.segments...)
	if mgr.activeMem != nil {
		scan = append(scan, mgr.activeMem)
	}
	for _, sf := range scan {
		ts, ok := sf.maxTimestamp(key)
		if !ok {
			continue
		}
		if !found || ts > max {
			max = ts
			found = true
		}
	}
	return max, found
}
