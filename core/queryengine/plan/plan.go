package plan

import "github.com/hanami/tidets/core/tsmodel"

// Kind 执行计划类型。
type Kind int

const (
	KindInsert Kind = iota
	KindSelect
	KindCreateTimeseries
	KindShowDevices
	KindShowTimeseries
	KindDelete
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
