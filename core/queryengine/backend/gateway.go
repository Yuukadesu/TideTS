package backend

import (
	"context"

	"github.com/hanami/tidets/core/dataplane"
	"github.com/hanami/tidets/core/storageengine"
	"github.com/hanami/tidets/core/tsmodel"
)

// EngineBackend 将写入网关与底层查询引擎组合为 queryengine Backend。
type EngineBackend struct {
	Writer *dataplane.Gateway
	Reader *storageengine.Engine
}

func (b *EngineBackend) Insert(ctx context.Context, key tsmodel.SeriesKey, p tsmodel.Point) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return b.Writer.Insert(key, p)
}

func (b *EngineBackend) InsertBatch(ctx context.Context, key tsmodel.SeriesKey, points []tsmodel.Point) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return b.Writer.InsertBatchPoints(key, points)
}

func (b *EngineBackend) Query(ctx context.Context, key tsmodel.SeriesKey, start, end int64, limit int) ([]tsmodel.Point, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return b.Reader.Query(key, start, end, limit)
}

func (b *EngineBackend) Count(ctx context.Context, key tsmodel.SeriesKey, start, end int64) (int, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	return b.Reader.Count(key, start, end)
}

func (b *EngineBackend) DeleteRange(ctx context.Context, key tsmodel.SeriesKey, start, end int64) (int, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	return b.Writer.DeleteRange(key, start, end)
}
