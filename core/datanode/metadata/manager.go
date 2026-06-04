package metadata

import (
	"github.com/hanami/tidets/core/schemaengine"
	"github.com/hanami/tidets/core/tsmodel"
)

// Manager DataNode 元数据目录：查询设备/测点路径树，持久化委托 schemaengine。
type Manager struct {
	schema *schemaengine.Service
}

// New 绑定 schema 服务。
func New(schema *schemaengine.Service) *Manager {
	return &Manager{schema: schema}
}

// ListDevices 列举设备路径（pattern 为空表示全部；支持前缀与通配符）。
func (m *Manager) ListDevices(pattern string) []DeviceInfo {
	if m.schema == nil {
		return nil
	}
	paths := m.schema.ListDevices(pattern)
	out := make([]DeviceInfo, 0, len(paths))
	for _, p := range paths {
		out = append(out, DeviceInfo{
			Path:           p,
			MeasurementCnt: len(m.schema.ListMeasurements(p)),
		})
	}
	return out
}

// ListTimeseries 列举设备下所有测点。
func (m *Manager) ListTimeseries(devicePath string) []TimeseriesInfo {
	if m.schema == nil {
		return nil
	}
	series := m.schema.ListMeasurements(devicePath)
	out := make([]TimeseriesInfo, 0, len(series))
	for _, ts := range series {
		out = append(out, TimeseriesInfo{
			DevicePath:  ts.DevicePath,
			Measurement: ts.Measurement,
			DataType:    ts.DataType,
			FullPath:    ts.FullPath(),
		})
	}
	return out
}

// ChildPaths 返回 prefix 下的直接子路径（类型与 schemaengine 一致）。
func (m *Manager) ChildPaths(prefix string) []schemaengine.ChildPath {
	if m.schema == nil {
		return nil
	}
	return m.schema.ChildPaths(prefix)
}

// HasTimeseries catalog 中是否已有该序列。
func (m *Manager) HasTimeseries(key tsmodel.SeriesKey) bool {
	if m.schema == nil {
		return false
	}
	return m.schema.Has(key)
}

// Schema 返回底层 schema 服务（DDL / 写入校验）。
func (m *Manager) Schema() *schemaengine.Service {
	return m.schema
}

// ReconcileFromStorage 用存储层已有序列补齐 catalog（启动时调用）。
func (m *Manager) ReconcileFromStorage(seriesTypes map[string]tsmodel.DataType) error {
	if m.schema == nil || len(seriesTypes) == 0 {
		return nil
	}
	keys := make([]tsmodel.SeriesKey, 0, len(seriesTypes))
	for keyStr := range seriesTypes {
		device, measurement := tsmodel.SplitSeriesKey(keyStr)
		if device == "" || measurement == "" {
			continue
		}
		keys = append(keys, tsmodel.SeriesKey{DevicePath: device, Measurement: measurement})
	}
	return m.schema.RegisterBatch(keys, func(key tsmodel.SeriesKey) tsmodel.DataType {
		return seriesTypes[key.String()]
	})
}
