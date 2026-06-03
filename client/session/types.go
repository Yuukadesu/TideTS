package session

// Point 查询返回的一个数据点。
type Point struct {
	Timestamp int64
	Value     Value
}

// BatchPoint 批量写入的一个测点。
type BatchPoint struct {
	Measurement string
	Timestamp   int64
	Value       Value
}
