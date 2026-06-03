package examples_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hanami/tidets/client/session"
)

// TestSessionInsertLifecycle 端到端：需先手动启动 DataNode，测试只调用 session API。
//
//	# 终端 1
//	go run ./cmd/datanode -data-dir ./data
//
//	# 终端 2
//	go test ./examples/... -run TestSessionInsertLifecycle -count=1
//
// 可选环境变量：TIDETS_HOST（默认 127.0.0.1）、TIDETS_PORT（默认 5556）。
func TestSessionInsertLifecycle(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	host := envOr("TIDETS_HOST", session.DefaultHost)
	port := envPortOr("TIDETS_PORT", session.DefaultPort)
	addr := fmt.Sprintf("%s:%d", host, port)

	if err := pingTCP(addr, 2*time.Second); err != nil {
		t.Skipf("datanode not running at %s, start it first: %v", addr, err)
	}

	ctx := context.Background()
	s, err := session.New(session.WithHost(host), session.WithPort(port))
	if err != nil {
		t.Fatalf("new session: %v", err)
	}
	if err := s.Open(ctx); err != nil {
		t.Fatalf("open session: %v", err)
	}

	device := "root.sg1.d1"
	measurement := "temperature"
	ts := time.Now().UnixMilli()
	value := session.Double(36.5)

	if err := s.InsertPoint(ctx, device, measurement, ts, value); err != nil {
		t.Fatalf("insert point: %v", err)
	}

	points, err := s.QueryRange(ctx, device, measurement, ts-1000, ts+1000)
	if err != nil {
		t.Fatalf("query range: %v", err)
	}
	if len(points) != 1 || points[0].Timestamp != ts || !points[0].Value.Equal(value) {
		t.Fatalf("unexpected points: %+v", points)
	}

	if err := s.Close(); err != nil {
		t.Fatalf("close session: %v", err)
	}

	if err := s.InsertPoint(ctx, device, measurement, ts+1, session.Double(1.0)); err == nil {
		t.Fatal("expected error when inserting after close")
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envPortOr(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil || p <= 0 {
			return fallback
		}
		return p
	}
	return fallback
}

func pingTCP(addr string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return err
	}
	return conn.Close()
}
