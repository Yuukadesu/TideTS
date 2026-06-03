package repl_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/hanami/tidets/cli/repl"
	"github.com/hanami/tidets/client/session"
)

type fakeSession struct {
	open bool
	sql  []string
}

func (f *fakeSession) Open(context.Context) error { f.open = true; return nil }
func (f *fakeSession) Close() error               { f.open = false; return nil }
func (f *fakeSession) IsOpen() bool               { return f.open }
func (f *fakeSession) SessionID() int64           { return 1 }
func (f *fakeSession) InsertPoint(context.Context, string, string, int64, session.Value) error {
	return nil
}
func (f *fakeSession) InsertBatch(context.Context, string, []session.BatchPoint) error {
	return nil
}
func (f *fakeSession) QueryRange(context.Context, string, string, int64, int64) ([]session.Point, error) {
	return nil, nil
}
func (f *fakeSession) QueryRangeWithLimit(context.Context, string, string, int64, int64, int) ([]session.Point, error) {
	return nil, nil
}
func (f *fakeSession) ExecuteSQL(_ context.Context, sql string) (*session.SQLResult, error) {
	f.sql = append(f.sql, sql)
	return &session.SQLResult{AffectedRows: 1}, nil
}
func (f *fakeSession) SetFetchSize(int) session.Session { return f }
func (f *fakeSession) Host() string                     { return "h" }
func (f *fakeSession) Port() int                        { return 1 }
func (f *fakeSession) Username() string                 { return "u" }
func (f *fakeSession) Password() string                 { return "p" }
func (f *fakeSession) FetchSize() int64                 { return 0 }
func (f *fakeSession) Version() session.Version         { return session.V_1_0 }

func TestRunExecutesSQLUntilExit(t *testing.T) {
	in := strings.NewReader("INSERT INTO root.d1(s1) VALUES (1, 1.0);\nexit\n")
	var out bytes.Buffer
	fs := &fakeSession{}
	if err := repl.Run(context.Background(), fs, repl.Options{In: in, Out: &out}); err != nil {
		t.Fatal(err)
	}
	if len(fs.sql) != 1 || fs.sql[0] != "INSERT INTO root.d1(s1) VALUES (1, 1.0)" {
		t.Fatalf("sql calls: %v", fs.sql)
	}
}
