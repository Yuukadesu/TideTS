package utils

import (
	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel"
)

// ValidatePoint 校验单点写入参数。
func ValidatePoint(key tsmodel.SeriesKey, p tsmodel.Point) error {
	if key.DevicePath == "" || key.Measurement == "" {
		return commons.ErrStorageDeviceMeasurementRequired
	}
	if p.Timestamp <= 0 {
		return commons.ErrStorageTimestampInvalid
	}
	return p.Value.Validate()
}

// CheckSeriesValueType 同序列测点类型须一致。
func CheckSeriesValueType(existing []tsmodel.Point, p tsmodel.Point) error {
	if len(existing) == 0 {
		return nil
	}
	if existing[0].Value.Type != p.Value.Type {
		return commons.ErrStorageDataTypeMismatch
	}
	return nil
}
