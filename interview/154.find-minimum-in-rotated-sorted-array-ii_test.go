package interview

import "testing"

func TestFindMinII(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"正常",
			[]int{1, 3, 5}, 1},
		{"有重复",
			[]int{2, 2, 2, 0, 1}, 0},
		{"两元素",
			[]int{2, 1}, 1},
		{"两元素相同",
			[]int{2, 2}, 2},
		{"无旋转",
			[]int{1, 2, 3}, 1},
		{"空数组",
			[]int{}, 0},
		{"单元素",
			[]int{5}, 5},
		{"全相同",
			[]int{2, 2, 2, 2}, 2},
		{"最小值在中间",
			[]int{3, 4, 5, 1, 2}, 1},
		{"最小值在末尾",
			[]int{1, 2, 3, 4, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinII(tt.nums); got != tt.want {
				t.Errorf("FindMinII(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
