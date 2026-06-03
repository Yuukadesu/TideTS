package segment

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hanami/tidets/core/storageengine/model"
)

func TestSealedSegmentUsesMmap(t *testing.T) {
	dir := t.TempDir()
	mgr, err := OpenManager(dir, 1)
	if err != nil {
		t.Fatal(err)
	}
	key := model.SeriesKey{DevicePath: "d1", Measurement: "s1"}
	if err := mgr.Flush(map[string][]model.Point{
		key.String(): {
			{Timestamp: 1, Value: model.NewDouble(1)},
			{Timestamp: 2, Value: model.NewDouble(2)},
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

func TestActiveSegmentStaysInMemory(t *testing.T) {
	dir := t.TempDir()
	mgr, err := OpenManager(dir, 100)
	if err != nil {
		t.Fatal(err)
	}
	key := model.SeriesKey{DevicePath: "d1", Measurement: "s1"}
	if err := mgr.Flush(map[string][]model.Point{
		key.String(): {{Timestamp: 1, Value: model.NewDouble(1)}},
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
