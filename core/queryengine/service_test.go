package queryengine_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/hanami/tidets/core/queryengine"
	"github.com/hanami/tidets/core/queryengine/result"
	querystore "github.com/hanami/tidets/core/queryengine/storage"
	"github.com/hanami/tidets/core/storageengine"
	"github.com/hanami/tidets/core/storageengine/model"
)

func TestServiceInsertSelectSQL(t *testing.T) {
	dir := t.TempDir()
	eng, err := storageengine.Open(filepath.Join(dir))
	if err != nil {
		t.Fatal(err)
	}
	defer eng.Close()

	svc := queryengine.NewService(&querystore.EngineAdapter{Engine: eng})
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
	if res.Rows[0].Timestamp != 100 || res.Rows[0].Value.Type != model.DataTypeDouble {
		t.Fatalf("row: %+v", res.Rows[0])
	}
	if res.Rows[0].Value.Double != 25.5 {
		t.Fatalf("value: %+v", res.Rows[0].Value)
	}
}

func TestServiceInsertUpsert(t *testing.T) {
	dir := t.TempDir()
	eng, err := storageengine.Open(filepath.Join(dir))
	if err != nil {
		t.Fatal(err)
	}
	defer eng.Close()

	svc := queryengine.NewService(&querystore.EngineAdapter{Engine: eng})
	ctx := context.Background()

	_, err = svc.Execute(ctx, "INSERT INTO root.d1(s1) VALUES (1, 1.0)")
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
