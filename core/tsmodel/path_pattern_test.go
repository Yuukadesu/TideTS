package tsmodel

import "testing"

func TestMatchDevicePattern(t *testing.T) {
	tests := []struct {
		device, pattern string
		want            bool
	}{
		{"root.sg1.d1", "", true},
		{"root.sg1.d1", "root.sg1", true},
		{"root.sg1.d1", "root.sg2", false},
		{"root.sg1.d1", "root.sg1.**", true},
		{"root.sg2.d1", "root.sg1.**", false},
		{"root.d1", "root.*", true},
		{"root.sg1.d1", "root.*", false},
		{"root.sg1.d1", "root.**", true},
		{"root.sg1", "root.sg1.**", true},
	}
	for _, tc := range tests {
		got := MatchDevicePattern(tc.device, tc.pattern)
		if got != tc.want {
			t.Errorf("MatchDevicePattern(%q, %q) = %v, want %v", tc.device, tc.pattern, got, tc.want)
		}
	}
}
