package ast

import "github.com/hanami/tidets/core/tsmodel"

// Stmt 根语句。
type Stmt interface {
	stmt()
}

// InsertStmt INSERT INTO device(measurement) VALUES (ts, val)[, ...]。
type InsertStmt struct {
	DevicePath  string
	Measurement string
	Rows        []InsertRow
}

// InsertRow 单条写入行。
type InsertRow struct {
	Timestamp int64
	Value     tsmodel.Value
}

func (*InsertStmt) stmt() {}

// CmpOp 时间谓词比较符。
type CmpOp int

const (
	CmpGTE CmpOp = iota
	CmpLTE
	CmpGT
	CmpLT
	CmpEQ
)

// TimePredicate WHERE time <op> <n>。
type TimePredicate struct {
	Op        CmpOp
	Timestamp int64
}

// SelectAgg SELECT 聚合类型。
type SelectAgg int

const (
	SelectRaw SelectAgg = iota
	SelectCount
)

// SelectStmt SELECT measurement FROM device [WHERE ...] [LIMIT n] 或 COUNT。
type SelectStmt struct {
	Aggregate   SelectAgg
	DevicePath  string
	Measurement string
	TimeWhere   []TimePredicate
	Limit       int // 0 表示由执行层默认
}

func (*SelectStmt) stmt() {}

// DeleteStmt DELETE FROM device(measurement) WHERE time ...。
type DeleteStmt struct {
	DevicePath  string
	Measurement string
	TimeWhere   []TimePredicate
}

func (*DeleteStmt) stmt() {}

// CreateTimeseriesStmt CREATE TIMESERIES device(measurement) WITH DATATYPE=type。
type CreateTimeseriesStmt struct {
	DevicePath  string
	Measurement string
	DataType    tsmodel.DataType
}

func (*CreateTimeseriesStmt) stmt() {}

// ShowDevicesStmt SHOW DEVICES [pattern]。
type ShowDevicesStmt struct {
	Pattern string // 空表示全部
}

func (*ShowDevicesStmt) stmt() {}

// ShowTimeseriesStmt SHOW TIMESERIES devicePath。
type ShowTimeseriesStmt struct {
	DevicePath string
}

func (*ShowTimeseriesStmt) stmt() {}
