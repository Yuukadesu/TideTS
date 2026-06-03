package storageengine

import (
	"testing"
)

func TestEngineCompaction(t *testing.T) {
	dir := t.TempDir()
	syncFlush := false
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}

	e, err := OpenWithOptions(Options{
		DataDir:          dir,
		FlushAt:          1,
		AsyncFlush:       &syncFlush,
		SealAfterFlushes: 1,
		CompactThreshold: 3,
		CompactMerge:     2,
	})
	if err != nil {
		t.Fatal(err)
	}

	for i := int64(1); i <= 5; i++ {
		if err := e.Insert(key, DoublePoint(i*10, float64(i))); err != nil {
			t.Fatal(err)
		}
	}

	sealed := e.segments.SealedFileCount()
	if sealed >= 5 {
		t.Fatalf("expected auto compact to reduce files, sealed=%d", sealed)
	}

	got, err := e.Query(key, 10, 50, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 5 {
		t.Fatalf("query after compact: %+v", got)
	}

	if err := e.Close(); err != nil {
		t.Fatal(err)
	}

	e2, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer e2.Close()

	got2, err := e2.Query(key, 10, 50, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got2) != 5 {
		t.Fatalf("query after reopen: %+v", got2)
	}
}

func TestEngineManualCompact(t *testing.T) {
	dir := t.TempDir()
	syncFlush := false
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}

	e, err := OpenWithOptions(Options{
		DataDir:          dir,
		FlushAt:          1,
		AsyncFlush:       &syncFlush,
		SealAfterFlushes: 1,
		CompactThreshold: 100, // 禁用自动
	})
	if err != nil {
		t.Fatal(err)
	}

	for i := int64(1); i <= 4; i++ {
		if err := e.Insert(key, DoublePoint(i, float64(i))); err != nil {
			t.Fatal(err)
		}
	}
	if e.segments.SealedFileCount() != 4 {
		t.Fatalf("want 4 sealed, got %d", e.segments.SealedFileCount())
	}

	if err := e.Compact(); err != nil {
		t.Fatal(err)
	}
	if e.segments.SealedFileCount() != 3 {
		t.Fatalf("after compact want 3 sealed, got %d", e.segments.SealedFileCount())
	}

	if err := e.Compact(); err != nil {
		t.Fatal(err)
	}
	if e.segments.SealedFileCount() != 2 {
		t.Fatalf("after second compact want 2 sealed, got %d", e.segments.SealedFileCount())
	}

	got, err := e.Query(key, 1, 4, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 4 {
		t.Fatalf("points: %+v", got)
	}
	_ = e.Close()
}
