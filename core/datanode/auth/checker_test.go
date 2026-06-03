package auth

import "testing"

func TestCheckerAuthenticate(t *testing.T) {
	c := Checker{}
	id, ok, err := c.Authenticate(RootUsername, RootPassword())
	if err != nil || !ok || id != RootUserID {
		t.Fatalf("root login: id=%d ok=%v err=%v", id, ok, err)
	}
	_, ok, err = c.Authenticate(RootUsername, "wrong")
	if err != nil || ok {
		t.Fatal("expected reject wrong password")
	}
}

func TestCheckerPrivilege(t *testing.T) {
	c := Checker{}
	if err := c.CheckPrivilege(RootUserID, RootUsername, "root.sg1.d1", PrivilegeWrite); err != nil {
		t.Fatal(err)
	}
	if err := c.CheckPrivilege(1, "guest", "root.sg1.d1", PrivilegeRead); err != ErrPermissionDenied {
		t.Fatalf("guest: %v", err)
	}
}
