package backend

import (
	"context"

	"github.com/hanami/tidets/core/datanode/metadata"
	"github.com/hanami/tidets/core/tsmodel"
)

// CatalogBackend DDL 与元数据 SHOW 能力。
type CatalogBackend interface {
	CreateTimeseries(ctx context.Context, devicePath, measurement string, dt tsmodel.DataType) error
	ListDevices(ctx context.Context, pattern string) ([]metadata.DeviceInfo, error)
	ListTimeseries(ctx context.Context, devicePath string) ([]metadata.TimeseriesInfo, error)
}

// MetadataCatalog 委托 metadata.Manager。
type MetadataCatalog struct {
	Meta *metadata.Manager
}

func (c *MetadataCatalog) CreateTimeseries(ctx context.Context, devicePath, measurement string, dt tsmodel.DataType) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if c.Meta == nil || c.Meta.Schema() == nil {
		return nil
	}
	_, err := c.Meta.Schema().CreateTimeseries(devicePath, measurement, dt)
	return err
}

func (c *MetadataCatalog) ListDevices(ctx context.Context, pattern string) ([]metadata.DeviceInfo, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if c.Meta == nil {
		return nil, nil
	}
	return c.Meta.ListDevices(pattern), nil
}

func (c *MetadataCatalog) ListTimeseries(ctx context.Context, devicePath string) ([]metadata.TimeseriesInfo, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if c.Meta == nil {
		return nil, nil
	}
	return c.Meta.ListTimeseries(devicePath), nil
}
