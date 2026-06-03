package session

import (
	"testing"

	"github.com/hanami/tidets/core/datanode/auth"
)

func TestManagerLoginClose(t *testing.T) {
	mgr := NewManager(auth.DefaultAuthenticator())

	sess, result := mgr.OpenGRPCSession(OpenParams{
		Login: LoginParams{
			Username:      "root",
			Password:      "root",
			ClientVersion: ClientVersionV10,
			FetchSize:     10000,
		},
		ClientAddr: "127.0.0.1",
		ClientPort: 12345,
	})
	if !result.OK {
		t.Fatalf("login: %s", result.Message)
	}
	if result.SessionID <= 0 {
		t.Fatalf("session id: %d", result.SessionID)
	}
	if sess.UserID() != 0 || sess.Username() != "root" {
		t.Fatalf("user: id=%d name=%s", sess.UserID(), sess.Username())
	}

	got, ok := mgr.Get(result.SessionID)
	if !ok || got.ID() != result.SessionID {
		t.Fatal("get session failed")
	}

	if err := mgr.CloseSession(result.SessionID); err != nil {
		t.Fatal(err)
	}
	if _, ok := mgr.Get(result.SessionID); ok {
		t.Fatal("session should be removed")
	}
}

func TestManagerLoginReject(t *testing.T) {
	mgr := NewManager(auth.DefaultAuthenticator())
	_, result := mgr.OpenGRPCSession(OpenParams{
		Login:      LoginParams{Username: "root", Password: "wrong"},
		ClientAddr: "127.0.0.1",
	})
	if result.OK {
		t.Fatal("expected login failure")
	}
	if mgr.Count() != 0 {
		t.Fatalf("sessions=%d want 0", mgr.Count())
	}
}
