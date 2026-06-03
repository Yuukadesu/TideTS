package tools

import (
	"fmt"
	"time"
)

// Stats 存储引擎运行指标。
type Stats struct {
	DataDir               string
	WALBytes              int64
	SegmentCount          int
	SealedSegmentCount    int
	ActiveSegmentBytes    int64
	MemTablePoints        int
	DelayedMemTablePoints int
	PendingFlushBatches   int
	LastFlushAt           time.Time
	AsyncFlushEnabled     bool
}

// String 便于日志输出。
func (s Stats) String() string {
	return fmt.Sprintf(
		"dir=%s wal=%dB segments=%d sealed=%d active=%dB mem=%d delayed=%d pending_flush=%d async=%v last_flush=%s",
		s.DataDir, s.WALBytes, s.SegmentCount, s.SealedSegmentCount, s.ActiveSegmentBytes,
		s.MemTablePoints, s.DelayedMemTablePoints, s.PendingFlushBatches,
		s.AsyncFlushEnabled, s.LastFlushAt.Format(time.RFC3339),
	)
}
