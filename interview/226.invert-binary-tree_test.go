package interview

import "testing"

func TestInvertTree(t *testing.T) {
    root := NewTree([]any{4, 2, 7, 1, 3, 6, 9})
    inverted := invertTree(root)
    if inverted.Left.Val != 7 || inverted.Right.Val != 2 {
        t.Errorf("invertTree() failed, got left=%d right=%d", inverted.Left.Val, inverted.Right.Val)
    }
    if inverted.Left.Left != nil && inverted.Left.Left.Val != 6 {
        t.Errorf("inverted.Left.Left = %d, want 6", inverted.Left.Left.Val)
    }
}
