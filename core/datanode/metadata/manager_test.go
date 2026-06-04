package metadata

import (
	"testing"

	"github.com/hanami/tidets/core/schemaengine"
	"github.com/hanami/tidets/core/tsmodel"
)

func openTestManager(t *testing.T) (*schemaengine.Service, *Manager) {
	t.Helper()
	dir := t.TempDir()
	schemaSvc, err := schemaengine.Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	return schemaSvc, New(schemaSvc)
}

func TestManagerListAfterInsert(t *testing.T) {
	schemaSvc, mgr := openTestManager(t)
	defer schemaSvc.Close()

	key := tsmodel.SeriesKey{DevicePath: "root.sg1.d1", Measurement: "temperature"}
	if _, err := schemaSvc.ValidateInsert(key, tsmodel.NewDouble(1)); err != nil {
		t.Fatal(err)
	}

	devices := mgr.ListDevices("root.sg1")
	if len(devices) != 1 || devices[0].Path != "root.sg1.d1" {
		t.Fatalf("devices=%+v", devices)
	}
	series := mgr.ListTimeseries("root.sg1.d1")
	if len(series) != 1 || series[0].Measurement != "temperature" {
		t.Fatalf("series=%+v", series)
	}
	children := mgr.ChildPaths("root.sg1")
	if len(children) != 1 || !children[0].IsDevice {
		t.Fatalf("children=%+v", children)
	}
}

func TestManagerReconcileFromStorage(t *testing.T) {
	schemaSvc, mgr := openTestManager(t)
	defer schemaSvc.Close()

	types := map[string]tsmodel.DataType{
		"root.sg1.d1.temperature": tsmodel.DataTypeDouble,
		"root.sg1.d1.humidity":    tsmodel.DataTypeFloat,
	}
	if err := mgr.ReconcileFromStorage(types); err != nil {
		t.Fatal(err)
	}
	if len(mgr.ListDevices("")) != 1 {
		t.Fatalf("devices=%+v", mgr.ListDevices(""))
	}
	if len(mgr.ListTimeseries("root.sg1.d1")) != 2 {
		t.Fatalf("series=%+v", mgr.ListTimeseries("root.sg1.d1"))
	}
}

func TestManagerListDevicesPattern(t *testing.T) {
	schemaSvc, mgr := openTestManager(t)
	defer schemaSvc.Close()

	for _, key := range []tsmodel.SeriesKey{
		{DevicePath: "root.d1", Measurement: "s1"},
		{DevicePath: "root.sg1.d1", Measurement: "s1"},
		{DevicePath: "root.sg1.d2", Measurement: "s1"},
		{DevicePath: "root.sg2.d1", Measurement: "s1"},
	} {
		if _, err := schemaSvc.ValidateInsert(key, tsmodel.NewDouble(1)); err != nil {
			t.Fatal(err)
		}
	}

	if len(mgr.ListDevices("root.sg1.**")) != 2 {
		t.Fatalf("sg1.** devices=%+v", mgr.ListDevices("root.sg1.**"))
	}
	if len(mgr.ListDevices("root.*")) != 1 {
		t.Fatalf("root.* devices=%+v", mgr.ListDevices("root.*"))
	}
}
