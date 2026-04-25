package interview

import (
    "testing"
)

func TestSubsetsWithDup(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {name: "基本", nums: []int{1, 2, 2}, want: 6},
        {name: "全相同", nums: []int{1, 1, 1}, want: 4},
        {name: "无重复", nums: []int{1, 2, 3}, want: 8},
        {name: "两对重复", nums: []int{1, 1, 2, 2}, want: 9},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := subsetsWithDup(tt.nums)
            if len(got) != tt.want {
                t.Errorf("subsetsWithDup(%v) returned %d subsets, want %d", tt.nums, len(got), tt.want)
            }
        })
    }
}
