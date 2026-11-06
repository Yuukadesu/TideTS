package parser_test

import (
	"testing"

	"github.com/hanami/tidets/core/sql/ast"
	sqlparser "github.com/hanami/tidets/core/sql/parser"
	"github.com/hanami/tidets/core/tsmodel"
)

func TestParseInsert(t *testing.T) {
	stmt, err := sqlparser.Parse("INSERT INTO root.sg1.d1(temperature) VALUES (100, 25.5)")
	if err != nil {
		t.Fatal(err)
	}
	ins, ok := stmt.(*ast.InsertStmt)
	if !ok {
		t.Fatalf("want InsertStmt, got %T", stmt)
	}
	if ins.DevicePath != "root.sg1.d1" || ins.Measurement != "temperature" || len(ins.Rows) != 1 || ins.Rows[0].Timestamp != 100 {
		t.Fatalf("insert: %+v", ins)
	}
	if ins.Rows[0].Value.Type != tsmodel.DataTypeDouble || ins.Rows[0].Value.Double != 25.5 {
		t.Fatalf("value: %+v", ins.Rows[0].Value)
	}
}

func TestParseInsertBatch(t *testing.T) {
	stmt, err := sqlparser.Parse("INSERT INTO root.d1(s1) VALUES (1, 1.0), (2, 2.0)")
	if err != nil {
		t.Fatal(err)
	}
	ins, ok := stmt.(*ast.InsertStmt)
	if !ok || len(ins.Rows) != 2 {
		t.Fatalf("batch insert: %+v", ins)
	}
}

func TestParseDelete(t *testing.T) {
	stmt, err := sqlparser.Parse("DELETE FROM root.d1(s1) WHERE time >= 1 AND time <= 3")
	if err != nil {
		t.Fatal(err)
	}
	del, ok := stmt.(*ast.DeleteStmt)
	if !ok || len(del.TimeWhere) != 2 {
		t.Fatalf("delete: %+v", del)
	}
}

func TestParseSelectCount(t *testing.T) {
	stmt, err := sqlparser.Parse("SELECT COUNT(s1) FROM root.d1 WHERE time >= 1 AND time <= 3")
	if err != nil {
		t.Fatal(err)
	}
	sel, ok := stmt.(*ast.SelectStmt)
	if !ok || sel.Aggregate != ast.SelectCount {
		t.Fatalf("count: %+v", sel)
	}
}

func TestParseSelectWhere(t *testing.T) {
	stmt, err := sqlparser.Parse("SELECT temperature FROM root.sg1.d1 WHERE time >= 100 AND time <= 200 LIMIT 5")
	if err != nil {
		t.Fatal(err)
	}
	sel, ok := stmt.(*ast.SelectStmt)
	if !ok {
		t.Fatalf("want SelectStmt, got %T", stmt)
	}
	if len(sel.TimeWhere) != 2 || sel.Limit != 5 {
		t.Fatalf("select: %+v", sel)
	}
}

func TestParseCreateTimeseries(t *testing.T) {
	stmt, err := sqlparser.Parse("CREATE TIMESERIES root.sg1.d1(temperature) WITH DATATYPE=DOUBLE")
	if err != nil {
		t.Fatal(err)
	}
	create, ok := stmt.(*ast.CreateTimeseriesStmt)
	if !ok {
		t.Fatalf("want CreateTimeseriesStmt, got %T", stmt)
	}
	if create.DevicePath != "root.sg1.d1" || create.Measurement != "temperature" {
		t.Fatalf("create: %+v", create)
	}
	if create.DataType != tsmodel.DataTypeDouble {
		t.Fatalf("datatype: %+v", create.DataType)
	}
}

func TestParseShowDevices(t *testing.T) {
	stmt, err := sqlparser.Parse("SHOW DEVICES root.sg1.**")
	if err != nil {
		t.Fatal(err)
	}
	show, ok := stmt.(*ast.ShowDevicesStmt)
	if !ok {
		t.Fatalf("want ShowDevicesStmt, got %T", stmt)
	}
	if show.Pattern != "root.sg1.**" {
		t.Fatalf("pattern: %q", show.Pattern)
	}

	stmt, err = sqlparser.Parse("SHOW DEVICES")
	if err != nil {
		t.Fatal(err)
	}
	show, ok = stmt.(*ast.ShowDevicesStmt)
	if !ok || show.Pattern != "" {
		t.Fatalf("show all: %+v err=%v", show, err)
	}
}

func TestParseShowTimeseries(t *testing.T) {
	stmt, err := sqlparser.Parse("SHOW TIMESERIES root.sg1.d1")
	if err != nil {
		t.Fatal(err)
	}
	show, ok := stmt.(*ast.ShowTimeseriesStmt)
	if !ok {
		t.Fatalf("want ShowTimeseriesStmt, got %T", stmt)
	}
	if show.DevicePath != "root.sg1.d1" {
		t.Fatalf("device: %q", show.DevicePath)
	}
}

func TestParseAccountStatements(t *testing.T) {
	stmt, err := sqlparser.Parse("CREATE USER alice 'secret'")
	if err != nil {
		t.Fatal(err)
	}
	created, ok := stmt.(*ast.CreateUserStmt)
	if !ok || created.Username != "alice" || created.Password != "secret" {
		t.Fatalf("create user: %+v", stmt)
	}

	stmt, err = sqlparser.Parse("ALTER USER alice SET PASSWORD 'secret2'")
	if err != nil {
		t.Fatal(err)
	}
	altered, ok := stmt.(*ast.AlterUserStmt)
	if !ok || altered.Username != "alice" || altered.Password != "secret2" {
		t.Fatalf("alter user: %+v", stmt)
	}

	stmt, err = sqlparser.Parse("GRANT READ, WRITE ON root.sg1.** TO alice")
	if err != nil {
		t.Fatal(err)
	}
	grant, ok := stmt.(*ast.GrantStmt)
	if !ok || grant.Revoke || grant.Username != "alice" || grant.Path != "root.sg1.**" || len(grant.Privs) != 2 {
		t.Fatalf("grant: %+v", stmt)
	}

	stmt, err = sqlparser.Parse("REVOKE WRITE ON root.sg1.* FROM alice")
	if err != nil {
		t.Fatal(err)
	}
	revoke, ok := stmt.(*ast.GrantStmt)
	if !ok || !revoke.Revoke || revoke.Path != "root.sg1.*" || len(revoke.Privs) != 1 || revoke.Privs[0] != ast.AccountWrite {
		t.Fatalf("revoke: %+v", stmt)
	}
}
