package interview

import "testing"

func TestLengthOfLIS(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"基本", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
        {"全降", []int{5, 4, 3, 2, 1}, 1},
        {"全升", []int{1, 2, 3, 4, 5}, 5},
        {"重复", []int{1, 1, 1}, 1},
        {"两个", []int{1, 3}, 2},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := lengthOfLIS(tt.nums)
            if got != tt.want {
                t.Errorf("lengthOfLIS(%v) = %d, want %d", tt.nums, got, tt.want)
            }
        })
    }
}
