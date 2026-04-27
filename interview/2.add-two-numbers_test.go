package interview

import "testing"

func TestAddTwoNumbers(t *testing.T) {
    tests := []struct {
        name string
        l1 *ListNode
        l2 *ListNode
        want []int
    }{
        {"基本", NewList(2, 4, 3), NewList(5, 6, 4), []int{7, 0, 8}},
        {"不同长度", NewList(9, 9, 9), NewList(1), []int{0, 0, 0, 1}},
        {"零", NewList(0), NewList(0), []int{0}},
        {"进位到新节点", NewList(5), NewList(5), []int{0, 1}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := addTwoNumbers(tt.l1, tt.l2)
            if !listEqual(got, tt.want) {
                t.Errorf("addTwoNumbers() = %v, want %v", listToSlice(got), tt.want)
            }
        })
    }
}
