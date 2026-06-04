package schemaengine

import (
	"fmt"
	"time"

	"github.com/hanami/tidets/core/tsmodel"
)

// Encoding 序列值编码方式（对齐 IoTDB 子集，当前仅 PLAIN）。
type Encoding uint8

const (
	EncodingPlain Encoding = 0
)

// Compressor 压缩算法（当前仅不压缩）。
type Compressor uint8

const (
	CompressorUncompressed Compressor = 0
)

// Timeseries 一条时间序列的 schema 定义。
type Timeseries struct {
	DevicePath  string
	Measurement string
	DataType    tsmodel.DataType
	Encoding    Encoding
	Compressor  Compressor
	CreatedAt   int64
}

// Key 返回 device + measurement 标识。
func (ts Timeseries) Key() tsmodel.SeriesKey {
	return tsmodel.SeriesKey{DevicePath: ts.DevicePath, Measurement: ts.Measurement}
}

// FullPath 返回 IoTDB 风格完整路径 root.sg1.d1.temperature。
func (ts Timeseries) FullPath() string {
	return fmt.Sprintf("%s.%s", ts.DevicePath, ts.Measurement)
}

// NewTimeseries 构造带默认 encoding/compressor 的 schema。
func NewTimeseries(devicePath, measurement string, dt tsmodel.DataType) Timeseries {
	return Timeseries{
		DevicePath:  devicePath,
		Measurement: measurement,
		DataType:    dt,
		Encoding:    EncodingPlain,
		Compressor:  CompressorUncompressed,
		CreatedAt:   time.Now().UnixMilli(),
	}
}
