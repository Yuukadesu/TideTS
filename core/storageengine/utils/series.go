package utils

// SplitSeriesKey 将 "device.measurement" 拆为设备路径与测点名。
func SplitSeriesKey(keyStr string) (device, measurement string) {
	for i := len(keyStr) - 1; i >= 0; i-- {
		if keyStr[i] == '.' {
			return keyStr[:i], keyStr[i+1:]
		}
	}
	return keyStr, ""
}

// DeviceFromSeriesKey 从序列键字符串解析设备路径。
func DeviceFromSeriesKey(keyStr string) string {
	device, _ := SplitSeriesKey(keyStr)
	return device
}
