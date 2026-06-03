package wal

import (
	"encoding/binary"
	"io"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/model"
	"github.com/hanami/tidets/core/storageengine/utils/codec"
)

// WAL 记录格式（学习项目：只维护当前布局，无多版本回放分支）。
// 变更记录布局后，请删除 data-dir/wal.log。
const (
	magic uint32 = 0x54534457 // "TSDW"

	opInsert      uint8 = 1
	opInsertBatch uint8 = 2
)

// InsertFn 回放或测试时逐条应用写入。
type InsertFn func(key model.SeriesKey, p model.Point) error

func writeHeader(w io.Writer, op uint8) error {
	if err := binary.Write(w, binary.LittleEndian, magic); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, op)
}

func readHeader(r io.Reader) (uint8, error) {
	var m uint32
	if err := binary.Read(r, binary.LittleEndian, &m); err != nil {
		return 0, err
	}
	if m != magic {
		return 0, commons.ErrWALCorruptRecord
	}
	var op uint8
	if err := binary.Read(r, binary.LittleEndian, &op); err != nil {
		return 0, err
	}
	return op, nil
}

func writeInsert(w io.Writer, key model.SeriesKey, p model.Point) error {
	if err := writeHeader(w, opInsert); err != nil {
		return err
	}
	return writePointFields(w, key, p)
}

func writePointFields(w io.Writer, key model.SeriesKey, p model.Point) error {
	if err := codec.WriteString(w, key.DevicePath); err != nil {
		return err
	}
	if err := codec.WriteString(w, key.Measurement); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, p.Timestamp); err != nil {
		return err
	}
	return p.Value.WritePayload(w)
}

func readPointFields(r io.Reader) (model.SeriesKey, model.Point, error) {
	device, err := codec.ReadString(r)
	if err != nil {
		return model.SeriesKey{}, model.Point{}, err
	}
	measurement, err := codec.ReadString(r)
	if err != nil {
		return model.SeriesKey{}, model.Point{}, err
	}
	var ts int64
	if err := binary.Read(r, binary.LittleEndian, &ts); err != nil {
		return model.SeriesKey{}, model.Point{}, err
	}
	val, err := model.ReadValuePayload(r)
	if err != nil {
		return model.SeriesKey{}, model.Point{}, err
	}
	return model.SeriesKey{DevicePath: device, Measurement: measurement},
		model.Point{Timestamp: ts, Value: val}, nil
}

func writeInsertBatch(w io.Writer, records []model.BatchRecord) error {
	if len(records) == 0 {
		return nil
	}
	if err := writeHeader(w, opInsertBatch); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, uint32(len(records))); err != nil {
		return err
	}
	for _, rec := range records {
		if err := writePointFields(w, rec.Key, rec.Point); err != nil {
			return err
		}
	}
	return nil
}

func readInsertBatch(r io.Reader, apply InsertFn) error {
	var n uint32
	if err := binary.Read(r, binary.LittleEndian, &n); err != nil {
		return err
	}
	for i := uint32(0); i < n; i++ {
		key, p, err := readPointFields(r)
		if err != nil {
			return err
		}
		if err := apply(key, p); err != nil {
			return err
		}
	}
	return nil
}
