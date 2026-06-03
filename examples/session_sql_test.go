package examples_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hanami/tidets/client/session"
)

// TestSessionExecuteSQL 端到端：ExecuteSQL INSERT + SELECT。
func TestSessionExecuteSQL(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	host := envOr("TIDETS_HOST", session.DefaultHost)
	port := envPortOr("TIDETS_PORT", session.DefaultPort)
	addr := fmt.Sprintf("%s:%d", host, port)

	if err := pingTCP(addr, 2*time.Second); err != nil {
		t.Skipf("datanode not running at %s: %v", addr, err)
	}

	ctx := context.Background()
	s, err := session.New(session.WithHost(host), session.WithPort(port))
	if err != nil {
		t.Fatal(err)
	}
	if err := s.Open(ctx); err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	if s.SessionID() <= 0 {
		t.Fatal("expected positive session id")
	}

	insertSQL := "INSERT INTO root.sg1.d1(temperature) VALUES (200, 36.5)"
	res, err := s.ExecuteSQL(ctx, insertSQL)
	if err != nil {
		t.Fatal(err)
	}
	if !res.IsInsert() || res.AffectedRows != 1 {
		t.Fatalf("insert: %+v", res)
	}

	selectSQL := "SELECT temperature FROM root.sg1.d1 WHERE time = 200"
	res, err = s.ExecuteSQL(ctx, selectSQL)
	if err != nil {
		t.Fatal(err)
	}
	if !res.IsSelect() || len(res.Rows) != 1 {
		t.Fatalf("select: %+v", res)
	}
	if res.Rows[0].Value.Type != session.DataTypeDouble || res.Rows[0].Value.Double != 36.5 {
		t.Fatalf("row: %+v", res.Rows[0])
	}
}
