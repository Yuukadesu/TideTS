package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/hanami/tidets/client/session"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		fail(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s, err := session.New(
		session.WithHost(cfg.host),
		session.WithPort(cfg.port),
		session.WithUsername(cfg.username),
		session.WithPassword(cfg.password),
		session.WithFetchSize(cfg.fetchSize),
	)
	if err != nil {
		fail(err)
	}
	if err := s.Open(ctx); err != nil {
		fail(err)
	}
	defer func() {
		if err := s.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "close session: %v\n", err)
		}
	}()

	wl, err := getWorkload(cfg)
	if err != nil {
		fail(err)
	}

	beforePath, afterPath := metricsPaths(cfg)
	if err := snapshotMetrics(cfg.metricsURL, beforePath); err != nil {
		fmt.Fprintf(os.Stderr, "snapshot metrics before: %v\n", err)
	}

	stats, err := runWorkload(ctx, cfg, s, wl)
	if err != nil {
		fail(err)
	}

	if err := snapshotMetrics(cfg.metricsURL, afterPath); err != nil {
		fmt.Fprintf(os.Stderr, "snapshot metrics after: %v\n", err)
	}

	printStats(stats)
	if cfg.jsonOut != "" {
		if err := writeStatsJSON(cfg.jsonOut, stats); err != nil {
			fail(err)
		}
	}
}

func metricsPaths(cfg benchConfig) (before, after string) {
	if cfg.resultDir == "" || cfg.metricsURL == "" {
		return "", ""
	}
	return filepath.Join(cfg.resultDir, "metrics_before.prom"), filepath.Join(cfg.resultDir, "metrics_after.prom")
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "bench: %v\n", err)
	os.Exit(1)
}
