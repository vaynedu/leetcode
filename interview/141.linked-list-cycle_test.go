package interview

import "testing"

func TestHasCycle(t *testing.T) {
    // 构造有环链表：1->2->3->2
    a := &ListNode{Val: 1}
    b := &ListNode{Val: 2}
    c := &ListNode{Val: 3}
    a.Next = b
    b.Next = c
    c.Next = b // 形成环

    tests := []struct {
        name string
        head *ListNode
        want bool
    }{
        {"有环", a, true},
        {"无环", NewList(1, 2, 3), false},
        {"空", nil, false},
        {"单节点自环", &ListNode{Val: 1}, false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := hasCycle(tt.head)
            if got != tt.want {
                t.Errorf("hasCycle() = %v, want %v", got, tt.want)
            }
        })
    }
}
