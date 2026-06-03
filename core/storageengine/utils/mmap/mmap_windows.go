//go:build windows

package mmap

import (
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
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
	h, err := windows.CreateFileMapping(windows.Handle(f.Fd()), nil, windows.PAGE_READONLY, 0, 0, nil)
	if err != nil {
		_ = f.Close()
		return nil, nil, err
	}
	addr, err := windows.MapViewOfFile(h, windows.FILE_MAP_READ, 0, 0, 0)
	if err != nil {
		_ = windows.CloseHandle(h)
		_ = f.Close()
		return nil, nil, err
	}
	data := unsafe.Slice((*byte)(unsafe.Pointer(addr)), int(size))
	unmap := func() error {
		if err := windows.UnmapViewOfFile(addr); err != nil {
			return err
		}
		if err := windows.CloseHandle(h); err != nil {
			return err
		}
		return f.Close()
	}
	return data, unmap, nil
}
