package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestStatsCollectorFinish(t *testing.T) {
	s := newStatsCollector("insert_batch")
	s.startedAt = time.Now().Add(-10 * time.Second)
	s.Record(10*time.Millisecond, 100, nil)
	s.Record(20*time.Millisecond, 200, nil)
	s.Record(30*time.Millisecond, 300, nil)

	got := s.Finish()
	if got.Op != "insert_batch" {
		t.Fatalf("op = %q, want insert_batch", got.Op)
	}
	if got.TotalRequests != 3 {
		t.Fatalf("requests = %d, want 3", got.TotalRequests)
	}
	if got.TotalItems != 600 {
		t.Fatalf("items = %d, want 600", got.TotalItems)
	}
	if got.AverageLatency != 20*time.Millisecond {
		t.Fatalf("avg latency = %s, want 20ms", got.AverageLatency)
	}
	if got.P50Latency != 20*time.Millisecond {
		t.Fatalf("p50 = %s, want 20ms", got.P50Latency)
	}
	if got.P95Latency != 20*time.Millisecond && got.P95Latency != 30*time.Millisecond {
		t.Fatalf("p95 = %s, want near upper tail", got.P95Latency)
	}
	if got.RequestsPerSec <= 0 || got.ItemsPerSec <= 0 {
		t.Fatalf("throughput must be positive: %+v", got)
	}
}

func TestBenchStatsJSON(t *testing.T) {
	stats := benchStats{
		Op:             "insert_batch",
		TotalRequests:  10,
		TotalItems:     1000,
		Errors:         0,
		TotalDuration:  13 * time.Millisecond,
		AverageLatency: 2 * time.Millisecond,
		P50Latency:     2 * time.Millisecond,
		P95Latency:     3 * time.Millisecond,
		P99Latency:     3 * time.Millisecond,
		RequestsPerSec: 753.69,
		ItemsPerSec:    75369.07,
	}

	data, err := json.Marshal(stats)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) == "{}" {
		t.Fatalf("json must not be empty: %s", data)
	}

	var decoded map[string]any
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded["op"] != "insert_batch" {
		t.Fatalf("op = %v, want insert_batch", decoded["op"])
	}
	if decoded["totalRequests"] != float64(10) {
		t.Fatalf("totalRequests = %v, want 10", decoded["totalRequests"])
	}
	if decoded["totalItems"] != float64(1000) {
		t.Fatalf("totalItems = %v, want 1000", decoded["totalItems"])
	}
}

func TestQuantileDuration(t *testing.T) {
	values := []time.Duration{40 * time.Millisecond, 10 * time.Millisecond, 30 * time.Millisecond, 20 * time.Millisecond}
	if got := quantileDuration(values, 0.50); got != 20*time.Millisecond && got != 30*time.Millisecond {
		t.Fatalf("p50 = %s, want median-ish", got)
	}
	if got := quantileDuration(values, 0.99); got != 30*time.Millisecond {
		t.Fatalf("p99 = %s, want 30ms with current quantile rule", got)
	}
}
