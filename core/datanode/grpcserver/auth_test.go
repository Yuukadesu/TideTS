package grpcserver

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/hanami/tidets/core/storageengine"
	"github.com/hanami/tidets/core/storageengine/wal"
	pb "github.com/hanami/tidets/protocol/grpc-datanode/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAccountSQLAndDeniedWrite(t *testing.T) {
	dir := t.TempDir()
	engine, err := storageengine.Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	srv, err := New(engine)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = srv.Close()
		_ = engine.Close()
	})

	ctx := context.Background()
	rootID := login(t, srv, "root", "root")
	if _, err := srv.ExecuteSQL(ctx, &pb.ExecuteSQLRequest{
		SessionId: rootID,
		Sql:       "CREATE USER alice 'secret'",
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := srv.ExecuteSQL(ctx, &pb.ExecuteSQLRequest{
		SessionId: rootID,
		Sql:       "GRANT READ ON root.sg1.** TO alice",
	}); err != nil {
		t.Fatal(err)
	}

	aliceID := login(t, srv, "alice", "secret")
	logBefore := fileSize(t, filepath.Join(dir, "system", "auth", "authlog.bin"))
	_, err = srv.ExecuteSQL(ctx, &pb.ExecuteSQLRequest{
		SessionId: aliceID,
		Sql:       "CREATE USER bob 'secret'",
	})
	if status.Code(err) != codes.PermissionDenied {
		t.Fatalf("alice create user: %v", err)
	}
	if fileSize(t, filepath.Join(dir, "system", "auth", "authlog.bin")) != logBefore {
		t.Fatal("denied account statement appended auth log")
	}

	walBefore, err := wal.FileSize(dir)
	if err != nil {
		t.Fatal(err)
	}
	schemaBefore := fileSize(t, filepath.Join(dir, "system", "schema", "mlog.bin"))
	_, err = srv.InsertPoint(ctx, &pb.InsertPointRequest{
		SessionId:   aliceID,
		DevicePath:  "root.sg1.d1",
		Measurement: "temperature",
		Timestamp:   1,
		Value: &pb.TSValue{
			DataType: pb.TSDataType_TS_DATA_TYPE_DOUBLE,
			Value:    &pb.TSValue_DoubleValue{DoubleValue: 1},
		},
	})
	if status.Code(err) != codes.PermissionDenied {
		t.Fatalf("alice insert: %v", err)
	}
	walAfter, err := wal.FileSize(dir)
	if err != nil {
		t.Fatal(err)
	}
	if walAfter != walBefore {
		t.Fatalf("denied insert wrote wal: before=%d after=%d", walBefore, walAfter)
	}
	if fileSize(t, filepath.Join(dir, "system", "schema", "mlog.bin")) != schemaBefore {
		t.Fatal("denied insert wrote schema log")
	}

	if _, err := srv.ExecuteSQL(ctx, &pb.ExecuteSQLRequest{
		SessionId: rootID,
		Sql:       "GRANT WRITE ON root.sg1.** TO alice",
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := srv.InsertPoint(ctx, &pb.InsertPointRequest{
		SessionId:   aliceID,
		DevicePath:  "root.sg1.d1",
		Measurement: "temperature",
		Timestamp:   1,
		Value: &pb.TSValue{
			DataType: pb.TSDataType_TS_DATA_TYPE_DOUBLE,
			Value:    &pb.TSValue_DoubleValue{DoubleValue: 1},
		},
	}); err != nil {
		t.Fatal(err)
	}

	_, err = srv.OpenSession(ctx, &pb.OpenSessionRequest{Username: "alice", Password: "nope"})
	if status.Code(err) != codes.Unauthenticated {
		t.Fatalf("bad password: %v", err)
	}
	_, err = srv.OpenSession(ctx, &pb.OpenSessionRequest{Username: "missing", Password: "nope"})
	if status.Code(err) != codes.Unauthenticated {
		t.Fatalf("missing user: %v", err)
	}
}

func login(t *testing.T, srv *Server, username, password string) int64 {
	t.Helper()
	resp, err := srv.OpenSession(context.Background(), &pb.OpenSessionRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.GetSessionId() <= 0 {
		t.Fatalf("session id %d", resp.GetSessionId())
	}
	return resp.GetSessionId()
}

func fileSize(t *testing.T, path string) int64 {
	t.Helper()
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return 0
		}
		t.Fatal(err)
	}
	return fi.Size()
}
