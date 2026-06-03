package storage

import (
	"context"

	"github.com/hanami/tidets/core/storageengine"
)

// Backend 查询层依赖的存储能力（与 SQL 解耦，便于后续扩展 Delete/Scan）。
type Backend interface {
	Insert(ctx context.Context, key storageengine.SeriesKey, p storageengine.Point) error
	Query(ctx context.Context, key storageengine.SeriesKey, start, end int64, limit int) ([]storageengine.Point, error)
}
