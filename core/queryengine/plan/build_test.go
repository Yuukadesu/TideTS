package plan_test

import (
	"testing"

	"github.com/hanami/tidets/core/queryengine/plan"
	"github.com/hanami/tidets/core/sql/ast"
	"github.com/hanami/tidets/core/tsmodel"
)

func TestBuildCreateTimeseries(t *testing.T) {
	p, err := plan.Build(&ast.CreateTimeseriesStmt{
		DevicePath:  "root.d1",
		Measurement: "s1",
		DataType:    tsmodel.DataTypeDouble,
	})
	if err != nil {
		t.Fatal(err)
	}
	if p.Kind != plan.KindCreateTimeseries || !p.NeedsWrite() {
		t.Fatalf("plan: %+v", p)
	}
	if p.DevicePath() != "root.d1" {
		t.Fatalf("device path: %q", p.DevicePath())
	}
}

func TestBuildShowDevicesPattern(t *testing.T) {
	p, err := plan.Build(&ast.ShowDevicesStmt{Pattern: "root.sg1.**"})
	if err != nil {
		t.Fatal(err)
	}
	if p.Kind != plan.KindShowDevices || p.NeedsWrite() {
		t.Fatalf("plan: %+v", p)
	}
	if p.DevicePath() != "root.sg1" {
		t.Fatalf("auth path: %q", p.DevicePath())
	}
}

func TestBuildInvalidDevicePath(t *testing.T) {
	_, err := plan.Build(&ast.ShowTimeseriesStmt{DevicePath: "invalid.d1"})
	if err == nil {
		t.Fatal("expected error for invalid path")
	}
}
