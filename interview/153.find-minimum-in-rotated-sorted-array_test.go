package interview

import "testing"

func TestFindMin153(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"正常", []int{3, 4, 5, 1, 2}, 1},
        {"正常2", []int{4, 5, 6, 7, 0, 1, 2}, 0},
        {"未旋转", []int{1, 2, 3, 4, 5}, 1},
        {"两个元素", []int{2, 1}, 1},
        {"两个元素2", []int{1, 2}, 1},
        {"单元素", []int{1}, 1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := findMin153(tt.nums)
            if got != tt.want {
                t.Errorf("findMin153(%v) = %d, want %d", tt.nums, got, tt.want)
            }
        })
    }
}
