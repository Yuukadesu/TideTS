package utils

import (
	"testing"

	"github.com/hanami/tidets/core/storageengine/model"
)

func TestMergeSortedPreferNewer(t *testing.T) {
	a := []model.Point{{Timestamp: 10, Value: model.NewDouble(1)}}
	b := []model.Point{{Timestamp: 10, Value: model.NewDouble(9)}, {Timestamp: 20, Value: model.NewDouble(2)}}
	got := MergeSortedPreferNewer(a, b)
	if len(got) != 2 || !got[0].Value.Equal(model.NewDouble(9)) {
		t.Fatalf("got %+v", got)
	}
}
