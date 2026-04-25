package interview

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m, n int
		want int
	}{
		{"基本", 3, 7, 28},
		{"1x1", 1, 1, 1},
		{"1行", 1, 5, 1},
		{"1列", 4, 1, 1},
		{"2x3", 2, 3, 3},
		{"3x3", 3, 3, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			if got != tt.want {
				t.Errorf("uniquePaths(%d,%d) = %d, want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}
