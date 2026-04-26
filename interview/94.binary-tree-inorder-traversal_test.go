package interview

import "testing"

func TestInorder94(t *testing.T) {
	//    1
	//  2   3
	// 4 5
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	result := Inorder94(root)
	expected := []int{4, 2, 5, 1, 3}
	if len(result) != len(expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestInorder94Iter(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	result := Inorder94Iter(root)
	expected := []int{4, 2, 5, 1, 3}
	if len(result) != len(expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}
