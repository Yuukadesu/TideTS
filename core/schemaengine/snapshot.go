package schemaengine

import (
	"encoding/binary"
	"io"
	"os"
	"path/filepath"

	"github.com/hanami/tidets/commons/errors"
)

const (
	snapshotMagic   uint32 = 0x544D5353 // "TSMS"
	snapshotVersion uint32 = 1
)

const snapshotFileName = "mtree.snapshot"

type snapshotMeta struct {
	MlogOffset int64
	Series     []Timeseries
}

func schemaDir(dataDir string) string {
	return filepath.Join(dataDir, "system", "schema")
}

func snapshotPath(dir string) string {
	return filepath.Join(dir, snapshotFileName)
}

func loadSnapshot(dir string) (*snapshotMeta, error) {
	path := snapshotPath(dir)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return decodeSnapshot(data)
}

func decodeSnapshot(data []byte) (*snapshotMeta, error) {
	if len(data) < 20 {
		return nil, commons.ErrSchemaSnapshotCorrupt
	}
	if binary.LittleEndian.Uint32(data[0:4]) != snapshotMagic {
		return nil, commons.ErrSchemaSnapshotCorrupt
	}
	if binary.LittleEndian.Uint32(data[4:8]) != snapshotVersion {
		return nil, commons.ErrSchemaSnapshotUnsupported(binary.LittleEndian.Uint32(data[4:8]), snapshotVersion)
	}
	off := int64(binary.LittleEndian.Uint64(data[8:16]))
	count := binary.LittleEndian.Uint32(data[16:20])
	r := &byteReader{data: data, off: 20}
	series := make([]Timeseries, 0, count)
	for i := uint32(0); i < count; i++ {
		ts, err := readTimeseriesPayload(r)
		if err != nil {
			return nil, err
		}
		series = append(series, ts)
	}
	return &snapshotMeta{MlogOffset: off, Series: series}, nil
}

type byteReader struct {
	data []byte
	off  int
}

func (r *byteReader) Read(p []byte) (int, error) {
	if r.off >= len(r.data) {
		return 0, io.EOF
	}
	n := copy(p, r.data[r.off:])
	r.off += n
	return n, nil
}

func saveSnapshot(dir string, meta snapshotMeta) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	payload, err := encodeSnapshot(meta)
	if err != nil {
		return err
	}
	path := snapshotPath(dir)
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, payload, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func encodeSnapshot(meta snapshotMeta) ([]byte, error) {
	// header: magic(4) + version(4) + mlogOffset(8) + count(4) = 20
	buf := make([]byte, 0, 20+len(meta.Series)*64)
	header := make([]byte, 20)
	binary.LittleEndian.PutUint32(header[0:4], snapshotMagic)
	binary.LittleEndian.PutUint32(header[4:8], snapshotVersion)
	binary.LittleEndian.PutUint64(header[8:16], uint64(meta.MlogOffset))
	binary.LittleEndian.PutUint32(header[16:20], uint32(len(meta.Series)))
	buf = append(buf, header...)

	w := &bytesWriter{buf: buf}
	for _, ts := range meta.Series {
		if err := writeTimeseriesPayload(w, ts); err != nil {
			return nil, err
		}
	}
	return w.buf, nil
}

type bytesWriter struct {
	buf []byte
}

func (w *bytesWriter) Write(p []byte) (int, error) {
	w.buf = append(w.buf, p...)
	return len(p), nil
}
