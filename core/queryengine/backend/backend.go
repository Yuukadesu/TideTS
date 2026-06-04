package backend

import (
	"context"

	"github.com/hanami/tidets/core/tsmodel"
)

// Backend 查询层依赖的存储能力（与 SQL 解耦，便于后续扩展 Delete/Scan）。
type Backend interface {
	Insert(ctx context.Context, key tsmodel.SeriesKey, p tsmodel.Point) error
	InsertBatch(ctx context.Context, key tsmodel.SeriesKey, points []tsmodel.Point) error
	Query(ctx context.Context, key tsmodel.SeriesKey, start, end int64, limit int) ([]tsmodel.Point, error)
	Count(ctx context.Context, key tsmodel.SeriesKey, start, end int64) (int, error)
	DeleteRange(ctx context.Context, key tsmodel.SeriesKey, start, end int64) (int, error)
}
