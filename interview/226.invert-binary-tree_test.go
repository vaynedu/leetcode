package interview

import "testing"

func TestInvertTree(t *testing.T) {
	// [4,2,7,1,3,6,9] 翻转后左变右
	root := NewTree([]any{4, 2, 7, 1, 3, 6, 9})
	inverted := invertTree(root)
	// 翻转后: 4的左是7(原来右), 右是2(原来左)
	if inverted.Left.Val != 7 || inverted.Right.Val != 2 {
		t.Errorf("invertTree() failed, got left=%d right=%d, want left=7 right=2", inverted.Left.Val, inverted.Right.Val)
	}
	// 2的左右交换: 左=3, 右=1
	if inverted.Right.Left.Val != 3 || inverted.Right.Right.Val != 1 {
		t.Errorf("inverted.Right children wrong")
	}
}
