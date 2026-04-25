package interview

import (
    "testing"
)

func TestSubsets(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int // 数量: 2^n
    }{
        {name: "基本", nums: []int{1, 2, 3}, want: 8},
        {name: "单个", nums: []int{1}, want: 2},
        {name: "空", nums: []int{}, want: 1},
        {name: "4元素", nums: []int{1, 2, 3, 4}, want: 16},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := subsets(tt.nums)
            if len(got) != tt.want {
                t.Errorf("subsets(%v) returned %d subsets, want %d", tt.nums, len(got), tt.want)
            }
        })
    }
}
