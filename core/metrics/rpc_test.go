package metrics

import (
	"context"
	"testing"

	dto "github.com/prometheus/client_model/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUnaryServerInterceptorRecordsMetrics(t *testing.T) {
	reg := NewRegistry()
	interceptor := reg.UnaryServerInterceptor()
	info := &grpc.UnaryServerInfo{FullMethod: "/tidets.datanode.v1.DataNodeSessionService/QueryRange"}

	if _, err := interceptor(context.Background(), nil, info, func(context.Context, any) (any, error) {
		return "ok", nil
	}); err != nil {
		t.Fatalf("interceptor ok handler: %v", err)
	}
	if _, err := interceptor(context.Background(), nil, info, func(context.Context, any) (any, error) {
		return nil, status.Error(codes.InvalidArgument, "bad query")
	}); status.Code(err) != codes.InvalidArgument {
		t.Fatalf("interceptor err handler code = %v, want %v", status.Code(err), codes.InvalidArgument)
	}

	families, err := reg.reg.Gather()
	if err != nil {
		t.Fatalf("gather metrics: %v", err)
	}

	if got := counterValue(t, families, "tidets_grpc_requests_total", map[string]string{"method": "QueryRange", "code": "OK"}); got != 1 {
		t.Fatalf("grpc ok counter = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_grpc_requests_total", map[string]string{"method": "QueryRange", "code": "InvalidArgument"}); got != 1 {
		t.Fatalf("grpc invalid counter = %v, want 1", got)
	}
	if got := histogramCount(t, families, "tidets_grpc_request_duration_seconds", map[string]string{"method": "QueryRange"}); got != 2 {
		t.Fatalf("grpc duration count = %d, want 2", got)
	}
}

func TestObserveRPCItemsRecordsVolumeMetrics(t *testing.T) {
	reg := NewRegistry()
	reg.ObserveRPCItems("InsertBatch", 3, 3)
	reg.ObserveRPCItems("QueryRange", 1, 5)
	reg.ObserveRPCItems("ExecuteSQL", 1, 2)

	families, err := reg.reg.Gather()
	if err != nil {
		t.Fatalf("gather metrics: %v", err)
	}

	if got := counterValue(t, families, "tidets_grpc_request_items_total", map[string]string{"method": "InsertBatch"}); got != 3 {
		t.Fatalf("insert batch request items = %v, want 3", got)
	}
	if got := counterValue(t, families, "tidets_grpc_response_items_total", map[string]string{"method": "QueryRange"}); got != 5 {
		t.Fatalf("query range response items = %v, want 5", got)
	}
	if got := counterValue(t, families, "tidets_grpc_response_items_total", map[string]string{"method": "ExecuteSQL"}); got != 2 {
		t.Fatalf("execute sql response items = %v, want 2", got)
	}
}

func gaugeValue(t *testing.T, families []*dto.MetricFamily, name string) float64 {
	t.Helper()
	for _, family := range families {
		if family.GetName() != name || len(family.GetMetric()) == 0 {
			continue
		}
		return family.GetMetric()[0].GetGauge().GetValue()
	}
	t.Fatalf("metric family %s not found", name)
	return 0
}

func counterValue(t *testing.T, families []*dto.MetricFamily, name string, labels map[string]string) float64 {
	t.Helper()
	metric := findMetric(t, families, name, labels)
	return metric.GetCounter().GetValue()
}

func histogramCount(t *testing.T, families []*dto.MetricFamily, name string, labels map[string]string) uint64 {
	t.Helper()
	metric := findMetric(t, families, name, labels)
	return metric.GetHistogram().GetSampleCount()
}

func findMetric(t *testing.T, families []*dto.MetricFamily, name string, labels map[string]string) *dto.Metric {
	t.Helper()
	for _, family := range families {
		if family.GetName() != name {
			continue
		}
		for _, metric := range family.GetMetric() {
			if metricHasLabels(metric, labels) {
				return metric
			}
		}
	}
	t.Fatalf("metric %s with labels %v not found", name, labels)
	return nil
}

func metricHasLabels(metric *dto.Metric, want map[string]string) bool {
	if len(metric.GetLabel()) != len(want) {
		return false
	}
	for _, label := range metric.GetLabel() {
		if want[label.GetName()] != label.GetValue() {
			return false
		}
	}
	return true
}
