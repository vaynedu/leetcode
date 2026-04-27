package interview

import "testing"

func TestReverseBetween(t *testing.T) {
    tests := []struct {
        name string
        head *ListNode
        left int
        right int
        want []int
    }{
        {"基本", NewList(1, 2, 3, 4, 5), 2, 4, []int{1, 4, 3, 2, 5}},
        {"反转全部", NewList(1, 2, 3), 1, 3, []int{3, 2, 1}},
        {"单节点", NewList(1), 1, 1, []int{1}},
        {"头部", NewList(1, 2, 3), 1, 2, []int{2, 1, 3}},
        {"尾部", NewList(1, 2, 3), 2, 3, []int{1, 3, 2}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := reverseBetween(tt.head, tt.left, tt.right)
            if !listEqual(got, tt.want) {
                t.Errorf("reverseBetween() = %v, want %v", listToSlice(got), tt.want)
            }
        })
    }
}
