package wal

import (
	"testing"

	"github.com/hanami/tidets/core/storageengine/model"
)

func TestWALBatchReplay(t *testing.T) {
	dir := t.TempDir()

	w, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	records := []model.BatchRecord{
		{Key: model.SeriesKey{DevicePath: "d1", Measurement: "s1"}, Point: model.Point{Timestamp: 1, Value: model.NewDouble(1)}},
		{Key: model.SeriesKey{DevicePath: "d1", Measurement: "s1"}, Point: model.Point{Timestamp: 2, Value: model.NewDouble(2)}},
	}
	if err := w.AppendInsertBatch(records); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	var n int
	if err := Replay(dir, func(key model.SeriesKey, p model.Point) error {
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
