package dataplane

import (
	commons "github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/schemaengine"
	"github.com/hanami/tidets/core/storageengine"
	"github.com/hanami/tidets/core/tsmodel"
)

// Gateway 统一写入入口：写入必经 schema 校验，再进入 storage。
type Gateway struct {
	Engine *storageengine.Engine
	Schema *schemaengine.Service
}

// New 创建数据面网关。
func New(engine *storageengine.Engine, schema *schemaengine.Service) *Gateway {
	return &Gateway{Engine: engine, Schema: schema}
}

// Insert 写入单点（ValidateInsert → Engine.Insert）。
func (g *Gateway) Insert(key tsmodel.SeriesKey, p tsmodel.Point) error {
	if _, err := g.Schema.ValidateInsert(key, p.Value); err != nil {
		return err
	}
	return g.Engine.Insert(key, p)
}

// InsertBatch 批量写入（每条先 ValidateInsert）。
func (g *Gateway) InsertBatch(records []storageengine.Record) error {
	for _, rec := range records {
		if _, err := g.Schema.ValidateInsert(rec.Key, rec.Point.Value); err != nil {
			return err
		}
	}
	return g.Engine.InsertBatch(records)
}

// InsertBatchPoints 同序列批量写入（每条先 ValidateInsert）。
func (g *Gateway) InsertBatchPoints(key tsmodel.SeriesKey, points []tsmodel.Point) error {
	if len(points) == 0 {
		return nil
	}
	records := make([]storageengine.Record, 0, len(points))
	for _, p := range points {
		if _, err := g.Schema.ValidateInsert(key, p.Value); err != nil {
			return err
		}
		records = append(records, storageengine.Record{Key: key, Point: p})
	}
	return g.Engine.InsertBatch(records)
}

// DeleteRange 删除时间范围内测点（序列须已注册 schema）。
func (g *Gateway) DeleteRange(key tsmodel.SeriesKey, start, end int64) (int, error) {
	if !g.Schema.Has(key) {
		return 0, commons.ErrStorageSchemaRequired
	}
	return g.Engine.DeleteRange(key, start, end)
}
