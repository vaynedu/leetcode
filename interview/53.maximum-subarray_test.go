package interview

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"基本", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},  // [4,-1,2,1]
		{"全正", []int{1, 2, 3, 4}, 10},
		{"全负", []int{-1, -2, -3}, -1},
		{"单个", []int{5}, 5},
		{"两个", []int{-1, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("maxSubArray(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
