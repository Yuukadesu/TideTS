package tsmodel

import "strings"

// MatchDevicePattern 判断设备路径是否匹配 SHOW DEVICES 的 pattern。
// 空 pattern 表示全部设备；支持前缀、root.sg1.** 子树、root.* 单层通配。
func MatchDevicePattern(devicePath, pattern string) bool {
	if devicePath == "" || !strings.HasPrefix(devicePath, "root") {
		return false
	}
	if pattern == "" {
		return true
	}

	if strings.HasSuffix(pattern, ".**") {
		prefix := strings.TrimSuffix(pattern, ".**")
		if prefix == "root" {
			return true
		}
		return devicePath == prefix || strings.HasPrefix(devicePath, prefix+".")
	}

	if strings.HasSuffix(pattern, ".*") {
		base := strings.TrimSuffix(pattern, ".*")
		if devicePath == base {
			return true
		}
		if !strings.HasPrefix(devicePath, base+".") {
			return false
		}
		rest := strings.TrimPrefix(devicePath, base+".")
		return rest != "" && !strings.Contains(rest, ".")
	}

	return devicePath == pattern || strings.HasPrefix(devicePath, pattern+".")
}
