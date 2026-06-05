package metrics

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
	"github.com/hanami/tidets/core/schemaengine"
	"github.com/hanami/tidets/core/storageengine"
)

func TestSQLOutcomeMetricsRecordSuccessParseAndExecuteErrors(t *testing.T) {
	dir := t.TempDir()
	engine, err := storageengine.Open(filepath.Join(dir, "data"))
	if err != nil {
		t.Fatalf("open engine: %v", err)
	}
	defer func() {
		if err := engine.Close(); err != nil {
			t.Fatalf("close engine: %v", err)
		}
	}()
	schemaSvc, err := schemaengine.Open(filepath.Join(dir, "data"))
	if err != nil {
		t.Fatalf("open schema: %v", err)
	}
	defer func() {
		if err := schemaSvc.Close(); err != nil {
			t.Fatalf("close schema: %v", err)
		}
	}()

	reg := NewRegistry()
	svc := queryengine.NewService(engine, dataplane.New(engine, schemaSvc), &backend.MetadataCatalog{Meta: metadata.New(schemaSvc)})
	svc.SetHooks(queryengine.Hooks{OnPlanExecuted: reg.ObserveSQL})

	ctx := context.Background()
	if _, err := svc.Execute(ctx, "INSERT INTO root.sg1.d1(temperature) VALUES (1, 25.5)"); err != nil {
		t.Fatalf("execute success sql: %v", err)
	}
	if _, err := svc.Execute(ctx, "SELECT FROM"); err == nil {
		t.Fatal("parse error = nil, want syntax error")
	} else if e, ok := commons.As(err); !ok || e.Op != "sql" {
		t.Fatalf("parse error = %v, want sql domain error", err)
	}
	if _, err := svc.Execute(ctx, "CREATE TIMESERIES root.sg1.d1(humidity) WITH DATATYPE=DOUBLE"); err != nil {
		t.Fatalf("create timeseries: %v", err)
	}
	if _, err := svc.Execute(ctx, "CREATE TIMESERIES root.sg1.d1(humidity) WITH DATATYPE=DOUBLE"); !errors.Is(err, commons.ErrSchemaTimeseriesExists) {
		t.Fatalf("execute error = %v, want schema exists", err)
	}

	families, err := reg.reg.Gather()
	if err != nil {
		t.Fatalf("gather metrics: %v", err)
	}

	if got := counterValue(t, families, "tidets_sql_requests_total", map[string]string{"kind": "insert", "success": "true", "error_class": "none"}); got != 1 {
		t.Fatalf("sql insert success counter = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_sql_requests_total", map[string]string{"kind": "unknown", "success": "false", "error_class": "parse"}); got != 1 {
		t.Fatalf("sql parse counter = %v, want 1", got)
	}
	if got := counterValue(t, families, "tidets_sql_requests_total", map[string]string{"kind": "create_timeseries", "success": "false", "error_class": "execute"}); got != 1 {
		t.Fatalf("sql execute counter = %v, want 1", got)
	}
}
