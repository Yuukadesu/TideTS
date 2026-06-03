package wal

import (
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
)

const checkpointFilename = "wal.checkpoint"

var checkpointMagic = [4]byte{'T', 'S', 'C', 'P'}

// ReadCheckpoint 读取 WAL 检查点（可跳过的 wal.log 字节偏移）。
// 不存在时返回 0, nil。
func ReadCheckpoint(dir string) (int64, error) {
	path := filepath.Join(dir, checkpointFilename)
	b, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return 0, nil
		}
		return 0, err
	}
	if len(b) != 16 {
		return 0, nil
	}
	if string(b[:4]) != string(checkpointMagic[:]) {
		return 0, nil
	}
	off := int64(binary.LittleEndian.Uint64(b[8:16]))
	if off < 0 {
		return 0, nil
	}
	return off, nil
}

// WriteCheckpoint 原子写入 WAL 检查点（可跳过的 wal.log 字节偏移）。
// 注意：检查点只影响回放起点，不会自动截断 wal.log。
func WriteCheckpoint(dir string, offset int64) error {
	if offset < 0 {
		offset = 0
	}
	tmp := filepath.Join(dir, checkpointFilename+".tmp")
	path := filepath.Join(dir, checkpointFilename)

	buf := make([]byte, 16)
	copy(buf[:4], checkpointMagic[:])
	// 4 bytes reserved for future
	binary.LittleEndian.PutUint64(buf[8:16], uint64(offset))

	if err := os.WriteFile(tmp, buf, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
