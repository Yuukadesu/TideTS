package metadata

import "github.com/hanami/tidets/core/tsmodel"

// DeviceInfo 设备节点摘要。
type DeviceInfo struct {
	Path           string
	MeasurementCnt int
}

// TimeseriesInfo 元数据目录中的序列条目（详细 schema 见 schemaengine）。
type TimeseriesInfo struct {
	DevicePath  string
	Measurement string
	DataType    tsmodel.DataType
	FullPath    string
}
