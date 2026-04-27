package interview

import "testing"

func TestMaxDepth(t *testing.T) {
    tests := []struct {
        name string
        root *TreeNode
        want int
    }{
        {"基本", NewTree([]any{3, 9, 20, nil, nil, 15, 7}), 3},
        {"单层", NewTree([]any{1, nil, 2}), 2},
        {"单节点", NewTree([]any{1}), 1},
        {"空树", NewTree([]any{}), 0},
        {"左斜", NewTree([]any{1, 2, nil, 3, nil, 4}), 4},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := maxDepth(tt.root)
            if got != tt.want {
                t.Errorf("maxDepth() = %v, want %v", got, tt.want)
            }
        })
    }
}
