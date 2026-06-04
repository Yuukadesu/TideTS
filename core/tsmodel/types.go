package tsmodel

import "fmt"

// SeriesKey 标识一条时间序列（设备路径 + 测点名）。
type SeriesKey struct {
	DevicePath  string
	Measurement string
}

func (k SeriesKey) String() string {
	return fmt.Sprintf("%s.%s", k.DevicePath, k.Measurement)
}

// Point 一个时间戳对应一个类型化单值。
type Point struct {
	Timestamp int64
	Value     Value
}
