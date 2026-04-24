package interview

import "testing"

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"目标在中间",
			[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{"目标在首尾",
			[]int{5, 7, 7, 8, 8, 10}, 5, []int{0, 0}},
		{"目标在末尾",
			[]int{5, 7, 7, 8, 8, 10}, 10, []int{5, 5}},
		{"目标不存在",
			[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{"空数组",
			[]int{}, 0, []int{-1, -1}},
		{"单元素-匹配",
			[]int{5}, 5, []int{0, 0}},
		{"单元素-不匹配",
			[]int{5}, 3, []int{-1, -1}},
		{"全部相同-匹配",
			[]int{2, 2, 2, 2}, 2, []int{0, 3}},
		{"全部相同-不匹配",
			[]int{2, 2, 2, 2}, 3, []int{-1, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRange(tt.nums, tt.target)
			if got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("SearchRange(%v, %d) = %v, want %v",
					tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
