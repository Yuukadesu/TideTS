//go:build !windows

package mmap

import (
	"os"

	"golang.org/x/sys/unix"
)

func mapReadOnly(path string) ([]byte, func() error, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		_ = f.Close()
		return nil, nil, err
	}
	size := fi.Size()
	if size == 0 {
		_ = f.Close()
		return nil, func() error { return nil }, nil
	}
	if size > int64(^uint(0)>>1) {
		_ = f.Close()
		return nil, nil, unix.EINVAL
	}
	data, err := unix.Mmap(int(f.Fd()), 0, int(size), unix.PROT_READ, unix.MAP_SHARED)
	if err != nil {
		_ = f.Close()
		return nil, nil, err
	}
	unmap := func() error {
		if err := unix.Munmap(data); err != nil {
			return err
		}
		return f.Close()
	}
	return data, unmap, nil
}
