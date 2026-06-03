package utils

import (
	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/model"
)

// ValidatePoint 校验单点写入参数。
func ValidatePoint(key model.SeriesKey, p model.Point) error {
	if key.DevicePath == "" || key.Measurement == "" {
		return commons.ErrStorageDeviceMeasurementRequired
	}
	if p.Timestamp <= 0 {
		return commons.ErrStorageTimestampInvalid
	}
	return p.Value.Validate()
}

// CheckSeriesValueType 同序列测点类型须一致。
func CheckSeriesValueType(existing []model.Point, p model.Point) error {
	if len(existing) == 0 {
		return nil
	}
	if existing[0].Value.Type != p.Value.Type {
		return commons.ErrStorageDataTypeMismatch
	}
	return nil
}
