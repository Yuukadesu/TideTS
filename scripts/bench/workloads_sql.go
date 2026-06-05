package main

import (
	"context"
	"fmt"
	"math"
	"sync/atomic"

	"github.com/hanami/tidets/client/session"
)

func newCountSQLWorkload(cfg benchConfig) workload {
	requests := max(1, int(math.Ceil(float64(cfg.points)/float64(cfg.rangeSize))))
	var nextWindow atomic.Int64
	return workload{
		name:     "count_sql",
		requests: requests,
		setup: func(ctx context.Context, cfg benchConfig, s session.Session) error {
			return preloadSeries(ctx, cfg, s)
		},
		warmup: func(ctx context.Context, cfg benchConfig, s session.Session) (int, error) {
			start, end := nextRange(&nextWindow, cfg.rangeSize, cfg.points)
			return execCount(ctx, cfg, s, start, end)
		},
		runOne: func(ctx context.Context, cfg benchConfig, s session.Session) (int, error) {
			start, end := nextRange(&nextWindow, cfg.rangeSize, cfg.points)
			return execCount(ctx, cfg, s, start, end)
		},
	}
}

func newDeleteSQLWorkload(cfg benchConfig) workload {
	requests := max(1, int(math.Ceil(float64(cfg.points)/float64(cfg.rangeSize))))
	var nextWindow atomic.Int64
	return workload{
		name:     "delete_sql",
		requests: requests,
		setup: func(ctx context.Context, cfg benchConfig, s session.Session) error {
			return preloadSeries(ctx, cfg, s)
		},
		runOne: func(ctx context.Context, cfg benchConfig, s session.Session) (int, error) {
			start, end := nextRange(&nextWindow, cfg.rangeSize, cfg.points)
			return execDelete(ctx, cfg, s, start, end)
		},
	}
}

func execCount(ctx context.Context, cfg benchConfig, s session.Session, start, end int64) (int, error) {
	sql := fmt.Sprintf("SELECT COUNT(%s) FROM %s WHERE time >= %d AND time <= %d", cfg.measurement, cfg.device, start, end)
	res, err := s.ExecuteSQL(ctx, sql)
	if err != nil {
		return 0, err
	}
	if len(res.Rows) == 0 {
		return 0, nil
	}
	return int(res.Rows[0].Value.Int64), nil
}

func execDelete(ctx context.Context, cfg benchConfig, s session.Session, start, end int64) (int, error) {
	sql := fmt.Sprintf("DELETE FROM %s(%s) WHERE time >= %d AND time <= %d", cfg.device, cfg.measurement, start, end)
	res, err := s.ExecuteSQL(ctx, sql)
	if err != nil {
		return 0, err
	}
	return res.AffectedRows, nil
}
