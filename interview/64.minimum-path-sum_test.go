package interview

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"基本", [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{"1x1", [][]int{{5}}, 5},
		{"1行", [][]int{{1, 2, 3}}, 6},
		{"1列", [][]int{{1}, {2}, {3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minPathSum(tt.grid)
			if got != tt.want {
				t.Errorf("minPathSum(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}
