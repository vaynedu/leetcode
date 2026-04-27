package interview

import "testing"

func TestGetIntersectionNode(t *testing.T) {
    // 构造相交链表: 1->2->3->4, 5->3->4
    a := &ListNode{Val: 1}
    b := &ListNode{Val: 2}
    c := &ListNode{Val: 3}
    d := &ListNode{Val: 4}
    a.Next = b
    b.Next = c
    c.Next = d
    e := &ListNode{Val: 5}
    e.Next = c // 相交于 c

    tests := []struct {
        name string
        headA *ListNode
        headB *ListNode
        wantVal int
    }{
        {"相交", a, e, 3},
        {"不相交", NewList(1, 2), NewList(3, 4), -1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := getIntersectionNode(tt.headA, tt.headB)
            if tt.wantVal == -1 {
                if got != nil {
                    t.Errorf("getIntersectionNode() = %v, want nil", got)
                }
            } else {
                if got == nil || got.Val != tt.wantVal {
                    t.Errorf("getIntersectionNode() = %v, want val %d", got, tt.wantVal)
                }
            }
        })
    }
}
