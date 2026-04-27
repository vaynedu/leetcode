package interview

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
    tests := []struct {
        name string
        root *TreeNode
        want int
    }{
        {"基本", NewTree([]any{1, 2, 3, 4, 5}), 3},
        {"单节点", NewTree([]any{1}), 0},
        {"直线", NewTree([]any{1, 2, nil, 3, nil, 4}), 3},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := diameterOfBinaryTree(tt.root)
            if got != tt.want {
                t.Errorf("diameterOfBinaryTree() = %d, want %d", got, tt.want)
            }
        })
    }
}
