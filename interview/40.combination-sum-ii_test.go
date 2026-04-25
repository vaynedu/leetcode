package interview

import (
    "testing"
)

func TestCombinationSum2(t *testing.T) {
    tests := []struct {
        name      string
        candidates []int
        target    int
        want      [][]int
    }{
        {
            name:      "基本",
            candidates: []int{10, 1, 2, 7, 6, 1, 5},
            target:    8,
            want:      [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
        },
        {
            name:      "有重复",
            candidates: []int{1, 1, 1, 1},
            target:    2,
            want:      [][]int{{1, 1}},
        },
        {
            name:      "无解",
            candidates: []int{1},
            target:    2,
            want:      nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := combinationSum2(tt.candidates, tt.target)
            if !intSliceEq(got, tt.want) {
                t.Errorf("combinationSum2(%v, %d) = %v, want %v", tt.candidates, tt.target, got, tt.want)
            }
        })
    }
}

