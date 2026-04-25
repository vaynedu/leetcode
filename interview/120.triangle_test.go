package interview

import "testing"

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		name    string
		triangle [][]int
		want    int
	}{
		{"基本", [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		{"单行", [][]int{{-1}}, -1},
		{"两行", [][]int{{2}, {3, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumTotal(tt.triangle)
			if got != tt.want {
				t.Errorf("minimumTotal() = %d, want %d", got, tt.want)
			}
		})
	}
}
