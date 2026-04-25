package interview

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"基本", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"有重复", []int{3, 2, 3, 1, 2, 4, 5, 5, 5}, 3, 5},
		{"最小k=1", []int{3, 2, 1, 5, 6, 4}, 1, 6},
		{"最大k", []int{3, 2, 1, 5, 6, 4}, 6, 1},
		{"两个元素", []int{1, 2}, 1, 2},
		{"两个元素k2", []int{1, 2}, 2, 1},
		{"降序", []int{5, 4, 3, 2, 1}, 2, 4},
		{"负数", []int{-1, -2, -3, -4, -5}, 1, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("findKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
