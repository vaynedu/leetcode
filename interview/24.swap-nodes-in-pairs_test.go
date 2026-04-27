package interview

import "testing"

func TestSwapPairs(t *testing.T) {
    tests := []struct {
        name string
        head *ListNode
        want []int
    }{
        {"基本", NewList(1, 2, 3, 4), []int{2, 1, 4, 3}},
        {"奇数", NewList(1, 2, 3), []int{2, 1, 3}},
        {"两个", NewList(1, 2), []int{2, 1}},
        {"单节点", NewList(1), []int{1}},
        {"空", nil, nil},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := swapPairs(tt.head)
            if !listEqual(got, tt.want) {
                t.Errorf("swapPairs() = %v, want %v", listToSlice(got), tt.want)
            }
        })
    }
}
