package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hanami/tidets/client/session"
)

type workloadFunc func(ctx context.Context, cfg benchConfig, s session.Session) (items int, err error)

type workload struct {
	name     string
	setup    func(ctx context.Context, cfg benchConfig, s session.Session) error
	warmup   workloadFunc
	runOne   workloadFunc
	requests int
}

func getWorkload(cfg benchConfig) (workload, error) {
	switch cfg.op {
	case "insert_point":
		return newInsertPointWorkload(cfg), nil
	case "insert_batch":
		return newInsertBatchWorkload(cfg), nil
	case "query_range":
		return newQueryRangeWorkload(cfg), nil
	case "count_sql":
		return newCountSQLWorkload(cfg), nil
	case "delete_sql":
		return newDeleteSQLWorkload(cfg), nil
	default:
		return workload{}, fmt.Errorf("unsupported op %q", cfg.op)
	}
}

func runWorkload(ctx context.Context, cfg benchConfig, s session.Session, wl workload) (benchStats, error) {
	if wl.setup != nil {
		if err := wl.setup(ctx, cfg, s); err != nil {
			return benchStats{}, err
		}
	}
	if wl.warmup != nil && cfg.warmup > 0 {
		if err := runWarmup(ctx, cfg, s, wl.warmup); err != nil {
			return benchStats{}, err
		}
	}
	stats := newStatsCollector(wl.name)
	stats.Start()

	requests := wl.requests
	if requests <= 0 {
		requests = cfg.points
	}
	type result struct {
		d   time.Duration
		n   int
		err error
	}
	results := make(chan result, requests)
	jobs := make(chan struct{}, requests)

	var wg sync.WaitGroup
	workerCount := cfg.concurrency
	if workerCount > requests {
		workerCount = requests
	}
	if workerCount <= 0 {
		workerCount = 1
	}
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range jobs {
				begin := time.Now()
				items, err := wl.runOne(ctx, cfg, s)
				results <- result{d: time.Since(begin), n: items, err: err}
			}
		}()
	}
	for i := 0; i < requests; i++ {
		jobs <- struct{}{}
	}
	close(jobs)
	wg.Wait()
	close(results)
	for res := range results {
		stats.Record(res.d, res.n, res.err)
	}
	return stats.Finish(), nil
}

func runWarmup(ctx context.Context, cfg benchConfig, s session.Session, fn workloadFunc) error {
	for i := 0; i < cfg.warmup; i++ {
		if _, err := fn(ctx, cfg, s); err != nil {
			return err
		}
	}
	return nil
}
