package segment

import (
	"path/filepath"

	"github.com/hanami/tidets/core/storageengine/model"
	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/storageengine/utils/mmap"
)

type mappedFile struct {
	data   []byte
	unmap  func() error
	chunks []chunkMeta
}

// file 表示一个 segment 文件。封存文件用 mmap 按需读；active.seg 仍全量驻内存以便追加。
type file struct {
	path   string
	index  fileIndex
	mem    map[string][]model.Point
	mapped *mappedFile
}

func (sf *file) close() error {
	if sf.mapped == nil {
		return nil
	}
	err := sf.mapped.unmap()
	sf.mapped = nil
	return err
}

func openFile(path string) (*file, error) {
	if filepath.Base(path) == ActiveFileName {
		return openFileInMemory(path)
	}
	return openFileMmap(path)
}

func openFileMmap(path string) (*file, error) {
	data, unmap, err := mmap.MapReadOnly(path)
	if err != nil {
		return nil, err
	}
	chunks, err := scanChunks(data)
	if err != nil {
		_ = unmap()
		return nil, err
	}
	sf := &file{
		path: path,
		mapped: &mappedFile{
			data:   data,
			unmap:  unmap,
			chunks: chunks,
		},
	}
	sf.index = buildFileIndexFromChunks(chunks)
	return sf, nil
}

func (sf *file) exportSeries() map[string][]model.Point {
	if sf.mem != nil {
		return sf.mem
	}
	out := make(map[string][]model.Point)
	for _, c := range sf.mapped.chunks {
		pts, err := queryChunk(sf.mapped.data, c, c.minTs, c.maxTs)
		if err != nil || len(pts) == 0 {
			continue
		}
		if existing, ok := out[c.keyStr]; ok {
			out[c.keyStr] = utils.MergeSorted(existing, pts)
		} else {
			out[c.keyStr] = pts
		}
	}
	return out
}

func (sf *file) forEachSeriesMaxTs(fn func(keyStr string, maxTs int64)) {
	if sf.mem != nil {
		for keyStr, pts := range sf.mem {
			if len(pts) == 0 {
				continue
			}
			fn(keyStr, pts[len(pts)-1].Timestamp)
		}
		return
	}
	seen := make(map[string]int64, len(sf.mapped.chunks))
	for _, c := range sf.mapped.chunks {
		if prev, ok := seen[c.keyStr]; !ok || c.maxTs > prev {
			seen[c.keyStr] = c.maxTs
		}
	}
	for k, ts := range seen {
		fn(k, ts)
	}
}

func (sf *file) maxTimestamp(key model.SeriesKey) (int64, bool) {
	keyStr := key.String()
	if sf.mem != nil {
		pts := sf.mem[keyStr]
		if len(pts) == 0 {
			return 0, false
		}
		return pts[len(pts)-1].Timestamp, true
	}
	var max int64
	found := false
	for _, c := range sf.mapped.chunks {
		if c.keyStr != keyStr {
			continue
		}
		if !found || c.maxTs > max {
			max = c.maxTs
			found = true
		}
	}
	return max, found
}

func (sf *file) query(key model.SeriesKey, start, end int64) []model.Point {
	if sf.mem != nil {
		return utils.QueryPointsInRange(sf.mem[key.String()], start, end)
	}
	keyStr := key.String()
	var merged []model.Point
	for _, c := range sf.mapped.chunks {
		if c.keyStr != keyStr {
			continue
		}
		if end < c.minTs || start > c.maxTs {
			continue
		}
		part, err := queryChunk(sf.mapped.data, c, start, end)
		if err != nil || len(part) == 0 {
			continue
		}
		merged = utils.MergeSorted(merged, part)
	}
	return merged
}
