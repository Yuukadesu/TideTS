package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/hanami/tidets/client/session"
)

type benchConfig struct {
	host        string
	port        int
	username    string
	password    string
	fetchSize   int
	op          string
	device      string
	measurement string
	points      int
	batchSize   int
	rangeSize   int
	concurrency int
	warmup      int
	metricsURL  string
	jsonOut     string
	resultDir   string
	timeout     time.Duration
}

func parseConfig() (benchConfig, error) {
	return parseConfigFrom(os.Args[1:])
}

func parseConfigFrom(args []string) (benchConfig, error) {
	cfg := benchConfig{}
	fs := flag.NewFlagSet("bench", flag.ContinueOnError)
	fs.StringVar(&cfg.host, "host", envOr("TIDETS_HOST", session.DefaultHost), "DataNode host")
	fs.IntVar(&cfg.port, "port", envIntOr("TIDETS_PORT", session.DefaultPort), "DataNode gRPC port")
	fs.StringVar(&cfg.username, "username", session.DefaultUsername, "session username")
	fs.StringVar(&cfg.password, "password", session.DefaultPassword, "session password")
	fs.IntVar(&cfg.fetchSize, "fetch-size", session.DefaultFetchSize, "session fetch size")
	fs.StringVar(&cfg.op, "op", "insert_batch", "benchmark op: insert_point|insert_batch|query_range|count_sql|delete_sql")
	fs.StringVar(&cfg.device, "device", "root.bench.d1", "target device path")
	fs.StringVar(&cfg.measurement, "measurement", "temperature", "target measurement")
	fs.IntVar(&cfg.points, "points", 10000, "total logical points to benchmark")
	fs.IntVar(&cfg.batchSize, "batch-size", 100, "batch size for insert_batch and preload")
	fs.IntVar(&cfg.rangeSize, "range-size", 1000, "query/count/delete time window size")
	fs.IntVar(&cfg.concurrency, "concurrency", 1, "number of concurrent workers")
	fs.IntVar(&cfg.warmup, "warmup", 100, "number of warmup logical units before benchmark")
	fs.StringVar(&cfg.metricsURL, "metrics", "", "optional metrics endpoint, e.g. http://127.0.0.1:9090/metrics")
	fs.StringVar(&cfg.jsonOut, "json-out", "", "optional JSON output file")
	fs.StringVar(&cfg.resultDir, "result-dir", "", "optional result directory for JSON and metrics snapshots")
	fs.DurationVar(&cfg.timeout, "timeout", 30*time.Second, "connection and request timeout")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}

	if cfg.host == "" {
		return cfg, fmt.Errorf("host is required")
	}
	if cfg.port <= 0 {
		return cfg, fmt.Errorf("port must be positive")
	}
	if cfg.points <= 0 {
		return cfg, fmt.Errorf("points must be positive")
	}
	if cfg.batchSize <= 0 {
		return cfg, fmt.Errorf("batch-size must be positive")
	}
	if cfg.rangeSize <= 0 {
		return cfg, fmt.Errorf("range-size must be positive")
	}
	if cfg.concurrency <= 0 {
		return cfg, fmt.Errorf("concurrency must be positive")
	}
	if cfg.warmup < 0 {
		return cfg, fmt.Errorf("warmup must be >= 0")
	}
	if cfg.resultDir != "" {
		if err := os.MkdirAll(cfg.resultDir, 0o755); err != nil {
			return cfg, err
		}
		if cfg.jsonOut == "" {
			cfg.jsonOut = filepath.Join(cfg.resultDir, "result.json")
		} else if !filepath.IsAbs(cfg.jsonOut) {
			cfg.jsonOut = filepath.Join(cfg.resultDir, cfg.jsonOut)
		}
	}
	return cfg, nil
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envIntOr(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		var n int
		if _, err := fmt.Sscanf(v, "%d", &n); err == nil && n > 0 {
			return n
		}
	}
	return fallback
}
