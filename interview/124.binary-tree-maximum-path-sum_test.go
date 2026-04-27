package interview

import "testing"

func TestMaxPathSum(t *testing.T) {
    tests := []struct {
        name string
        root *TreeNode
        want int
    }{
        {"基本", NewTree([]any{-10, 9, 20, nil, nil, 15, 7}), 42},
        {"单节点", NewTree([]any{1}), 1},
        {"全负取最大", NewTree([]any{-3}), -3},
        {"单边", NewTree([]any{2, -1}), 2},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := maxPathSum(tt.root)
            if got != tt.want {
                t.Errorf("maxPathSum() = %d, want %d", got, tt.want)
            }
        })
    }
}
