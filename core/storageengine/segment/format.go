package segment

// .seg 磁盘格式（学习项目：只维护当前 version，不读旧文件）。
// 变更 magic、version 或 chunk 布局后，请删除 data-dir/segments 并重启。
const (
	magic   uint32 = 0x47455354 // "TSEG"
	version uint32 = 1          // 多批次 chunk + DataType 类型化值列

	endMagic uint32 = 0x53444E45 // "ENDS"

	SubDir                   = "segments"
	ActiveFileName           = "active.seg"
	DefaultFlushPoints       = 4096
	DefaultSealAfterFlushes  = 4
	DefaultCompactThreshold  = 4
	DefaultCompactMergeCount = 2
)
