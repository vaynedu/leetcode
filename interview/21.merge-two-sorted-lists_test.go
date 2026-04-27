package interview

import "testing"

func TestMergeTwoLists(t *testing.T) {
    tests := []struct {
        name string
        l1 *ListNode
        l2 *ListNode
        want []int
    }{
        {"基本", NewList(1, 2, 4), NewList(1, 3, 4), []int{1, 1, 2, 3, 4, 4}},
        {"空l1", nil, NewList(1), []int{1}},
        {"空l2", NewList(1), nil, []int{1}},
        {"交替", NewList(1, 3, 5), NewList(2, 4, 6), []int{1, 2, 3, 4, 5, 6}},
        {"重复", NewList(1, 1, 1), NewList(1, 1, 1), []int{1, 1, 1, 1, 1, 1}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := mergeTwoLists(tt.l1, tt.l2)
            if !listEqual(got, tt.want) {
                t.Errorf("mergeTwoLists() = %v, want %v", listToSlice(got), tt.want)
            }
        })
    }
}
