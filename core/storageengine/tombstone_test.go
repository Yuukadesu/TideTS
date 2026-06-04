package storageengine

import (
	"path/filepath"
	"testing"

	"github.com/hanami/tidets/core/storageengine/wal"
	"github.com/hanami/tidets/core/tsmodel"
)

func TestTombstoneIndexFilter(t *testing.T) {
	idx := newTombstoneIndex(nil)
	key := "root.d1.s1"
	if err := idx.Mark(key, 100, 200); err != nil {
		t.Fatal(err)
	}

	pts := []tsmodel.Point{
		{Timestamp: 50, Value: tsmodel.NewDouble(1)},
		{Timestamp: 150, Value: tsmodel.NewDouble(2)},
		{Timestamp: 250, Value: tsmodel.NewDouble(3)},
	}
	out := idx.Filter(key, pts)
	if len(out) != 2 || out[0].Timestamp != 50 || out[1].Timestamp != 250 {
		t.Fatalf("filter: %+v", out)
	}
}

func TestTombstoneLogRoundTrip(t *testing.T) {
	dir := t.TempDir()
	log, err := openTombstoneLog(dir, wal.SyncAlways)
	if err != nil {
		t.Fatal(err)
	}
	idx := newTombstoneIndex(log)
	if err := idx.Mark("root.d1.s1", 10, 20); err != nil {
		t.Fatal(err)
	}
	if err := idx.Mark("root.d1.s1", 30, 40); err != nil {
		t.Fatal(err)
	}
	if err := idx.Close(); err != nil {
		t.Fatal(err)
	}

	reopened, err := openTombstoneLog(dir, wal.SyncAlways)
	if err != nil {
		t.Fatal(err)
	}
	loaded := newTombstoneIndex(reopened)
	if err := loadTombstoneLog(dir, loaded.Restore); err != nil {
		t.Fatal(err)
	}
	pts := []tsmodel.Point{
		{Timestamp: 5, Value: tsmodel.NewDouble(1)},
		{Timestamp: 15, Value: tsmodel.NewDouble(2)},
		{Timestamp: 25, Value: tsmodel.NewDouble(3)},
		{Timestamp: 35, Value: tsmodel.NewDouble(4)},
	}
	got := loaded.Filter("root.d1.s1", pts)
	if len(got) != 2 || got[0].Timestamp != 5 || got[1].Timestamp != 25 {
		t.Fatalf("round trip filter: %+v", got)
	}
	if err := loaded.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestTombstonePruneRewrite(t *testing.T) {
	dir := t.TempDir()
	log, err := openTombstoneLog(dir, wal.SyncAlways)
	if err != nil {
		t.Fatal(err)
	}
	idx := newTombstoneIndex(log)
	if err := idx.Mark("root.d1.s1", 10, 20); err != nil {
		t.Fatal(err)
	}
	if err := idx.Mark("root.d1.s1", 30, 40); err != nil {
		t.Fatal(err)
	}
	drop := map[string]map[timeRange]struct{}{
		"root.d1.s1": {
			{start: 10, end: 20}: {},
		},
	}
	if err := idx.Prune(drop); err != nil {
		t.Fatal(err)
	}
	if err := idx.Close(); err != nil {
		t.Fatal(err)
	}

	reopened, err := openTombstoneLog(dir, wal.SyncAlways)
	if err != nil {
		t.Fatal(err)
	}
	loaded := newTombstoneIndex(reopened)
	if err := loadTombstoneLog(dir, loaded.Restore); err != nil {
		t.Fatal(err)
	}
	got := loaded.Snapshot()["root.d1.s1"]
	if len(got) != 1 || got[0].start != 30 || got[0].end != 40 {
		t.Fatalf("pruned ranges: %+v", got)
	}
	if _, err := filepath.Abs(filepath.Join(dir, tombstoneLogFilename)); err != nil {
		t.Fatal(err)
	}
	if err := loaded.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestMergeTimeRanges(t *testing.T) {
	merged := mergeTimeRanges([]timeRange{
		{100, 120},
		{115, 150},
		{200, 210},
	})
	if len(merged) != 2 || merged[0].end != 150 || merged[1].start != 200 {
		t.Fatalf("merged: %+v", merged)
	}
}
