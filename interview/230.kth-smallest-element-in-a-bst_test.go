package interview

import "testing"

func TestKthSmallest(t *testing.T) {
    root := NewTree([]any{5, 3, 6, 2, 4, nil, nil, 1})
    tests := []struct {
        name string
        root *TreeNode
        k int
        want int
    }{
        {"第1小", root, 1, 1},
        {"第3小", root, 3, 4},
        {"第5小", root, 5, 6},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := kthSmallest(tt.root, tt.k)
            if got != tt.want {
                t.Errorf("kthSmallest(k=%d) = %d, want %d", tt.k, got, tt.want)
            }
        })
    }
}
