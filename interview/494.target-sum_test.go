package interview

import "testing"

func TestFindTargetSumWays(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        target int
        want int
    }{
        {"基本", []int{1, 1, 1, 1, 1}, 3, 5},
        {"零", []int{0}, 0, 2},
        {"单个正", []int{1}, 1, 1},
        {"单个负", []int{1}, -1, 1},
        {"无解", []int{1}, 2, 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := findTargetSumWays(tt.nums, tt.target)
            if got != tt.want {
                t.Errorf("findTargetSumWays(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
            }
        })
    }
}
