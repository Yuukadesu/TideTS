package metrics

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
)

func TestHTTPServerExposesMetricsEndpoint(t *testing.T) {
	reg := NewRegistry()
	reg.ObserveRPC("QueryRange", codes.OK, time.Millisecond)
	reg.ObserveSQL(-1, false, "parse", time.Millisecond)
	server := NewHTTPServer(":0", reg)
	ts := httptest.NewServer(server.Handler)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/metrics")
	if err != nil {
		t.Fatalf("get /metrics: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code = %d, want 200", resp.StatusCode)
	}
	text := string(body)
	if !strings.Contains(text, "tidets_grpc_requests_total") {
		t.Fatalf("metrics body does not contain TideTS metrics")
	}
	if !strings.Contains(text, "tidets_build_info") || !strings.Contains(text, "tidets_node_uptime_seconds") || !strings.Contains(text, "tidets_node_start_time_seconds") {
		t.Fatalf("metrics body does not contain node info metrics")
	}
	if !strings.Contains(text, "tidets_sql_requests_total") {
		t.Fatalf("metrics body does not contain sql outcome metrics")
	}
}
