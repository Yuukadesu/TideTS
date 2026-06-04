package utils

import (
	"testing"

	"github.com/hanami/tidets/core/tsmodel"
)

func TestMergeSortedPreferNewer(t *testing.T) {
	a := []tsmodel.Point{{Timestamp: 10, Value: tsmodel.NewDouble(1)}}
	b := []tsmodel.Point{{Timestamp: 10, Value: tsmodel.NewDouble(9)}, {Timestamp: 20, Value: tsmodel.NewDouble(2)}}
	got := MergeSortedPreferNewer(a, b)
	if len(got) != 2 || !got[0].Value.Equal(tsmodel.NewDouble(9)) {
		t.Fatalf("got %+v", got)
	}
}
