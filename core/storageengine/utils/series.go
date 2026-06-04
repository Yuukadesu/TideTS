package utils

import "github.com/hanami/tidets/core/tsmodel"

// SplitSeriesKey 将 "device.measurement" 拆为设备路径与测点名。
func SplitSeriesKey(keyStr string) (device, measurement string) {
	return tsmodel.SplitSeriesKey(keyStr)
}

// DeviceFromSeriesKey 从序列键字符串解析设备路径。
func DeviceFromSeriesKey(keyStr string) string {
	return tsmodel.DeviceFromSeriesKey(keyStr)
}
