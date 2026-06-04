package segment

import (
	"encoding/binary"
	"errors"
	"io"
	"os"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/storageengine/utils/codec"
	"github.com/hanami/tidets/core/tsmodel"
)

func openFileInMemory(path string) (*file, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var m, ver uint32
	if err := binary.Read(f, binary.LittleEndian, &m); err != nil {
		return nil, err
	}
	if m != magic {
		return nil, commons.ErrSegmentCorrupt
	}
	if err := binary.Read(f, binary.LittleEndian, &ver); err != nil {
		return nil, err
	}
	if ver != version {
		return nil, commons.ErrSegmentUnsupportedVersion(ver, version)
	}

	series := make(map[string][]tsmodel.Point)
	if err := readBody(f, series); err != nil {
		return nil, err
	}

	sf := &file{path: path, mem: series}
	sf.index = buildFileIndex(series)
	return sf, nil
}

// readBody 读取文件体：重复 [chunkCount → chunks]，以 chunkCount=0 + endMagic 结束。
func readBody(f *os.File, series map[string][]tsmodel.Point) error {
	for {
		var chunkCount uint32
		err := binary.Read(f, binary.LittleEndian, &chunkCount)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		if chunkCount == 0 {
			var em uint32
			if err := binary.Read(f, binary.LittleEndian, &em); err != nil {
				return err
			}
			if em == endMagic {
				return nil
			}
			return commons.ErrSegmentCorrupt
		}
		for i := uint32(0); i < chunkCount; i++ {
			if err := readChunk(f, series); err != nil {
				return err
			}
		}
	}
}

func readChunk(f *os.File, series map[string][]tsmodel.Point) error {
	device, err := codec.ReadString(f)
	if err != nil {
		return err
	}
	measurement, err := codec.ReadString(f)
	if err != nil {
		return err
	}
	var n uint32
	if err := binary.Read(f, binary.LittleEndian, &n); err != nil {
		return err
	}
	var minTs, maxTs int64
	if err := binary.Read(f, binary.LittleEndian, &minTs); err != nil {
		return err
	}
	if err := binary.Read(f, binary.LittleEndian, &maxTs); err != nil {
		return err
	}
	_ = minTs
	_ = maxTs

	var dt uint8
	if err := binary.Read(f, binary.LittleEndian, &dt); err != nil {
		return err
	}
	pts := make([]tsmodel.Point, n)
	for j := uint32(0); j < n; j++ {
		if err := binary.Read(f, binary.LittleEndian, &pts[j].Timestamp); err != nil {
			return err
		}
	}
	vals, err := tsmodel.ReadValuesColumn(f, tsmodel.DataType(dt), n)
	if err != nil {
		return err
	}
	for j := uint32(0); j < n; j++ {
		pts[j].Value = vals[j]
	}
	return mergeChunkSeries(series, device, measurement, pts)
}

func mergeChunkSeries(series map[string][]tsmodel.Point, device, measurement string, pts []tsmodel.Point) error {
	key := tsmodel.SeriesKey{DevicePath: device, Measurement: measurement}
	keyStr := key.String()
	if existing, ok := series[keyStr]; ok {
		series[keyStr] = utils.MergeSorted(existing, pts)
	} else {
		series[keyStr] = pts
	}
	return nil
}
