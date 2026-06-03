package wal

// SyncMode WAL 刷盘策略（阶段 2）。
type SyncMode int

const (
	// SyncAlways 每条 batch 后 fsync（默认，最安全）。
	SyncAlways SyncMode = iota
	// SyncOnFlush 仅在 MemTable flush 完成后 fsync WAL（更高吞吐，依赖 flush 落盘）。
	SyncOnFlush
)
