package interview

import (
    "testing"
)

func TestPermuteUnique(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want [][]int
    }{
        {
            name: "基本",
            nums: []int{1, 1, 2},
            want: [][]int{
                {1, 1, 2}, {1, 2, 1}, {2, 1, 1},
            },
        },
        {
            name: "全重复",
            nums: []int{1, 1},
            want: [][]int{{1, 1}},
        },
        {
            name: "三个不同",
            nums: []int{1, 2, 3},
            want: [][]int{
                {1, 2, 3}, {1, 3, 2},
                {2, 1, 3}, {2, 3, 1},
                {3, 1, 2}, {3, 2, 1},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := permuteUnique(tt.nums)
            if !intSliceEq(got, tt.want) {
                t.Errorf("permuteUnique(%v) = %v, want %v", tt.nums, got, tt.want)
            }
        })
    }
}

