package main

import (
	"context"
	"fmt"
	"math"
	"sync/atomic"

	"github.com/hanami/tidets/client/session"
)

func newQueryRangeWorkload(cfg benchConfig) workload {
	requests := max(1, int(math.Ceil(float64(cfg.points)/float64(cfg.rangeSize))))
	var nextWindow atomic.Int64
	return workload{
		name:     "query_range",
		requests: requests,
		setup: func(ctx context.Context, cfg benchConfig, s session.Session) error {
			return preloadSeries(ctx, cfg, s)
		},
		warmup: func(ctx context.Context, cfg benchConfig, s session.Session) (int, error) {
			start, end := nextRange(&nextWindow, cfg.rangeSize, cfg.points)
			points, err := s.QueryRange(ctx, cfg.device, cfg.measurement, start, end)
			return len(points), err
		},
		runOne: func(ctx context.Context, cfg benchConfig, s session.Session) (int, error) {
			start, end := nextRange(&nextWindow, cfg.rangeSize, cfg.points)
			points, err := s.QueryRange(ctx, cfg.device, cfg.measurement, start, end)
			return len(points), err
		},
	}
}

func preloadSeries(ctx context.Context, cfg benchConfig, s session.Session) error {
	var ts int64
	for inserted := 0; inserted < cfg.points; inserted += cfg.batchSize {
		size := min(cfg.batchSize, cfg.points-inserted)
		batch := make([]session.BatchPoint, 0, size)
		for i := 0; i < size; i++ {
			ts++
			batch = append(batch, session.BatchPoint{
				Measurement: cfg.measurement,
				Timestamp:   ts,
				Value:       session.Double(float64(ts)),
			})
		}
		if err := s.InsertBatch(ctx, cfg.device, batch); err != nil {
			return fmt.Errorf("preload series: %w", err)
		}
	}
	return nil
}

func nextRange(counter *atomic.Int64, rangeSize, totalPoints int) (int64, int64) {
	window := counter.Add(1) - 1
	start := window*int64(rangeSize) + 1
	last := int64(totalPoints)
	end := start + int64(rangeSize) - 1
	if start > last {
		start = ((start - 1) % last) + 1
		end = start + int64(rangeSize) - 1
	}
	if end > last {
		end = last
	}
	return start, end
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
