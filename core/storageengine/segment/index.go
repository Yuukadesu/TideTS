package segment

import "github.com/hanami/tidets/core/tsmodel"

type fileIndex struct {
	minTs     int64
	maxTs     int64
	hasSeries map[string]struct{}
}

func buildFileIndex(series map[string][]tsmodel.Point) fileIndex {
	idx := fileIndex{hasSeries: make(map[string]struct{}, len(series))}
	var initialized bool
	for keyStr, pts := range series {
		if len(pts) == 0 {
			continue
		}
		idx.hasSeries[keyStr] = struct{}{}
		if !initialized {
			idx.minTs = pts[0].Timestamp
			idx.maxTs = pts[len(pts)-1].Timestamp
			initialized = true
			continue
		}
		if pts[0].Timestamp < idx.minTs {
			idx.minTs = pts[0].Timestamp
		}
		if pts[len(pts)-1].Timestamp > idx.maxTs {
			idx.maxTs = pts[len(pts)-1].Timestamp
		}
	}
	return idx
}

func (idx fileIndex) canSkip(key tsmodel.SeriesKey, start, end int64) bool {
	if len(idx.hasSeries) == 0 {
		return true
	}
	if _, ok := idx.hasSeries[key.String()]; !ok {
		return true
	}
	if idx.maxTs != 0 && end < idx.minTs {
		return true
	}
	if start > idx.maxTs {
		return true
	}
	return false
}

func (mgr *Manager) rebuildFileIndex(sf *file) {
	if sf.mem != nil {
		sf.index = buildFileIndex(sf.mem)
		return
	}
	if sf.mapped != nil {
		sf.index = buildFileIndexFromChunks(sf.mapped.chunks)
	}
}
