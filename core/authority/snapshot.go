package authority

import (
	"encoding/binary"
	"io"
	"os"
	"path/filepath"

	commons "github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel/codec"
)

const (
	snapshotMagic   uint32 = 0x54534153 // "TSAS"
	snapshotVersion uint32 = 1
)

const snapshotName = "auth.snapshot"

type snapshotState struct {
	Offset int64
	Users  []User
	Grants []Grant
}

func snapshotPath(dir string) string {
	return filepath.Join(dir, snapshotName)
}

func loadSnapshot(dir string) (*snapshotState, error) {
	data, err := os.ReadFile(snapshotPath(dir))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return decodeSnapshot(data)
}

func decodeSnapshot(data []byte) (*snapshotState, error) {
	if len(data) < 20 {
		return nil, commons.ErrAuthSnapshotCorrupt
	}
	if binary.LittleEndian.Uint32(data[0:4]) != snapshotMagic {
		return nil, commons.ErrAuthSnapshotCorrupt
	}
	ver := binary.LittleEndian.Uint32(data[4:8])
	if ver != snapshotVersion {
		return nil, commons.ErrAuthSnapshotUnsupported(ver, snapshotVersion)
	}
	st := &snapshotState{Offset: int64(binary.LittleEndian.Uint64(data[8:16]))}
	nUsers := binary.LittleEndian.Uint32(data[16:20])
	r := &byteReader{data: data, off: 20}
	st.Users = make([]User, 0, nUsers)
	for i := uint32(0); i < nUsers; i++ {
		u, err := readUserPayload(r)
		if err != nil {
			return nil, commons.ErrAuthSnapshotCorrupt
		}
		st.Users = append(st.Users, u)
	}
	var nGrants uint32
	if err := binary.Read(r, binary.LittleEndian, &nGrants); err != nil {
		return nil, commons.ErrAuthSnapshotCorrupt
	}
	st.Grants = make([]Grant, 0, nGrants)
	for i := uint32(0); i < nGrants; i++ {
		g, err := readGrantPayload(r)
		if err != nil {
			return nil, commons.ErrAuthSnapshotCorrupt
		}
		st.Grants = append(st.Grants, g)
	}
	if r.off != len(data) {
		return nil, commons.ErrAuthSnapshotCorrupt
	}
	if st.Offset < 0 {
		return nil, commons.ErrAuthSnapshotCorrupt
	}
	return st, nil
}

func persistSnapshot(dir string, payload []byte) error {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	path := snapshotPath(dir)
	tmp := path + ".tmp"
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	if _, err := f.Write(payload); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func encodeSnapshot(st snapshotState) ([]byte, error) {
	header := make([]byte, 20)
	binary.LittleEndian.PutUint32(header[0:4], snapshotMagic)
	binary.LittleEndian.PutUint32(header[4:8], snapshotVersion)
	binary.LittleEndian.PutUint64(header[8:16], uint64(st.Offset))
	binary.LittleEndian.PutUint32(header[16:20], uint32(len(st.Users)))
	w := &bytesWriter{buf: append([]byte{}, header...)}
	for _, u := range st.Users {
		if err := writeUserPayload(w, u); err != nil {
			return nil, err
		}
	}
	if err := binary.Write(w, binary.LittleEndian, uint32(len(st.Grants))); err != nil {
		return nil, err
	}
	for _, g := range st.Grants {
		if err := binary.Write(w, binary.LittleEndian, g.UserID); err != nil {
			return nil, err
		}
		if err := codec.WriteString(w, g.Path); err != nil {
			return nil, err
		}
		if err := binary.Write(w, binary.LittleEndian, uint8(g.Privilege)); err != nil {
			return nil, err
		}
	}
	return w.buf, nil
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

type bytesWriter struct {
	buf []byte
}

func (w *bytesWriter) Write(p []byte) (int, error) {
	w.buf = append(w.buf, p...)
	return len(p), nil
}
