package interview

import "testing"

func TestIsValidBST(t *testing.T) {
    tests := []struct {
        name string
        root *TreeNode
        want bool
    }{
        {"有效", NewTree([]any{2, 1, 3}), true},
        {"无效", NewTree([]any{5, 1, 4, nil, nil, 3, 6}), false},
        {"右斜", NewTree([]any{2, 1, 3}), true},
        {"右子节点", NewTree([]any{5, 1, 6, nil, 3, 4, 7}), false},
        {"单节点", NewTree([]any{1}), true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := isValidBST(tt.root)
            if got != tt.want {
                t.Errorf("isValidBST() = %v, want %v", got, tt.want)
            }
        })
    }
}
