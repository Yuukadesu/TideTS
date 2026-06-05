package metrics

import (
	"path/filepath"
	"testing"

	"github.com/hanami/tidets/core/storageengine"
	"github.com/prometheus/client_golang/prometheus"
)

func TestStorageCollectorReportsEngineStats(t *testing.T) {
	dir := t.TempDir()
	engine, err := storageengine.OpenWithOptions(storageengine.Options{
		DataDir: filepath.Join(dir, "data"),
		FlushAt: 64,
	})
	if err != nil {
		t.Fatalf("open engine: %v", err)
	}
	defer func() {
		if err := engine.Close(); err != nil {
			t.Fatalf("close engine: %v", err)
		}
	}()

	key := storageengine.SeriesKey{DevicePath: "root.sg1.d1", Measurement: "temperature"}
	if err := engine.Insert(key, storageengine.DoublePoint(100, 25.5)); err != nil {
		t.Fatalf("insert: %v", err)
	}
	if err := engine.Flush(); err != nil {
		t.Fatalf("flush: %v", err)
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(NewStorageCollector(engine))

	families, err := reg.Gather()
	if err != nil {
		t.Fatalf("gather metrics: %v", err)
	}

	if got := gaugeValue(t, families, "tidets_storage_segment_count"); got < 1 {
		t.Fatalf("segment count = %v, want >= 1", got)
	}
	if got := gaugeValue(t, families, "tidets_storage_memtable_points"); got != 0 {
		t.Fatalf("memtable points = %v, want 0 after flush", got)
	}
	if got := gaugeValue(t, families, "tidets_storage_last_flush_timestamp_seconds"); got <= 0 {
		t.Fatalf("last flush timestamp = %v, want > 0", got)
	}
}

func TestStorageHooksRecordCorePathMetrics(t *testing.T) {
	dir := t.TempDir()
	asyncFlush := false
	engine, err := storageengine.OpenWithOptions(storageengine.Options{
		DataDir:          filepath.Join(dir, "data"),
		AsyncFlush:       &asyncFlush,
		FlushAt:          1,
		SealAfterFlushes: 1,
		CompactThreshold: 2,
		CompactMerge:     2,
	})
	if err != nil {
		t.Fatalf("open engine: %v", err)
	}
	defer func() {
		if err := engine.Close(); err != nil {
			t.Fatalf("close engine: %v", err)
		}
	}()

	reg := NewRegistry()
	engine.SetHooks(reg.StorageHooks())

	key := storageengine.SeriesKey{DevicePath: "root.sg1.d1", Measurement: "temperature"}
	if err := engine.Insert(key, storageengine.DoublePoint(100, 25.5)); err != nil {
		t.Fatalf("insert 1: %v", err)
	}
	if err := engine.Insert(key, storageengine.DoublePoint(101, 26.0)); err != nil {
		t.Fatalf("insert 2: %v", err)
	}
	if err := engine.InsertBatch([]storageengine.Record{
		{Key: key, Point: storageengine.DoublePoint(102, 26.5)},
		{Key: key, Point: storageengine.DoublePoint(103, 27.0)},
	}); err != nil {
		t.Fatalf("insert batch: %v", err)
	}

	points, err := engine.Query(key, 100, 103, 0)
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	if len(points) != 4 {
		t.Fatalf("query len = %d, want 4", len(points))
	}
	n, err := engine.Count(key, 100, 103)
	if err != nil {
		t.Fatalf("count: %v", err)
	}
	if n != 4 {
		t.Fatalf("count = %d, want 4", n)
	}
	deleted, err := engine.DeleteRange(key, 100, 101)
	if err != nil {
		t.Fatalf("delete range: %v", err)
	}
	if deleted != 2 {
		t.Fatalf("deleted = %d, want 2", deleted)
	}
	if err := engine.Flush(); err != nil {
		t.Fatalf("flush before gather: %v", err)
	}

	families, err := reg.reg.Gather()
	if err != nil {
		t.Fatalf("gather metrics: %v", err)
	}

	if got := counterValue(t, families, "tidets_storage_write_requests_total", map[string]string{"op": "insert"}); got != 2 {
		t.Fatalf("insert write requests = %v, want 2", got)
	}
	if got := counterValue(t, families, "tidets_storage_write_points_total", map[string]string{"op": "insert"}); got != 2 {
		t.Fatalf("insert write points = %v, want 2", got)
	}
	if got := counterValue(t, families, "tidets_storage_write_requests_total", map[string]string{"op": "insert_batch"}); got != 1 {
		t.Fatalf("insert_batch write requests = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_storage_write_points_total", map[string]string{"op": "insert_batch"}); got != 2 {
		t.Fatalf("insert_batch write points = %v, want 2", got)
	}
	if got := counterValue(t, families, "tidets_storage_write_requests_total", map[string]string{"op": "delete_range"}); got != 1 {
		t.Fatalf("delete_range write requests = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_storage_write_points_total", map[string]string{"op": "delete_range"}); got != 2 {
		t.Fatalf("delete_range write points = %v, want 2", got)
	}
	if got := counterValue(t, families, "tidets_storage_read_requests_total", map[string]string{"op": "query"}); got != 1 {
		t.Fatalf("query read requests = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_storage_read_points_total", map[string]string{"op": "query"}); got != 4 {
		t.Fatalf("query read points = %v, want 4", got)
	}
	if got := counterValue(t, families, "tidets_storage_read_requests_total", map[string]string{"op": "count"}); got != 1 {
		t.Fatalf("count read requests = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_storage_read_points_total", map[string]string{"op": "count"}); got != 4 {
		t.Fatalf("count read points = %v, want 4", got)
	}
	if got := counterValue(t, families, "tidets_storage_wal_events_total", map[string]string{"op": "append_insert"}); got != 2 {
		t.Fatalf("wal append_insert events = %v, want 2", got)
	}
	if got := counterValue(t, families, "tidets_storage_wal_records_total", map[string]string{"op": "append_insert"}); got != 2 {
		t.Fatalf("wal append_insert records = %v, want 2", got)
	}
	if got := counterValue(t, families, "tidets_storage_wal_events_total", map[string]string{"op": "append_insert_batch"}); got != 1 {
		t.Fatalf("wal append_insert_batch events = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_storage_wal_records_total", map[string]string{"op": "append_insert_batch"}); got != 2 {
		t.Fatalf("wal append_insert_batch records = %v, want 2", got)
	}
	if got := counterValue(t, families, "tidets_storage_wal_events_total", map[string]string{"op": "append_delete_range"}); got != 1 {
		t.Fatalf("wal append_delete_range events = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_storage_tombstone_events_total", map[string]string{"op": "mark"}); got != 1 {
		t.Fatalf("tombstone mark events = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_storage_tombstone_ranges_total", map[string]string{"op": "mark"}); got != 1 {
		t.Fatalf("tombstone mark ranges = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_storage_flush_total", nil); got < 3 {
		t.Fatalf("flush total = %v, want >= 3", got)
	}
	if got := counterValue(t, families, "tidets_storage_flush_points_total", nil); got < 4 {
		t.Fatalf("flush points = %v, want >= 4", got)
	}
	if got := counterValue(t, families, "tidets_storage_compact_total", nil); got < 1 {
		t.Fatalf("compact total = %v, want >= 1", got)
	}
	if got := counterValue(t, families, "tidets_storage_compact_input_files_total", nil); got < 2 {
		t.Fatalf("compact input files = %v, want >= 2", got)
	}
	if got := counterValue(t, families, "tidets_storage_compact_output_files_total", nil); got < 1 {
		t.Fatalf("compact output files = %v, want >= 1", got)
	}
}
