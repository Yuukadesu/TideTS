package ast

import "github.com/hanami/tidets/core/storageengine/model"

// Stmt 根语句。
type Stmt interface {
	stmt()
}

// InsertStmt INSERT INTO device(measurement) VALUES (ts, val)。
type InsertStmt struct {
	DevicePath  string
	Measurement string
	Timestamp   int64
	Value       model.Value
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

// SelectStmt SELECT measurement FROM device [WHERE ...] [LIMIT n]。
type SelectStmt struct {
	DevicePath  string
	Measurement string
	TimeWhere   []TimePredicate
	Limit       int // 0 表示由执行层默认
}

func (*SelectStmt) stmt() {}
