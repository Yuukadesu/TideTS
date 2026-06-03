package session

import "testing"

func TestSQLResultKind(t *testing.T) {
	ins := &SQLResult{AffectedRows: 1}
	if !ins.IsInsert() || ins.IsSelect() {
		t.Fatalf("insert: %+v", ins)
	}

	sel := &SQLResult{Rows: []Point{{Timestamp: 1, Value: Double(1)}}}
	if !sel.IsSelect() || sel.IsInsert() {
		t.Fatalf("select: %+v", sel)
	}
}
