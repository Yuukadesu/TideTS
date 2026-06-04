package schemaengine

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel"
)

func TestServiceCreateAndValidate(t *testing.T) {
	dir := t.TempDir()
	svc, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer svc.Close()

	key := tsmodel.SeriesKey{DevicePath: "root.sg1.d1", Measurement: "temperature"}
	_, err = svc.CreateTimeseries(key.DevicePath, key.Measurement, tsmodel.DataTypeDouble)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := svc.CreateTimeseries(key.DevicePath, key.Measurement, tsmodel.DataTypeDouble); !commons.Is(err, commons.ErrSchemaTimeseriesExists) {
		t.Fatalf("want exists, got %v", err)
	}

	ts, err := svc.ValidateInsert(key, tsmodel.NewDouble(25.5))
	if err != nil {
		t.Fatal(err)
	}
	if ts.DataType != tsmodel.DataTypeDouble {
		t.Fatalf("datatype=%v", ts.DataType)
	}

	_, err = svc.ValidateInsert(key, tsmodel.NewInt32(1))
	if !commons.Is(err, commons.ErrSchemaDataTypeMismatch) {
		t.Fatalf("want mismatch, got %v", err)
	}

	mlogPath := filepath.Join(dir, "system", "schema", mlogFileName)
	if _, err := os.Stat(mlogPath); err != nil {
		t.Fatalf("mlog missing: %v", err)
	}
}

func TestServiceAutoRegisterOnInsert(t *testing.T) {
	dir := t.TempDir()
	svc, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer svc.Close()

	key := tsmodel.SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}
	ts, err := svc.ValidateInsert(key, tsmodel.NewFloat(1.2))
	if err != nil {
		t.Fatal(err)
	}
	if !svc.Has(key) || ts.DataType != tsmodel.DataTypeFloat {
		t.Fatalf("auto register failed: %+v", ts)
	}
}

func TestServiceReload(t *testing.T) {
	dir := t.TempDir()
	svc, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	key := tsmodel.SeriesKey{DevicePath: "root.a.d1", Measurement: "m1"}
	if _, err := svc.CreateTimeseries(key.DevicePath, key.Measurement, tsmodel.DataTypeInt64); err != nil {
		t.Fatal(err)
	}
	if err := svc.Close(); err != nil {
		t.Fatal(err)
	}

	svc2, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer svc2.Close()
	ts, ok := svc2.Get(key)
	if !ok || ts.DataType != tsmodel.DataTypeInt64 {
		t.Fatalf("reload failed: ok=%v ts=%+v", ok, ts)
	}
}

func TestServiceSnapshotTrigger(t *testing.T) {
	dir := t.TempDir()
	svc, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer svc.Close()

	for i := 0; i < snapshotEveryOps; i++ {
		key := tsmodel.SeriesKey{DevicePath: "root.sg1.d1", Measurement: fmt.Sprintf("m%d", i)}
		if _, err := svc.ValidateInsert(key, tsmodel.NewDouble(float64(i))); err != nil {
			t.Fatal(err)
		}
	}

	snapPath := filepath.Join(dir, "system", "schema", snapshotFileName)
	if _, err := os.Stat(snapPath); err != nil {
		t.Fatalf("snapshot not created: %v", err)
	}
}
