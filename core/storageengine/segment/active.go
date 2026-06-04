package segment

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/tsmodel"
)

func (mgr *Manager) ensureActiveHeader() error {
	if _, err := os.Stat(mgr.activePath); err == nil {
		return nil
	}
	f, err := os.OpenFile(mgr.activePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := binary.Write(f, binary.LittleEndian, magic); err != nil {
		return err
	}
	if err := binary.Write(f, binary.LittleEndian, version); err != nil {
		return err
	}
	return f.Sync()
}

func (mgr *Manager) appendToActive(series map[string][]tsmodel.Point) error {
	if err := mgr.ensureActiveHeader(); err != nil {
		return err
	}
	f, err := os.OpenFile(mgr.activePath, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	var chunkCount uint32
	for _, pts := range series {
		if len(pts) > 0 {
			chunkCount++
		}
	}
	if chunkCount == 0 {
		return nil
	}
	if err := binary.Write(f, binary.LittleEndian, chunkCount); err != nil {
		return err
	}
	for keyStr, pts := range series {
		if len(pts) == 0 {
			continue
		}
		device, measurement := utils.SplitSeriesKey(keyStr)
		if err := writeChunk(f, device, measurement, pts); err != nil {
			return err
		}
	}
	if err := f.Sync(); err != nil {
		return err
	}

	mgr.mergeIntoActiveMem(series)
	mgr.activeFlushes++
	return nil
}

func (mgr *Manager) mergeIntoActiveMem(series map[string][]tsmodel.Point) {
	if mgr.activeMem == nil {
		mgr.activeMem = &file{path: mgr.activePath, mem: make(map[string][]tsmodel.Point)}
	}
	for keyStr, pts := range series {
		existing := mgr.activeMem.mem[keyStr]
		mgr.activeMem.mem[keyStr] = utils.MergeSorted(existing, pts)
	}
	mgr.rebuildFileIndex(mgr.activeMem)
}

// SealActive 封存 active.seg 为编号 .seg 文件（关闭引擎时调用）。
func (mgr *Manager) SealActive() error {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	return mgr.sealActiveLocked()
}

func (mgr *Manager) sealActiveLocked() error {
	if mgr.activeMem == nil {
		return nil
	}
	series := mgr.activeMem.exportSeries()
	if len(series) > 0 {
		mgr.nextID++
		name := filepath.Join(mgr.dir, fmt.Sprintf("%06d.seg", mgr.nextID))
		if err := writeFile(name, series); err != nil {
			return err
		}
		sf, err := openFileMmap(name)
		if err != nil {
			return err
		}
		mgr.segments = append(mgr.segments, sf)
	}
	if err := os.Remove(mgr.activePath); err != nil && !os.IsNotExist(err) {
		return err
	}
	mgr.activeMem = nil
	mgr.activePath = filepath.Join(mgr.dir, ActiveFileName)
	mgr.activeFlushes = 0
	return nil
}
