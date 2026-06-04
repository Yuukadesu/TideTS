package tsmodel

import (
	"errors"
	"testing"

	commons "github.com/hanami/tidets/commons/errors"
)

func TestParseDataType(t *testing.T) {
	dt, err := ParseDataType("double")
	if err != nil || dt != DataTypeDouble {
		t.Fatalf("double: dt=%v err=%v", dt, err)
	}
	dt, err = ParseDataType("TEXT")
	if err != nil || dt != DataTypeText {
		t.Fatalf("TEXT: dt=%v err=%v", dt, err)
	}
	_, err = ParseDataType("INVALID")
	if !errors.Is(err, commons.ErrSQLDataTypeInvalid) {
		t.Fatalf("expected ErrSQLDataTypeInvalid, got %v", err)
	}
}
