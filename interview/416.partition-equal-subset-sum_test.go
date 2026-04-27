package interview

import "testing"

func TestCanPartition(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want bool
    }{
        {"基本", []int{1, 5, 11, 5}, true},
        {"可分", []int{1, 2, 3, 6}, true},
        {"不可分", []int{1, 2, 3, 5}, false},
        {"单元素", []int{1}, false},
        {"两元素", []int{1, 1}, true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := canPartition(tt.nums)
            if got != tt.want {
                t.Errorf("canPartition(%v) = %v, want %v", tt.nums, got, tt.want)
            }
        })
    }
}
