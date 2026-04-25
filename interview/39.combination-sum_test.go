package interview

import (
    "testing"
)

func TestCombinationSum(t *testing.T) {
    tests := []struct {
        name      string
        candidates []int
        target    int
        want      int // 数量
    }{
        {name: "基本", candidates: []int{2, 3, 6, 7}, target: 7, want: 2},
        {name: "两个数", candidates: []int{2, 3, 5}, target: 8, want: 3},
        {name: "全1", candidates: []int{1}, target: 1, want: 1},
        {name: "无解", candidates: []int{2}, target: 1, want: 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := combinationSum(tt.candidates, tt.target)
            if len(got) != tt.want {
                t.Errorf("combinationSum(%v, %d) returned %d combos, want %d", tt.candidates, tt.target, len(got), tt.want)
            }
            // 验证每个组合的和
            for _, combo := range got {
                sum := 0
                for _, v := range combo {
                    sum += v
                }
                if sum != tt.target {
                    t.Errorf("combinationSum(%v, %d) produced combo %v with sum %d", tt.candidates, tt.target, combo, sum)
                }
            }
        })
    }
}
