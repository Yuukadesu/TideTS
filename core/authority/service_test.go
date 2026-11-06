package authority

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	commons "github.com/hanami/tidets/commons/errors"
)

func TestOpenSeedsRoot(t *testing.T) {
	dir := t.TempDir()
	s, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	id, ok, err := s.Authenticate(rootUsername, rootPassword)
	if err != nil || !ok || id != 1 {
		t.Fatalf("login id=%d ok=%v err=%v", id, ok, err)
	}
	if err := s.CheckPrivilege(id, "root.sg1.d1", PrivilegeWrite); err != nil {
		t.Fatal(err)
	}
	if err := s.Close(); err != nil {
		t.Fatal(err)
	}

	s2, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer s2.Close()
	id, ok, err = s2.Authenticate(rootUsername, rootPassword)
	if err != nil || !ok || id != 1 {
		t.Fatalf("reopen id=%d ok=%v err=%v", id, ok, err)
	}
}

func TestAuthenticateSameFailure(t *testing.T) {
	s := openAuth(t)
	if _, ok, err := s.Authenticate("root", "wrong"); err != nil || ok {
		t.Fatalf("wrong password ok=%v err=%v", ok, err)
	}
	if _, ok, err := s.Authenticate("missing", "root"); err != nil || ok {
		t.Fatalf("missing user ok=%v err=%v", ok, err)
	}
	if _, ok, err := s.Authenticate("Root", "root"); err != nil || ok {
		t.Fatalf("case ok=%v err=%v", ok, err)
	}
	if _, ok, err := s.Authenticate("root", ""); err != nil || ok {
		t.Fatalf("empty password ok=%v err=%v", ok, err)
	}
}

func TestCreateGrantRevokeAndReplay(t *testing.T) {
	dir := t.TempDir()
	s, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.CreateUser("alice", "secret"); err != nil {
		t.Fatal(err)
	}
	if err := s.CreateUser("alice", "secret"); !commons.Is(err, commons.ErrAuthUserExists) {
		t.Fatalf("dup: %v", err)
	}
	if err := s.CreateUser("root", "secret"); !commons.Is(err, commons.ErrAuthReservedUser) {
		t.Fatalf("root: %v", err)
	}
	if err := s.Grant("alice", "root.sg1.**", PrivilegeRead, PrivilegeWrite); err != nil {
		t.Fatal(err)
	}
	before := authLogSize(t, dir)
	if err := s.Grant("alice", "root.sg1.**", PrivilegeRead); err != nil {
		t.Fatal(err)
	}
	if authLogSize(t, dir) != before {
		t.Fatal("duplicate grant appended a log record")
	}
	id, ok, err := s.Authenticate("alice", "secret")
	if err != nil || !ok {
		t.Fatalf("alice login ok=%v err=%v", ok, err)
	}
	if err := s.CheckPrivilege(id, "root.sg1.d1", PrivilegeRead); err != nil {
		t.Fatal(err)
	}
	if err := s.CheckPrivilege(id, "root.sg1.d1", PrivilegeWrite); err != nil {
		t.Fatal(err)
	}
	if err := s.CheckPrivilege(id, "root.sg2.d1", PrivilegeRead); !commons.Is(err, commons.ErrAuthPermissionDenied) {
		t.Fatalf("other path: %v", err)
	}
	if err := s.Revoke("alice", "root.sg1.**", PrivilegeWrite); err != nil {
		t.Fatal(err)
	}
	if err := s.CheckPrivilege(id, "root.sg1.d1", PrivilegeWrite); !commons.Is(err, commons.ErrAuthPermissionDenied) {
		t.Fatalf("revoked write: %v", err)
	}
	if err := s.CheckPrivilege(id, "root.sg1.d1", PrivilegeRead); err != nil {
		t.Fatal(err)
	}
	if err := s.Revoke("alice", "root.sg1.**", PrivilegeWrite); !commons.Is(err, commons.ErrAuthGrantNotFound) {
		t.Fatalf("missing revoke: %v", err)
	}
	if err := s.SetPassword("alice", "next"); err != nil {
		t.Fatal(err)
	}
	if err := s.Close(); err != nil {
		t.Fatal(err)
	}

	s2, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer s2.Close()
	id, ok, err = s2.Authenticate("alice", "next")
	if err != nil || !ok {
		t.Fatalf("reopen alice ok=%v err=%v", ok, err)
	}
	if err := s2.CheckPrivilege(id, "root.sg1.d1", PrivilegeRead); err != nil {
		t.Fatal(err)
	}
	if err := s2.CheckPrivilege(id, "root.sg1.d1", PrivilegeWrite); !commons.Is(err, commons.ErrAuthPermissionDenied) {
		t.Fatalf("write survived revoke: %v", err)
	}
}

func TestPasswordTooLong(t *testing.T) {
	s := openAuth(t)
	err := s.CreateUser("alice", strings.Repeat("a", 73))
	if !commons.Is(err, commons.ErrAuthPasswordTooLong) {
		t.Fatalf("got %v", err)
	}
}

func TestCorruptLogDoesNotSeed(t *testing.T) {
	dir := t.TempDir()
	authDir := filepath.Join(dir, "system", "auth")
	if err := os.MkdirAll(authDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(authDir, authLogName), []byte("bad"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(dir); !commons.Is(err, commons.ErrAuthLogCorrupt) {
		t.Fatalf("open: %v", err)
	}
	if _, err := Open(dir); !commons.Is(err, commons.ErrAuthLogCorrupt) {
		t.Fatalf("second open seeded or changed error: %v", err)
	}
}

func TestSnapshotOffsetPastLog(t *testing.T) {
	dir := t.TempDir()
	s, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
	authDir := filepath.Join(dir, "system", "auth")
	raw, err := os.ReadFile(filepath.Join(authDir, snapshotName))
	if err != nil {
		t.Fatal(err)
	}
	st, err := decodeSnapshot(raw)
	if err != nil {
		t.Fatal(err)
	}
	st.Offset = st.Offset + 100
	payload, err := encodeSnapshot(*st)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(authDir, snapshotName), payload, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(dir); err == nil {
		t.Fatal("expected offset past log to fail")
	}
}

func TestSnapshotWithoutSuperuser(t *testing.T) {
	dir := t.TempDir()
	authDir := filepath.Join(dir, "system", "auth")
	if err := os.MkdirAll(authDir, 0o700); err != nil {
		t.Fatal(err)
	}
	payload, err := encodeSnapshot(snapshotState{
		Offset: 0,
		Users: []User{{
			ID: 2, Username: "alice", PasswordHash: "hash", IsSuper: false,
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(authDir, snapshotName), payload, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(dir); !commons.Is(err, commons.ErrAuthNoSuperuser) {
		t.Fatalf("open: %v", err)
	}
}

func TestUnknownUserDenied(t *testing.T) {
	s := openAuth(t)
	if err := s.CheckPrivilege(99, "root.sg1.d1", PrivilegeRead); !commons.Is(err, commons.ErrAuthPermissionDenied) {
		t.Fatalf("got %v", err)
	}
	if err := s.CheckPrivilege(1, "", PrivilegeRead); err != nil {
		t.Fatal(err)
	}
}

func openAuth(t *testing.T) *Service {
	t.Helper()
	s, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = s.Close() })
	return s
}

func authLogSize(t *testing.T, dataDir string) int64 {
	t.Helper()
	fi, err := os.Stat(filepath.Join(dataDir, "system", "auth", authLogName))
	if err != nil {
		t.Fatal(err)
	}
	return fi.Size()
}
