package storageengine

import (
	"path/filepath"
	"testing"

	"github.com/hanami/tidets/core/storageengine/model"
	"github.com/hanami/tidets/core/storageengine/wal"
)

func testOptions(dir string, flushAt int) Options {
	syncFlush := false
	return Options{
		DataDir:          dir,
		FlushAt:          flushAt,
		AsyncFlush:       &syncFlush,
		SealAfterFlushes: 100,
	}
}

func TestEngineWALTruncateOnIdle(t *testing.T) {
	dir := t.TempDir()
	opts := testOptions(dir, 2)
	opts.WALSync = wal.SyncOnFlush
	opts.WALTruncate = true

	e, err := OpenWithOptions(opts)
	if err != nil {
		t.Fatal(err)
	}
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}
	if err := e.Insert(key, DoublePoint(1, 1)); err != nil {
		t.Fatal(err)
	}
	if err := e.Insert(key, DoublePoint(2, 2)); err != nil {
		t.Fatal(err)
	}
	// flushAt=2 会触发 flush；同步 flush 下，此时应已空闲并触发 truncate/reset 逻辑。
	_ = e.Close()

	// wal.log 应被截断为 0 或接近 0（不同文件系统可能保留空文件）。
	if sz, err := wal.FileSize(dir); err != nil {
		t.Fatal(err)
	} else if sz != 0 {
		t.Fatalf("wal.log should be truncated to 0, got %d", sz)
	}
}

func openTest(t *testing.T, dir string, flushAt int) *Engine {
	t.Helper()
	e, err := OpenWithOptions(testOptions(dir, flushAt))
	if err != nil {
		t.Fatal(err)
	}
	return e
}

func TestEngineInsertQueryReplay(t *testing.T) {
	dir := t.TempDir()

	e1 := openTest(t, dir, 4096)
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "temperature"}
	if err := e1.Insert(key, DoublePoint(100, 25.5)); err != nil {
		t.Fatal(err)
	}
	if err := e1.Insert(key, DoublePoint(200, 26.0)); err != nil {
		t.Fatal(err)
	}
	if err := e1.Close(); err != nil {
		t.Fatal(err)
	}

	e2, err := Open(filepath.Join(dir))
	if err != nil {
		t.Fatal(err)
	}
	defer e2.Close()

	got, err := e2.Query(key, 100, 200, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || !got[0].Value.Equal(model.NewDouble(25.5)) || got[1].Timestamp != 200 {
		t.Fatalf("unexpected points: %+v", got)
	}
}

func TestEngineFlushToSegment(t *testing.T) {
	dir := t.TempDir()
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}

	e := openTest(t, dir, 2)
	if err := e.Insert(key, DoublePoint(1, 1)); err != nil {
		t.Fatal(err)
	}
	if err := e.Insert(key, DoublePoint(2, 2)); err != nil {
		t.Fatal(err)
	}
	if e.MemTablePointCount() != 0 {
		t.Fatalf("memtable should be empty after flush, got %d points", e.MemTablePointCount())
	}
	if e.StableTime(key.DevicePath) != 2 {
		t.Fatalf("stable time want 2, got %d", e.StableTime(key.DevicePath))
	}

	got, err := e.Query(key, 1, 2, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("expected 2 points from segment, got %+v", got)
	}

	if err := e.Insert(key, DoublePoint(3, 3)); err != nil {
		t.Fatal(err)
	}
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}

	e2, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer e2.Close()

	all, err := e2.Query(key, 1, 3, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 3 {
		t.Fatalf("expected 3 points after reopen, got %+v", all)
	}
	if e2.StableTime(key.DevicePath) != 3 {
		t.Fatalf("stable time after reopen want 3, got %d", e2.StableTime(key.DevicePath))
	}
}

func TestEngineInsertBatch(t *testing.T) {
	dir := t.TempDir()
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}

	e := openTest(t, dir, 4096)
	defer e.Close()

	err := e.InsertBatch([]Record{
		{Key: key, Point: DoublePoint(10, 1)},
		{Key: key, Point: DoublePoint(20, 2)},
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := e.Query(key, 10, 20, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("got %+v", got)
	}
}

func TestEngineOutOfOrderDelayedMemTable(t *testing.T) {
	dir := t.TempDir()
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "temp"}

	e := openTest(t, dir, 2)

	if err := e.Insert(key, DoublePoint(100, 1)); err != nil {
		t.Fatal(err)
	}
	if err := e.Insert(key, DoublePoint(200, 2)); err != nil {
		t.Fatal(err)
	}
	if e.MemTablePointCount() != 0 {
		t.Fatal("expected flush")
	}

	if err := e.Insert(key, DoublePoint(150, 1.5)); err != nil {
		t.Fatal(err)
	}
	if e.DelayedMemTablePointCount() != 1 {
		t.Fatalf("delayed memtable want 1 point, got %d", e.DelayedMemTablePointCount())
	}

	got, err := e.Query(key, 100, 200, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 {
		t.Fatalf("expected 3 points (2 disk + 1 delayed), got %+v", got)
	}

	if err := e.Close(); err != nil {
		t.Fatal(err)
	}

	e2, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer e2.Close()

	got2, err := e2.Query(key, 100, 200, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got2) != 3 {
		t.Fatalf("after reopen: %+v", got2)
	}
}

func TestEngineQueryLimit(t *testing.T) {
	dir := t.TempDir()
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}
	e := openTest(t, dir, 4096)
	defer e.Close()

	for i := int64(1); i <= 5; i++ {
		if err := e.Insert(key, DoublePoint(i*10, float64(i))); err != nil {
			t.Fatal(err)
		}
	}
	got, err := e.Query(key, 10, 50, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].Timestamp != 10 || got[1].Timestamp != 20 {
		t.Fatalf("limit: %+v", got)
	}
}

func TestEngineStats(t *testing.T) {
	dir := t.TempDir()
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}
	e := openTest(t, dir, 4096)
	defer e.Close()

	if err := e.Insert(key, DoublePoint(1, 1)); err != nil {
		t.Fatal(err)
	}
	st := e.Stats()
	if st.MemTablePoints != 1 {
		t.Fatalf("stats mem points: %+v", st)
	}
	if st.DataDir == "" {
		t.Fatal("stats data dir empty")
	}
}

func TestEngineCrashRecovery(t *testing.T) {
	dir := t.TempDir()
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}

	e1 := openTest(t, dir, 4096)
	if err := e1.Insert(key, DoublePoint(10, 1)); err != nil {
		t.Fatal(err)
	}
	if err := e1.Insert(key, DoublePoint(20, 2)); err != nil {
		t.Fatal(err)
	}
	// 模拟进程崩溃：不 Close，直接丢弃引擎（WAL 未清空）
	e1.mu.Lock()
	_ = e1.wal.Sync()
	e1.mu.Unlock()

	e2, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer e2.Close()

	got, err := e2.Query(key, 10, 20, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("recovery: %+v", got)
	}
}

func TestEngineAsyncFlush(t *testing.T) {
	dir := t.TempDir()
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}
	async := true
	e, err := OpenWithOptions(Options{DataDir: dir, FlushAt: 2, AsyncFlush: &async, SealAfterFlushes: 100})
	if err != nil {
		t.Fatal(err)
	}

	if err := e.Insert(key, DoublePoint(1, 1)); err != nil {
		t.Fatal(err)
	}
	if err := e.Insert(key, DoublePoint(2, 2)); err != nil {
		t.Fatal(err)
	}
	// MemTable 应已 swap，查询仍能看到 pending flush 数据
	got, err := e.Query(key, 1, 2, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("query during async flush: %+v", got)
	}
	e.WaitFlushed()
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestEngineWALSyncOnFlushMode(t *testing.T) {
	dir := t.TempDir()
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}
	syncFlush := false
	e, err := OpenWithOptions(Options{
		DataDir:    dir,
		FlushAt:    2,
		AsyncFlush: &syncFlush,
		WALSync:    wal.SyncOnFlush,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := e.Insert(key, DoublePoint(1, 1)); err != nil {
		t.Fatal(err)
	}
	if err := e.Insert(key, DoublePoint(2, 2)); err != nil {
		t.Fatal(err)
	}
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestEngineUpsertSameTimestamp(t *testing.T) {
	dir := t.TempDir()
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}

	e := openTest(t, dir, 4096)
	defer e.Close()

	if err := e.Insert(key, DoublePoint(5, 1)); err != nil {
		t.Fatal(err)
	}
	if err := e.Insert(key, DoublePoint(5, 9)); err != nil {
		t.Fatal(err)
	}

	got, err := e.Query(key, 5, 5, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || !got[0].Value.Equal(model.NewDouble(9)) {
		t.Fatalf("upsert: %+v", got)
	}
}

func TestEngineTypedValues(t *testing.T) {
	dir := t.TempDir()
	intKey := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "count"}
	textKey := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "status"}

	e := openTest(t, dir, 4096)
	defer e.Close()

	if err := e.Insert(intKey, Point{Timestamp: 1, Value: model.NewInt32(42)}); err != nil {
		t.Fatal(err)
	}
	if err := e.Insert(textKey, Point{Timestamp: 2, Value: model.NewText("ok")}); err != nil {
		t.Fatal(err)
	}
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}

	e2, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer e2.Close()

	gotInt, err := e2.Query(intKey, 1, 1, 0)
	if err != nil || len(gotInt) != 1 || !gotInt[0].Value.Equal(model.NewInt32(42)) {
		t.Fatalf("int query: %+v err=%v", gotInt, err)
	}
	gotText, err := e2.Query(textKey, 2, 2, 0)
	if err != nil || len(gotText) != 1 || !gotText[0].Value.Equal(model.NewText("ok")) {
		t.Fatalf("text query: %+v err=%v", gotText, err)
	}
}

func TestEngineDataTypeMismatch(t *testing.T) {
	dir := t.TempDir()
	key := SeriesKey{DevicePath: "root.sg1.d1", Measurement: "s1"}
	e := openTest(t, dir, 4096)
	defer e.Close()

	if err := e.Insert(key, DoublePoint(1, 1)); err != nil {
		t.Fatal(err)
	}
	if err := e.Insert(key, Point{Timestamp: 2, Value: model.NewInt32(2)}); err == nil {
		t.Fatal("expected data type mismatch")
	}
}
