package plan

import (
	"github.com/hanami/tidets/core/sql/ast"
	"github.com/hanami/tidets/core/tsmodel"
)

// Kind 执行计划类型。
type Kind int

const (
	KindInsert Kind = iota
	KindSelect
	KindCreateTimeseries
	KindShowDevices
	KindShowTimeseries
	KindDelete
	KindAccount
)

// Plan 物理计划（单机最简）。
type Plan struct {
	Kind Kind

	Insert           *Insert
	Select           *Select
	CreateTimeseries *CreateTimeseries
	ShowDevices      *ShowDevices
	ShowTimeseries   *ShowTimeseries
	Delete           *Delete
	Account          *Account
}

// AccountOp 账户管理操作。
type AccountOp int

const (
	OpCreateUser AccountOp = iota
	OpSetPassword
	OpGrant
	OpRevoke
)

// Account 账户语句的物理计划。超级用户检查在进入存储引擎之前完成。
type Account struct {
	Op       AccountOp
	Username string
	Password string
	Path     string
	Privs    []ast.AccountPriv
}

type Insert struct {
	Key    tsmodel.SeriesKey
	Points []tsmodel.Point
}

type Select struct {
	Key       tsmodel.SeriesKey
	Start     int64
	End       int64
	Limit     int
	Aggregate SelectAgg
}

type SelectAgg int

const (
	SelectRaw SelectAgg = iota
	SelectCount
)

type Delete struct {
	Key   tsmodel.SeriesKey
	Start int64
	End   int64
}

type CreateTimeseries struct {
	DevicePath  string
	Measurement string
	DataType    tsmodel.DataType
}

type ShowDevices struct {
	Pattern string
}

type ShowTimeseries struct {
	DevicePath string
}
