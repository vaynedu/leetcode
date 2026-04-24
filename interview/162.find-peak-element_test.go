package interview

import "testing"

func TestFindPeakElement(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"正常", []int{1, 2, 3, 1}, 2},
        {"递增", []int{1, 2, 3, 4, 5}, 4},
        {"递减", []int{5, 4, 3, 2, 1}, 0},
        {"两个元素-升", []int{1, 2}, 1},
        {"两个元素-降", []int{2, 1}, 0},
        {"两个元素-相等", []int{1, 1}, 0},
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
