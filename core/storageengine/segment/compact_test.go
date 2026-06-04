package segment

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hanami/tidets/core/tsmodel"
)

func TestCompactReducesFiles(t *testing.T) {
	dir := t.TempDir()
	mgr, err := OpenManagerWithCompact(dir, 1, CompactOptions{Threshold: 4, MergeCount: 2})
	if err != nil {
		t.Fatal(err)
	}

	key := tsmodel.SeriesKey{DevicePath: "d1", Measurement: "s1"}
	for i := int64(1); i <= 3; i++ {
		if err := mgr.Flush(map[string][]tsmodel.Point{
			key.String(): {{Timestamp: i, Value: tsmodel.NewDouble(float64(i))}},
		}); err != nil {
			t.Fatal(err)
		}
	}

	if mgr.SealedFileCount() != 3 {
		t.Fatalf("before compact sealed=%d", mgr.SealedFileCount())
	}
	if err := mgr.Compact(); err != nil {
		t.Fatal(err)
	}
	if mgr.SealedFileCount() != 2 {
		t.Fatalf("after compact sealed=%d", mgr.SealedFileCount())
	}

	pts := mgr.Query(key, 1, 3)
	if len(pts) != 3 {
		t.Fatalf("query: %+v", pts)
	}

	entries, _ := os.ReadDir(filepath.Join(dir, SubDir))
	segFiles := 0
	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".seg" && e.Name() != ActiveFileName {
			segFiles++
		}
	}
	if segFiles != 2 {
		t.Fatalf("disk seg files=%d", segFiles)
	}
}
