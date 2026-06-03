package mmap

// MapReadOnly 只读映射整个文件，返回映射区与释放函数。
func MapReadOnly(path string) (data []byte, unmap func() error, err error) {
	return mapReadOnly(path)
}
