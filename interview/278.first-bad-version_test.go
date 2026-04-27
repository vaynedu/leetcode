package interview

import "testing"

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name       string
		n          int
		badVersion int
		want       int
	}{
		{"基本", 5, 4, 4},
		{"第一个", 3, 1, 1},
		{"全是", 3, 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstBadVersion(tt.n, tt.badVersion)
			if got != tt.want {
				t.Errorf("firstBadVersion(%d, %d) = %d, want %d", tt.n, tt.badVersion, got, tt.want)
			}
		})
	}
}
