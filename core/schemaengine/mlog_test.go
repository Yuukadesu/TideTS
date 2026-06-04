package schemaengine

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hanami/tidets/core/tsmodel"
)

func TestMlogReplay(t *testing.T) {
	dir := t.TempDir()
	schemaDir := schemaDir(dir)
	if err := os.MkdirAll(schemaDir, 0o755); err != nil {
		t.Fatal(err)
	}

	ts := NewTimeseries("root.sg1.d1", "temperature", tsmodel.DataTypeDouble)
	ml, err := openMlog(schemaDir)
	if err != nil {
		t.Fatal(err)
	}
	if err := ml.appendCreate(ts); err != nil {
		t.Fatal(err)
	}
	if err := ml.close(); err != nil {
		t.Fatal(err)
	}

	tree := newMTree()
	if err := replayMlog(schemaDir, 0, tree.put); err != nil {
		t.Fatal(err)
	}
	got, ok := tree.get(ts.Key())
	if !ok || got.DataType != tsmodel.DataTypeDouble {
		t.Fatalf("replay failed: ok=%v got=%+v", ok, got)
	}
}

func TestSnapshotRoundTrip(t *testing.T) {
	dir := t.TempDir()
	schemaDir := schemaDir(dir)
	meta := snapshotMeta{
		MlogOffset: 42,
		Series: []Timeseries{
			NewTimeseries("root.a.d1", "s1", tsmodel.DataTypeInt32),
			NewTimeseries("root.a.d1", "s2", tsmodel.DataTypeFloat),
		},
	}
	if err := saveSnapshot(schemaDir, meta); err != nil {
		t.Fatal(err)
	}
	loaded, err := loadSnapshot(schemaDir)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.MlogOffset != 42 || len(loaded.Series) != 2 {
		t.Fatalf("loaded=%+v", loaded)
	}
	tree := newMTreeFromSeries(loaded.Series)
	if len(tree.ListDevices("")) != 1 {
		t.Fatalf("devices=%v", tree.ListDevices(""))
	}
	if len(tree.ListMeasurements("root.a.d1")) != 2 {
		t.Fatalf("measurements=%v", tree.ListMeasurements("root.a.d1"))
	}
	if _, err := os.Stat(filepath.Join(schemaDir, snapshotFileName)); err != nil {
		t.Fatal(err)
	}
}
