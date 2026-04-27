package interview

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
    tests := []struct {
        name string
        head *ListNode
        n int
        want []int
    }{
        {"基本", NewList(1, 2, 3, 4, 5), 2, []int{1, 2, 3, 5}},
        {"删除头", NewList(1, 2), 2, []int{2}},
        {"删除尾", NewList(1, 2), 1, []int{1}},
        {"单节点", NewList(1), 1, []int{}},
        {"最后两个", NewList(1, 2, 3), 1, []int{1, 2}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := removeNthFromEnd(tt.head, tt.n)
            if !listEqual(got, tt.want) {
                t.Errorf("removeNthFromEnd() = %v, want %v", listToSlice(got), tt.want)
            }
        })
    }
}
