package schemaengine

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
	"path/filepath"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel"
	"github.com/hanami/tidets/core/tsmodel/codec"
)

const (
	mlogMagic uint32 = 0x544D4C53 // "TSML"

	opCreateTimeseries uint8 = 1
	opBatchCreate      uint8 = 2
)

const mlogFileName = "mlog.bin"

type mlog struct {
	file *os.File
	w    *bufio.Writer
}

func mlogPath(dir string) string {
	return filepath.Join(dir, mlogFileName)
}

func openMlog(dir string) (*mlog, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	path := mlogPath(dir)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, err
	}
	return &mlog{file: f, w: bufio.NewWriter(f)}, nil
}

func (m *mlog) offset() (int64, error) {
	if m == nil || m.file == nil {
		return 0, nil
	}
	if err := m.w.Flush(); err != nil {
		return 0, err
	}
	fi, err := m.file.Stat()
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func (m *mlog) appendCreate(ts Timeseries) error {
	if err := writeHeader(m.w, opCreateTimeseries); err != nil {
		return err
	}
	if err := writeTimeseriesPayload(m.w, ts); err != nil {
		return err
	}
	return m.sync()
}

func (m *mlog) appendBatch(series []Timeseries) error {
	if len(series) == 0 {
		return nil
	}
	if err := writeHeader(m.w, opBatchCreate); err != nil {
		return err
	}
	if err := binary.Write(m.w, binary.LittleEndian, uint32(len(series))); err != nil {
		return err
	}
	for _, ts := range series {
		if err := writeTimeseriesPayload(m.w, ts); err != nil {
			return err
		}
	}
	return m.sync()
}

func (m *mlog) sync() error {
	if err := m.w.Flush(); err != nil {
		return err
	}
	return m.file.Sync()
}

func (m *mlog) close() error {
	if m == nil {
		return nil
	}
	if m.w != nil {
		if err := m.w.Flush(); err != nil {
			return err
		}
	}
	if m.file != nil {
		return m.file.Close()
	}
	return nil
}

func writeHeader(w io.Writer, op uint8) error {
	if err := binary.Write(w, binary.LittleEndian, mlogMagic); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, op)
}

func readHeader(r io.Reader) (uint8, error) {
	var magic uint32
	if err := binary.Read(r, binary.LittleEndian, &magic); err != nil {
		return 0, err
	}
	if magic != mlogMagic {
		return 0, commons.ErrSchemaMlogCorrupt
	}
	var op uint8
	if err := binary.Read(r, binary.LittleEndian, &op); err != nil {
		return 0, err
	}
	return op, nil
}

func writeTimeseriesPayload(w io.Writer, ts Timeseries) error {
	if err := codec.WriteString(w, ts.DevicePath); err != nil {
		return err
	}
	if err := codec.WriteString(w, ts.Measurement); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, uint8(ts.DataType)); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, uint8(ts.Encoding)); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, uint8(ts.Compressor)); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, ts.CreatedAt)
}

func readTimeseriesPayload(r io.Reader) (Timeseries, error) {
	device, err := codec.ReadString(r)
	if err != nil {
		return Timeseries{}, err
	}
	measurement, err := codec.ReadString(r)
	if err != nil {
		return Timeseries{}, err
	}
	var dt, enc, comp uint8
	if err := binary.Read(r, binary.LittleEndian, &dt); err != nil {
		return Timeseries{}, err
	}
	if err := binary.Read(r, binary.LittleEndian, &enc); err != nil {
		return Timeseries{}, err
	}
	if err := binary.Read(r, binary.LittleEndian, &comp); err != nil {
		return Timeseries{}, err
	}
	var created int64
	if err := binary.Read(r, binary.LittleEndian, &created); err != nil {
		return Timeseries{}, err
	}
	return Timeseries{
		DevicePath:  device,
		Measurement: measurement,
		DataType:    tsmodel.DataType(dt),
		Encoding:    Encoding(enc),
		Compressor:  Compressor(comp),
		CreatedAt:   created,
	}, nil
}

type applyFn func(ts Timeseries) error

func replayMlog(dir string, fromOffset int64, apply applyFn) error {
	path := mlogPath(dir)
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	if fromOffset > 0 {
		if _, err := f.Seek(fromOffset, io.SeekStart); err != nil {
			return err
		}
	}

	r := bufio.NewReader(f)
	for {
		op, err := readHeader(r)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch op {
		case opCreateTimeseries:
			ts, err := readTimeseriesPayload(r)
			if err != nil {
				return err
			}
			if err := apply(ts); err != nil {
				return err
			}
		case opBatchCreate:
			var count uint32
			if err := binary.Read(r, binary.LittleEndian, &count); err != nil {
				return err
			}
			for i := uint32(0); i < count; i++ {
				ts, err := readTimeseriesPayload(r)
				if err != nil {
					return err
				}
				if err := apply(ts); err != nil {
					return err
				}
			}
		default:
			return commons.ErrSchemaMlogCorrupt
		}
	}
}
