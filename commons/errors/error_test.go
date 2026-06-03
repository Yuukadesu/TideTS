package commons

import (
	"errors"
	"testing"

	"google.golang.org/grpc/status"
)

func TestErrorIsAndGRPC(t *testing.T) {
	err := ErrStorageInvalidRange
	if !Is(err, ErrStorageInvalidRange) {
		t.Fatal("Is should match sentinel")
	}
	st, ok := status.FromError(ToGRPCStatus(err))
	if !ok || st.Code().String() != "InvalidArgument" {
		t.Fatalf("grpc: %+v", st)
	}
}

func TestWrapUnwrap(t *testing.T) {
	inner := errors.New("disk")
	err := Wrap("segment", CodeCorrupt, "compact open", inner)
	if !errors.Is(err, inner) {
		t.Fatal("unwrap")
	}
}
