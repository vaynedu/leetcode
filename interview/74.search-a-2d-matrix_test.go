package interview

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{"正常-存在",
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{"正常-不存在",
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
		{"左上角",
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 1, true},
		{"右下角",
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 60, true},
		{"空矩阵",
			[][]int{}, 1, false},
		{"单行",
			[][]int{{1, 3, 5}}, 3, true},
		{"单列",
			[][]int{{1}, {3}, {5}}, 5, true},
		{"单元素-匹配",
			[][]int{{1}}, 1, true},
		{"单元素-不匹配",
			[][]int{{1}}, 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("SearchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
