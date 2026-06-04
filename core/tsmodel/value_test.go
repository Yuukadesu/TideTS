package tsmodel

import (
	"bytes"
	"testing"
)

func TestValuePayloadRoundTrip(t *testing.T) {
	cases := []Value{
		NewBoolean(true),
		NewInt32(-42),
		NewInt64(1 << 40),
		NewFloat(3.14),
		NewDouble(36.5),
		NewText("ok"),
	}
	for _, want := range cases {
		var buf bytes.Buffer
		if err := want.WritePayload(&buf); err != nil {
			t.Fatal(err)
		}
		got, err := ReadValuePayload(&buf)
		if err != nil {
			t.Fatal(err)
		}
		if !got.Equal(want) {
			t.Fatalf("want %+v got %+v", want, got)
		}
	}
}
