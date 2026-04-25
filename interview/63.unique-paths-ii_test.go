package interview

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]int
		want   int
	}{
		{"基本", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{"起点障碍", [][]int{{1, 0}}, 0},
		{"终点障碍", [][]int{{0, 0}, {0, 1}}, 0},
		{"无障碍", [][]int{{0, 0}, {0, 0}}, 2},
		{"一格", [][]int{{0}}, 1},
		{"一格障碍", [][]int{{1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePathsWithObstacles(tt.grid)
			if got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %d, want %d", got, tt.want)
			}
		})
	}
}
