package interview

import "testing"

func TestSearchRotatedII(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"正常-存在中间",
			[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{"正常-不存在",
			[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{"全部相同-存在",
			[]int{1, 1, 1, 1, 1}, 1, true},
		{"全部相同-不存在",
			[]int{1, 1, 1, 1, 1}, 0, false},
		{"无旋转",
			[]int{1, 3, 5}, 3, true},
		{"空数组",
			[]int{}, 1, false},
		{"单元素-匹配",
			[]int{3}, 3, true},
		{"单元素-不匹配",
			[]int{3}, 2, false},
		{"重复+目标在左",
			[]int{2, 2, 2, 3, 4, 0, 1, 2}, 3, true},
		{"重复+目标不存在",
			[]int{2, 2, 2, 3, 4, 0, 1, 2}, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchRotatedII(tt.nums, tt.target); got != tt.want {
				t.Errorf("SearchRotatedII(%v, %d) = %v, want %v",
					tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
