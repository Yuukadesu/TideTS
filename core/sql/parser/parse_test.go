package parser_test

import (
	"testing"

	"github.com/hanami/tidets/core/sql/ast"
	sqlparser "github.com/hanami/tidets/core/sql/parser"
	"github.com/hanami/tidets/core/storageengine/model"
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
	if ins.DevicePath != "root.sg1.d1" || ins.Measurement != "temperature" || ins.Timestamp != 100 {
		t.Fatalf("insert: %+v", ins)
	}
	if ins.Value.Type != model.DataTypeDouble || ins.Value.Double != 25.5 {
		t.Fatalf("value: %+v", ins.Value)
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
