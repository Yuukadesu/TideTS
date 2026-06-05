package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

type benchStats struct {
	Op             string        `json:"op"`
	TotalRequests  int           `json:"totalRequests"`
	TotalItems     int           `json:"totalItems"`
	Errors         int           `json:"errors"`
	TotalDuration  time.Duration `json:"totalDuration"`
	AverageLatency time.Duration `json:"averageLatency"`
	P50Latency     time.Duration `json:"p50Latency"`
	P95Latency     time.Duration `json:"p95Latency"`
	P99Latency     time.Duration `json:"p99Latency"`
	RequestsPerSec float64       `json:"requestsPerSec"`
	ItemsPerSec    float64       `json:"itemsPerSec"`
}

type statsCollector struct {
	op        string
	durations []time.Duration
	errors    int
	items     int
	startedAt time.Time
	endedAt   time.Time
}

func newStatsCollector(op string) *statsCollector {
	return &statsCollector{op: op}
}

func (s *statsCollector) Start() {
	s.startedAt = time.Now()
}

func (s *statsCollector) Record(duration time.Duration, items int, err error) {
	s.durations = append(s.durations, duration)
	s.items += items
	if err != nil {
		s.errors++
	}
}

func (s *statsCollector) Finish() benchStats {
	s.endedAt = time.Now()
	total := s.endedAt.Sub(s.startedAt)
	if total <= 0 {
		total = time.Nanosecond
	}
	stats := benchStats{
		Op:            s.op,
		TotalRequests: len(s.durations),
		TotalItems:    s.items,
		Errors:        s.errors,
		TotalDuration: total,
	}
	if len(s.durations) == 0 {
		return stats
	}
	stats.AverageLatency = avgDuration(s.durations)
	stats.P50Latency = quantileDuration(s.durations, 0.50)
	stats.P95Latency = quantileDuration(s.durations, 0.95)
	stats.P99Latency = quantileDuration(s.durations, 0.99)
	stats.RequestsPerSec = float64(stats.TotalRequests) / total.Seconds()
	stats.ItemsPerSec = float64(stats.TotalItems) / total.Seconds()
	return stats
}

func printStats(stats benchStats) {
	fmt.Printf("op=%s\n", stats.Op)
	fmt.Printf("requests=%d items=%d errors=%d\n", stats.TotalRequests, stats.TotalItems, stats.Errors)
	fmt.Printf("total=%s avg=%s p50=%s p95=%s p99=%s\n",
		stats.TotalDuration, stats.AverageLatency, stats.P50Latency, stats.P95Latency, stats.P99Latency)
	fmt.Printf("throughput req/s=%.2f items/s=%.2f\n", stats.RequestsPerSec, stats.ItemsPerSec)
}

func writeStatsJSON(path string, stats benchStats) error {
	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}

func avgDuration(values []time.Duration) time.Duration {
	if len(values) == 0 {
		return 0
	}
	var total time.Duration
	for _, v := range values {
		total += v
	}
	return total / time.Duration(len(values))
}

func quantileDuration(values []time.Duration, q float64) time.Duration {
	if len(values) == 0 {
		return 0
	}
	cp := append([]time.Duration(nil), values...)
	sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })
	idx := int(float64(len(cp)-1) * q)
	if idx < 0 {
		idx = 0
	}
	if idx >= len(cp) {
		idx = len(cp) - 1
	}
	return cp[idx]
}
