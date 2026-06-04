package wal

import (
	"testing"

	"github.com/hanami/tidets/core/tsmodel"
)

func TestCheckpointReadWrite(t *testing.T) {
	dir := t.TempDir()

	if off, err := ReadCheckpoint(dir); err != nil || off != 0 {
		t.Fatalf("read empty checkpoint: off=%d err=%v", off, err)
	}

	if err := WriteCheckpoint(dir, 1234); err != nil {
		t.Fatalf("write checkpoint: %v", err)
	}
	if off, err := ReadCheckpoint(dir); err != nil || off != 1234 {
		t.Fatalf("read checkpoint: off=%d err=%v", off, err)
	}

	if err := WriteCheckpoint(dir, -1); err != nil {
		t.Fatalf("write negative checkpoint: %v", err)
	}
	if off, err := ReadCheckpoint(dir); err != nil || off != 0 {
		t.Fatalf("read clamped checkpoint: off=%d err=%v", off, err)
	}
}

func TestCheckpointResetClearsFile(t *testing.T) {
	dir := t.TempDir()

	if err := WriteCheckpoint(dir, 999); err != nil {
		t.Fatalf("write checkpoint: %v", err)
	}
	if _, err := Reset(dir); err != nil {
		t.Fatalf("reset: %v", err)
	}
	off, err := ReadCheckpoint(dir)
	if err != nil {
		t.Fatalf("read checkpoint after reset: %v", err)
	}
	if off != 0 {
		t.Fatalf("checkpoint should be cleared after reset, got %d", off)
	}
}

func TestReplayHonorsCheckpointOffset(t *testing.T) {
	dir := t.TempDir()

	w, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}

	// batch 1
	if err := w.AppendInsertBatch([]tsmodel.BatchRecord{
		{Key: tsmodel.SeriesKey{DevicePath: "d1", Measurement: "s1"}, Point: tsmodel.Point{Timestamp: 1, Value: tsmodel.NewDouble(1)}},
	}); err != nil {
		t.Fatal(err)
	}
	off, err := w.Offset()
	if err != nil {
		t.Fatal(err)
	}
	if err := WriteCheckpoint(dir, off); err != nil {
		t.Fatal(err)
	}

	// batch 2
	if err := w.AppendInsertBatch([]tsmodel.BatchRecord{
		{Key: tsmodel.SeriesKey{DevicePath: "d1", Measurement: "s1"}, Point: tsmodel.Point{Timestamp: 2, Value: tsmodel.NewDouble(2)}},
	}); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	var n int
	if err := Replay(dir, func(key tsmodel.SeriesKey, p tsmodel.Point) error {
		n++
		if p.Timestamp != 2 {
			t.Fatalf("expected replay from checkpoint to include only ts=2, got %d", p.Timestamp)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("expected replay count=1, got %d", n)
	}
}
