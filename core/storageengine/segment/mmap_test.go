package segment

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hanami/tidets/core/tsmodel"
)

func TestSealedSegmentUsesMmap(t *testing.T) {
	dir := t.TempDir()
	mgr, err := OpenManager(dir, 1)
	if err != nil {
		t.Fatal(err)
	}
	key := tsmodel.SeriesKey{DevicePath: "d1", Measurement: "s1"}
	if err := mgr.Flush(map[string][]tsmodel.Point{
		key.String(): {
			{Timestamp: 1, Value: tsmodel.NewDouble(1)},
			{Timestamp: 2, Value: tsmodel.NewDouble(2)},
		},
	}); err != nil {
		t.Fatal(err)
	}
	if mgr.SealedFileCount() != 1 {
		t.Fatalf("sealed=%d", mgr.SealedFileCount())
	}

	pts := mgr.Query(key, 1, 2)
	if len(pts) != 2 || pts[0].Value.Double != 1 || pts[1].Timestamp != 2 {
		t.Fatalf("query: %+v", pts)
	}

	entries, _ := os.ReadDir(filepath.Join(dir, SubDir))
	for _, e := range entries {
		if e.Name() == ActiveFileName {
			continue
		}
		if filepath.Ext(e.Name()) != ".seg" {
			continue
		}
		sf, err := openFileMmap(filepath.Join(dir, SubDir, e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		if sf.mapped == nil {
			t.Fatal("expected mmap-backed sealed segment")
		}
		if sf.mem != nil {
			t.Fatal("sealed segment should not load all series into memory")
		}
		_ = sf.close()
	}
}

func TestQueryTimestampsMatchesQuery(t *testing.T) {
	dir := t.TempDir()
	mgr, err := OpenManager(dir, 1)
	if err != nil {
		t.Fatal(err)
	}
	key := tsmodel.SeriesKey{DevicePath: "d1", Measurement: "s1"}
	points := []tsmodel.Point{
		{Timestamp: 10, Value: tsmodel.NewDouble(1)},
		{Timestamp: 20, Value: tsmodel.NewDouble(2)},
		{Timestamp: 30, Value: tsmodel.NewDouble(3)},
	}
	if err := mgr.Flush(map[string][]tsmodel.Point{key.String(): points}); err != nil {
		t.Fatal(err)
	}

	got := mgr.Query(key, 10, 30)
	tsOnly := mgr.QueryTimestamps(key, 10, 30)
	if len(got) != len(tsOnly) || len(got) != 3 {
		t.Fatalf("query=%+v tsOnly=%+v", got, tsOnly)
	}
	for i := range got {
		if got[i].Timestamp != tsOnly[i].Timestamp {
			t.Fatalf("ts mismatch at %d: %+v vs %+v", i, got[i], tsOnly[i])
		}
		if tsOnly[i].Value.Type != tsmodel.DataTypeUnknown {
			t.Fatalf("timestamp-only query should not load values: %+v", tsOnly[i])
		}
	}
}

func TestActiveSegmentStaysInMemory(t *testing.T) {
	dir := t.TempDir()
	mgr, err := OpenManager(dir, 100)
	if err != nil {
		t.Fatal(err)
	}
	key := tsmodel.SeriesKey{DevicePath: "d1", Measurement: "s1"}
	if err := mgr.Flush(map[string][]tsmodel.Point{
		key.String(): {{Timestamp: 1, Value: tsmodel.NewDouble(1)}},
	}); err != nil {
		t.Fatal(err)
	}
	if mgr.activeMem == nil || mgr.activeMem.mem == nil {
		t.Fatal("active.seg should stay in memory")
	}
	if mgr.activeMem.mapped != nil {
		t.Fatal("active.seg should not use mmap while appendable")
	}
}
