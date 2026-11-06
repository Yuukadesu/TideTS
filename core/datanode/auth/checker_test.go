package auth

import (
	"testing"

	"github.com/hanami/tidets/core/authority"
)

func TestCheckerAuthenticate(t *testing.T) {
	c := openChecker(t)
	id, ok, err := c.Authenticate(RootUsername, RootPassword())
	if err != nil || !ok || id == RootUserID || id <= 0 {
		t.Fatalf("root login: id=%d ok=%v err=%v", id, ok, err)
	}
	_, ok, err = c.Authenticate(RootUsername, "wrong")
	if err != nil || ok {
		t.Fatal("expected reject wrong password")
	}
	_, ok, err = c.Authenticate("nobody", "wrong")
	if err != nil || ok {
		t.Fatal("expected reject unknown user")
	}
}

func TestCheckerPrivilege(t *testing.T) {
	c := openChecker(t)
	id, ok, err := c.Authenticate(RootUsername, RootPassword())
	if err != nil || !ok {
		t.Fatal(err)
	}
	if err := c.CheckPrivilege(id, RootUsername, "root.sg1.d1", PrivilegeWrite); err != nil {
		t.Fatal(err)
	}
	if err := c.CheckPrivilege(99, "guest", "root.sg1.d1", PrivilegeRead); err != ErrPermissionDenied {
		t.Fatalf("guest: %v", err)
	}
	if err := c.CheckPrivilege(id, RootUsername, "", PrivilegeRead); err != nil {
		t.Fatal(err)
	}
}

func openChecker(t *testing.T) Checker {
	t.Helper()
	accounts, err := authority.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = accounts.Close() })
	return NewChecker(accounts)
}
