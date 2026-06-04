package queryengine_test

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	commons "github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/datanode/metadata"
	"github.com/hanami/tidets/core/dataplane"
	"github.com/hanami/tidets/core/queryengine"
	"github.com/hanami/tidets/core/queryengine/backend"
	"github.com/hanami/tidets/core/queryengine/result"
	"github.com/hanami/tidets/core/schemaengine"
	"github.com/hanami/tidets/core/storageengine"
	"github.com/hanami/tidets/core/tsmodel"
)

func openTestService(t *testing.T) (*queryengine.Service, func()) {
	t.Helper()
	dir := t.TempDir()
	eng, err := storageengine.Open(filepath.Join(dir, "data"))
	if err != nil {
		t.Fatal(err)
	}
	schemaSvc, err := schemaengine.Open(filepath.Join(dir, "data"))
	if err != nil {
		_ = eng.Close()
		t.Fatal(err)
	}
	metaMgr := metadata.New(schemaSvc)
	gw := dataplane.New(eng, schemaSvc)
	catalog := &backend.MetadataCatalog{Meta: metaMgr}
	return queryengine.NewService(eng, gw, catalog), func() {
		_ = schemaSvc.Close()
		_ = eng.Close()
	}
}

func TestServiceInsertSelectSQL(t *testing.T) {
	svc, cleanup := openTestService(t)
	defer cleanup()
	ctx := context.Background()

	insertSQL := "INSERT INTO root.sg1.d1(temperature) VALUES (100, 25.5)"
	res, err := svc.Execute(ctx, insertSQL)
	if err != nil {
		t.Fatal(err)
	}
	if res.Kind != result.KindInsert || res.AffectedRows != 1 {
		t.Fatalf("insert result: %+v", res)
	}

	selectSQL := "SELECT temperature FROM root.sg1.d1 WHERE time >= 100 AND time <= 100"
	res, err = svc.Execute(ctx, selectSQL)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Rows) != 1 {
		t.Fatalf("select rows: %+v", res.Rows)
	}
	if res.Rows[0].Timestamp != 100 || res.Rows[0].Value.Type != tsmodel.DataTypeDouble {
		t.Fatalf("row: %+v", res.Rows[0])
	}
	if res.Rows[0].Value.Double != 25.5 {
		t.Fatalf("value: %+v", res.Rows[0].Value)
	}
}

func TestServiceInsertUpsert(t *testing.T) {
	svc, cleanup := openTestService(t)
	defer cleanup()
	ctx := context.Background()

	_, err := svc.Execute(ctx, "INSERT INTO root.d1(s1) VALUES (1, 1.0)")
	if err != nil {
		t.Fatal(err)
	}
	_, err = svc.Execute(ctx, "INSERT INTO root.d1(s1) VALUES (1, 2.0)")
	if err != nil {
		t.Fatal(err)
	}

	res, err := svc.Execute(ctx, "SELECT s1 FROM root.d1 WHERE time = 1")
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Rows) != 1 || res.Rows[0].Value.Double != 2.0 {
		t.Fatalf("upsert: %+v", res.Rows)
	}
}

func TestServiceQueryRange(t *testing.T) {
	svc, cleanup := openTestService(t)
	defer cleanup()
	ctx := context.Background()

	_, err := svc.Execute(ctx, "INSERT INTO root.d1(s1) VALUES (1, 1.0)")
	if err != nil {
		t.Fatal(err)
	}
	_, err = svc.Execute(ctx, "INSERT INTO root.d1(s1) VALUES (2, 2.0)")
	if err != nil {
		t.Fatal(err)
	}

	res, err := svc.QueryRange(ctx, tsmodel.SeriesKey{DevicePath: "root.d1", Measurement: "s1"}, 1, 2, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Rows) != 2 {
		t.Fatalf("query range rows: %+v", res.Rows)
	}
}

func TestServiceCreateTimeseries(t *testing.T) {
	svc, cleanup := openTestService(t)
	defer cleanup()
	ctx := context.Background()

	res, err := svc.Execute(ctx, "CREATE TIMESERIES root.d1(s1) WITH DATATYPE=DOUBLE")
	if err != nil {
		t.Fatal(err)
	}
	if res.Kind != result.KindCreateTimeseries || res.AffectedRows != 1 {
		t.Fatalf("create: %+v", res)
	}

	_, err = svc.Execute(ctx, "CREATE TIMESERIES root.d1(s1) WITH DATATYPE=DOUBLE")
	if !errors.Is(err, commons.ErrSchemaTimeseriesExists) {
		t.Fatalf("duplicate create: %v", err)
	}

	_, err = svc.Execute(ctx, "INSERT INTO root.d1(s1) VALUES (1, 1.0)")
	if err != nil {
		t.Fatal(err)
	}
	_, err = svc.Execute(ctx, "INSERT INTO root.d1(s1) VALUES (2, 1)")
	if !errors.Is(err, commons.ErrSchemaDataTypeMismatch) {
		t.Fatalf("type mismatch: %v", err)
	}
}

func TestServiceShowDevicesAndTimeseries(t *testing.T) {
	svc, cleanup := openTestService(t)
	defer cleanup()
	ctx := context.Background()

	for _, sql := range []string{
		"CREATE TIMESERIES root.sg1.d1(s1) WITH DATATYPE=DOUBLE",
		"CREATE TIMESERIES root.sg1.d2(s1) WITH DATATYPE=FLOAT",
		"CREATE TIMESERIES root.sg2.d1(s1) WITH DATATYPE=INT64",
	} {
		if _, err := svc.Execute(ctx, sql); err != nil {
			t.Fatal(err)
		}
	}

	res, err := svc.Execute(ctx, "SHOW DEVICES root.sg1.**")
	if err != nil {
		t.Fatal(err)
	}
	if res.Kind != result.KindShowDevices || len(res.CatalogRows) != 2 {
		t.Fatalf("show devices: %+v", res)
	}

	res, err = svc.Execute(ctx, "SHOW TIMESERIES root.sg1.d1")
	if err != nil {
		t.Fatal(err)
	}
	if res.Kind != result.KindShowTimeseries || len(res.CatalogRows) != 1 {
		t.Fatalf("show timeseries: %+v", res)
	}
	if res.CatalogRows[0].Columns["DataType"] != "DOUBLE" {
		t.Fatalf("datatype column: %+v", res.CatalogRows[0])
	}
}

func TestServiceBatchInsertCountDelete(t *testing.T) {
	svc, cleanup := openTestService(t)
	defer cleanup()
	ctx := context.Background()

	res, err := svc.Execute(ctx, "INSERT INTO root.d1(s1) VALUES (1, 1.0), (2, 2.0), (3, 3.0)")
	if err != nil {
		t.Fatal(err)
	}
	if res.AffectedRows != 3 {
		t.Fatalf("batch insert: %+v", res)
	}

	res, err = svc.Execute(ctx, "SELECT COUNT(s1) FROM root.d1 WHERE time >= 1 AND time <= 3")
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Rows) != 1 || res.Rows[0].Value.Int64 != 3 {
		t.Fatalf("count: %+v", res.Rows)
	}

	res, err = svc.Execute(ctx, "DELETE FROM root.d1(s1) WHERE time >= 2 AND time <= 3")
	if err != nil {
		t.Fatal(err)
	}
	if res.AffectedRows != 2 {
		t.Fatalf("delete: %+v", res)
	}

	res, err = svc.Execute(ctx, "SELECT COUNT(s1) FROM root.d1 WHERE time >= 1 AND time <= 3")
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Rows) != 1 || res.Rows[0].Value.Int64 != 1 {
		t.Fatalf("count after delete: %+v", res.Rows)
	}
}
