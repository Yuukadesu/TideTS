package wal

import (
	"testing"

	"github.com/hanami/tidets/core/tsmodel"
)

func TestWALBatchReplay(t *testing.T) {
	dir := t.TempDir()

	w, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	records := []tsmodel.BatchRecord{
		{Key: tsmodel.SeriesKey{DevicePath: "d1", Measurement: "s1"}, Point: tsmodel.Point{Timestamp: 1, Value: tsmodel.NewDouble(1)}},
		{Key: tsmodel.SeriesKey{DevicePath: "d1", Measurement: "s1"}, Point: tsmodel.Point{Timestamp: 2, Value: tsmodel.NewDouble(2)}},
	}
	if err := w.AppendInsertBatch(records); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	var n int
	if err := Replay(dir, func(key tsmodel.SeriesKey, p tsmodel.Point) error {
		n++
		if key.DevicePath != "d1" || p.Timestamp <= 0 {
			t.Fatalf("unexpected %v %v", key, p)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("replay count=%d", n)
	}
}
