package plan

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
	}
	return ""
}

// NeedsWrite 是否为写操作。
func (p Plan) NeedsWrite() bool {
	return p.Kind == KindInsert
}
