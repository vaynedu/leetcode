package interview

import "testing"

func TestDetectCycle(t *testing.T) {
    // 构造 1->2->3->2 环，入口是 2
    a := &ListNode{Val: 1}
    b := &ListNode{Val: 2}
    c := &ListNode{Val: 3}
    a.Next = b
    b.Next = c
    c.Next = b

    tests := []struct {
        name string
        head *ListNode
        wantVal int
    }{
        {"有环入口2", a, 2},
        {"无环", NewList(1, 2, 3), -1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := detectCycle(tt.head)
            if tt.wantVal == -1 {
                if got != nil {
                    t.Errorf("detectCycle() = %v, want nil", got)
                }
            } else {
                if got == nil || got.Val != tt.wantVal {
                    t.Errorf("detectCycle() = %v, want val %d", got, tt.wantVal)
                }
            }
        })
    }
}
