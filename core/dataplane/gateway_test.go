package dataplane_test

import (
	"testing"

	"github.com/hanami/tidets/core/dataplane"
	"github.com/hanami/tidets/core/schemaengine"
	"github.com/hanami/tidets/core/storageengine"
	"github.com/hanami/tidets/core/tsmodel"
)

func TestGatewayInsertRequiresSchema(t *testing.T) {
	dir := t.TempDir()
	eng, err := storageengine.Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer eng.Close()

	schemaSvc, err := schemaengine.Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer schemaSvc.Close()

	gw := dataplane.New(eng, schemaSvc)
	key := tsmodel.SeriesKey{DevicePath: "root.sg1.d1", Measurement: "temperature"}
	if err := gw.Insert(key, tsmodel.Point{Timestamp: 1, Value: tsmodel.NewDouble(1.0)}); err != nil {
		t.Fatal(err)
	}
	if !schemaSvc.Has(key) {
		t.Fatal("schema not registered after gateway insert")
	}
}
