package plan

import "strings"

// DevicePath 返回计划涉及的设备路径（用于鉴权）。
func (p Plan) DevicePath() string {
	switch p.Kind {
	case KindInsert:
		if p.Insert != nil {
			return p.Insert.Key.DevicePath
		}
	case KindSelect:
		if p.Select != nil {
			return p.Select.Key.DevicePath
		}
	case KindDelete:
		if p.Delete != nil {
			return p.Delete.Key.DevicePath
		}
	case KindCreateTimeseries:
		if p.CreateTimeseries != nil {
			return p.CreateTimeseries.DevicePath
		}
	case KindShowDevices:
		if p.ShowDevices != nil {
			pattern := p.ShowDevices.Pattern
			if pattern == "" {
				return "root"
			}
			if strings.HasSuffix(pattern, ".**") {
				return strings.TrimSuffix(pattern, ".**")
			}
			if strings.HasSuffix(pattern, ".*") {
				return strings.TrimSuffix(pattern, ".*")
			}
			return pattern
		}
	case KindShowTimeseries:
		if p.ShowTimeseries != nil {
			return p.ShowTimeseries.DevicePath
		}
	}
	return ""
}

// NeedsWrite 是否为写操作。
func (p Plan) NeedsWrite() bool {
	return p.Kind == KindInsert || p.Kind == KindCreateTimeseries || p.Kind == KindDelete
}
