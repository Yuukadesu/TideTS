package storage

import (
	"context"

	"github.com/hanami/tidets/core/storageengine"
)

// EngineAdapter 将 storageengine.Engine 适配为 SQL 执行层 Backend。
type EngineAdapter struct {
	Engine *storageengine.Engine
}

func (a *EngineAdapter) Insert(ctx context.Context, key storageengine.SeriesKey, p storageengine.Point) error {
	_ = ctx
	return a.Engine.Insert(key, p)
}

func (a *EngineAdapter) Query(ctx context.Context, key storageengine.SeriesKey, start, end int64, limit int) ([]storageengine.Point, error) {
	_ = ctx
	return a.Engine.Query(key, start, end, limit)
}
