package segment

import (
	"encoding/binary"
	"os"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/model"
	"github.com/hanami/tidets/core/storageengine/utils"
	"github.com/hanami/tidets/core/storageengine/utils/codec"
)

func writeFile(path string, series map[string][]model.Point) error {
	if len(series) == 0 {
		return commons.ErrSegmentNothingToFlush
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
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

	var chunkCount uint32
	for _, pts := range series {
		if len(pts) > 0 {
			chunkCount++
		}
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
	// 批次结束：chunkCount=0 + endMagic（与 active.seg 追加格式一致）
	if err := binary.Write(f, binary.LittleEndian, uint32(0)); err != nil {
		return err
	}
	if err := binary.Write(f, binary.LittleEndian, endMagic); err != nil {
		return err
	}
	return f.Sync()
}

func writeChunk(f *os.File, device, measurement string, pts []model.Point) error {
	if err := codec.WriteString(f, device); err != nil {
		return err
	}
	if err := codec.WriteString(f, measurement); err != nil {
		return err
	}
	n := uint32(len(pts))
	if err := binary.Write(f, binary.LittleEndian, n); err != nil {
		return err
	}
	if err := binary.Write(f, binary.LittleEndian, pts[0].Timestamp); err != nil {
		return err
	}
	if err := binary.Write(f, binary.LittleEndian, pts[len(pts)-1].Timestamp); err != nil {
		return err
	}
	dt := pts[0].Value.Type
	if err := binary.Write(f, binary.LittleEndian, uint8(dt)); err != nil {
		return err
	}
	for _, p := range pts {
		if err := binary.Write(f, binary.LittleEndian, p.Timestamp); err != nil {
			return err
		}
	}
	return model.WriteValuesColumn(f, dt, pts)
}
