package segment

import (
	"bytes"
	"encoding/binary"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel"
)

const fileHeaderSize = 8 // magic + version

// chunkMeta 记录 mmap 文件中单个 chunk 的元数据与列偏移（不拷贝点数据）。
type chunkMeta struct {
	keyStr string
	n      uint32
	minTs  int64
	maxTs  int64
	dt     tsmodel.DataType
	tsOff  int
	valOff int
}

func scanChunks(data []byte) ([]chunkMeta, error) {
	if len(data) < fileHeaderSize {
		return nil, commons.ErrSegmentCorrupt
	}
	if binary.LittleEndian.Uint32(data[0:4]) != magic {
		return nil, commons.ErrSegmentCorrupt
	}
	if binary.LittleEndian.Uint32(data[4:8]) != version {
		return nil, commons.ErrSegmentUnsupportedVersion(binary.LittleEndian.Uint32(data[4:8]), version)
	}

	off := fileHeaderSize
	var out []chunkMeta
	for {
		if off+4 > len(data) {
			if off == len(data) {
				return out, nil
			}
			return nil, commons.ErrSegmentCorrupt
		}
		chunkCount := binary.LittleEndian.Uint32(data[off : off+4])
		off += 4
		if chunkCount == 0 {
			if off+4 > len(data) {
				return nil, commons.ErrSegmentCorrupt
			}
			if binary.LittleEndian.Uint32(data[off:off+4]) != endMagic {
				return nil, commons.ErrSegmentCorrupt
			}
			return out, nil
		}
		for i := uint32(0); i < chunkCount; i++ {
			meta, next, err := parseChunkMeta(data, off)
			if err != nil {
				return nil, err
			}
			out = append(out, meta)
			off = next
		}
	}
}

func parseChunkMeta(data []byte, off int) (chunkMeta, int, error) {
	device, off, err := readStringAt(data, off)
	if err != nil {
		return chunkMeta{}, off, err
	}
	measurement, off, err := readStringAt(data, off)
	if err != nil {
		return chunkMeta{}, off, err
	}
	if off+4 > len(data) {
		return chunkMeta{}, off, commons.ErrSegmentCorrupt
	}
	n := binary.LittleEndian.Uint32(data[off : off+4])
	off += 4
	if off+16 > len(data) {
		return chunkMeta{}, off, commons.ErrSegmentCorrupt
	}
	minTs := int64(binary.LittleEndian.Uint64(data[off : off+8]))
	off += 8
	maxTs := int64(binary.LittleEndian.Uint64(data[off : off+8]))
	off += 8
	if off+1 > len(data) {
		return chunkMeta{}, off, commons.ErrSegmentCorrupt
	}
	dt := tsmodel.DataType(data[off])
	off += 1

	tsOff := off
	tsBytes := int(n) * 8
	if tsOff+tsBytes > len(data) {
		return chunkMeta{}, off, commons.ErrSegmentCorrupt
	}
	valOff := tsOff + tsBytes
	valBytes, err := valueColumnByteSize(data, valOff, dt, n)
	if err != nil {
		return chunkMeta{}, off, err
	}
	if valOff+valBytes > len(data) {
		return chunkMeta{}, off, commons.ErrSegmentCorrupt
	}

	key := tsmodel.SeriesKey{DevicePath: device, Measurement: measurement}
	return chunkMeta{
		keyStr: key.String(),
		n:      n,
		minTs:  minTs,
		maxTs:  maxTs,
		dt:     dt,
		tsOff:  tsOff,
		valOff: valOff,
	}, valOff + valBytes, nil
}

func readStringAt(data []byte, off int) (string, int, error) {
	if off+2 > len(data) {
		return "", off, commons.ErrSegmentCorrupt
	}
	n := int(binary.LittleEndian.Uint16(data[off : off+2]))
	off += 2
	if off+n > len(data) {
		return "", off, commons.ErrSegmentCorrupt
	}
	s := string(data[off : off+n])
	off += n
	return s, off, nil
}

func fixedValueSize(dt tsmodel.DataType) int {
	switch dt {
	case tsmodel.DataTypeBoolean:
		return 1
	case tsmodel.DataTypeInt32, tsmodel.DataTypeFloat:
		return 4
	case tsmodel.DataTypeInt64, tsmodel.DataTypeDouble:
		return 8
	default:
		return 0
	}
}

func valueColumnByteSize(data []byte, off int, dt tsmodel.DataType, n uint32) (int, error) {
	if w := fixedValueSize(dt); w > 0 {
		return int(n) * w, nil
	}
	if dt != tsmodel.DataTypeText {
		return 0, commons.ErrStorageUnsupportedDataType(uint8(dt))
	}
	start := off
	for i := uint32(0); i < n; i++ {
		if off+2 > len(data) {
			return 0, commons.ErrSegmentCorrupt
		}
		ln := int(binary.LittleEndian.Uint16(data[off : off+2]))
		off += 2 + ln
		if off > len(data) {
			return 0, commons.ErrSegmentCorrupt
		}
	}
	return off - start, nil
}

func queryChunkTimestamps(data []byte, meta chunkMeta, start, end int64) ([]tsmodel.Point, error) {
	if meta.n == 0 || end < meta.minTs || start > meta.maxTs {
		return nil, nil
	}
	ts := readTimestampColumn(data, meta.tsOff, meta.n)
	lo, hi := rangeBounds(ts, start, end)
	if lo > hi {
		return nil, nil
	}
	count := int(hi - lo + 1)
	out := make([]tsmodel.Point, count)
	for i := 0; i < count; i++ {
		out[i].Timestamp = ts[lo+uint32(i)]
	}
	return out, nil
}

func queryChunk(data []byte, meta chunkMeta, start, end int64) ([]tsmodel.Point, error) {
	if meta.n == 0 || end < meta.minTs || start > meta.maxTs {
		return nil, nil
	}
	ts := readTimestampColumn(data, meta.tsOff, meta.n)
	lo, hi := rangeBounds(ts, start, end)
	if lo > hi {
		return nil, nil
	}
	count := hi - lo + 1
	vals, err := readValuesColumnSlice(data, meta.valOff, meta.dt, lo, count)
	if err != nil {
		return nil, err
	}
	out := make([]tsmodel.Point, count)
	for i := uint32(0); i < count; i++ {
		out[i] = tsmodel.Point{Timestamp: ts[lo+i], Value: vals[i]}
	}
	return out, nil
}

func readTimestampColumn(data []byte, off int, n uint32) []int64 {
	out := make([]int64, n)
	for i := uint32(0); i < n; i++ {
		out[i] = int64(binary.LittleEndian.Uint64(data[off : off+8]))
		off += 8
	}
	return out
}

func rangeBounds(ts []int64, start, end int64) (lo, hi uint32) {
	if len(ts) == 0 {
		return 0, ^uint32(0)
	}
	lo = uint32(searchGE(ts, start))
	if lo >= uint32(len(ts)) || ts[lo] > end {
		return 0, ^uint32(0)
	}
	hi = uint32(searchLE(ts, end))
	if hi < lo {
		return 0, ^uint32(0)
	}
	return lo, hi
}

func searchGE(ts []int64, v int64) int {
	lo, hi := 0, len(ts)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if ts[mid] < v {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func searchLE(ts []int64, v int64) int {
	lo, hi := 0, len(ts)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if ts[mid] <= v {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}

func readValuesColumnSlice(data []byte, off int, dt tsmodel.DataType, startIdx, count uint32) ([]tsmodel.Value, error) {
	if count == 0 {
		return nil, nil
	}
	if w := fixedValueSize(dt); w > 0 {
		off += int(startIdx) * w
		r := bytes.NewReader(data[off:])
		return tsmodel.ReadValuesColumn(r, dt, count)
	}
	if dt != tsmodel.DataTypeText {
		return nil, commons.ErrStorageUnsupportedDataType(uint8(dt))
	}
	for i := uint32(0); i < startIdx; i++ {
		_, next, err := readStringAt(data, off)
		if err != nil {
			return nil, err
		}
		off = next
	}
	r := bytes.NewReader(data[off:])
	return tsmodel.ReadValuesColumn(r, dt, count)
}

func buildFileIndexFromChunks(chunks []chunkMeta) fileIndex {
	idx := fileIndex{hasSeries: make(map[string]struct{}, len(chunks))}
	var initialized bool
	for _, c := range chunks {
		idx.hasSeries[c.keyStr] = struct{}{}
		if !initialized {
			idx.minTs = c.minTs
			idx.maxTs = c.maxTs
			initialized = true
			continue
		}
		if c.minTs < idx.minTs {
			idx.minTs = c.minTs
		}
		if c.maxTs > idx.maxTs {
			idx.maxTs = c.maxTs
		}
	}
	return idx
}
