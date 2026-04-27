package interview

import "testing"

func TestFindPeakElement(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"基本", []int{1, 2, 3, 1}, 2},
        {"两个", []int{1, 2}, 1},
        {"递减", []int{3, 2, 1}, 0},
        {"单元素", []int{1}, 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := findPeakElement(tt.nums)
            if got != tt.want {
                t.Errorf("findPeakElement(%v) = %d, want %d", tt.nums, got, tt.want)
            }
        })
    }
}
