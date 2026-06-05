package main

import (
	"context"
	"math"
	"sync/atomic"

	"github.com/hanami/tidets/client/session"
)

func newInsertPointWorkload(cfg benchConfig) workload {
	var nextTS atomic.Int64
	return workload{
		name:     "insert_point",
		requests: cfg.points,
		warmup: func(ctx context.Context, cfg benchConfig, s session.Session) (int, error) {
			ts := nextTS.Add(1)
			return 1, s.InsertPoint(ctx, cfg.device, cfg.measurement, ts, session.Double(float64(ts)))
		},
		runOne: func(ctx context.Context, cfg benchConfig, s session.Session) (int, error) {
			ts := nextTS.Add(1)
			return 1, s.InsertPoint(ctx, cfg.device, cfg.measurement, ts, session.Double(float64(ts)))
		},
	}
}

func newInsertBatchWorkload(cfg benchConfig) workload {
	var nextTS atomic.Int64
	requests := int(math.Ceil(float64(cfg.points) / float64(cfg.batchSize)))
	return workload{
		name:     "insert_batch",
		requests: requests,
		warmup: func(ctx context.Context, cfg benchConfig, s session.Session) (int, error) {
			points := makeBatchPoints(&nextTS, cfg.batchSize, cfg.measurement)
			return len(points), s.InsertBatch(ctx, cfg.device, points)
		},
		runOne: func(ctx context.Context, cfg benchConfig, s session.Session) (int, error) {
			points := makeBatchPoints(&nextTS, cfg.batchSize, cfg.measurement)
			return len(points), s.InsertBatch(ctx, cfg.device, points)
		},
	}
}

func makeBatchPoints(counter *atomic.Int64, size int, measurement string) []session.BatchPoint {
	points := make([]session.BatchPoint, 0, size)
	for i := 0; i < size; i++ {
		ts := counter.Add(1)
		points = append(points, session.BatchPoint{
			Measurement: measurement,
			Timestamp:   ts,
			Value:       session.Double(float64(ts)),
		})
	}
	return points
}
