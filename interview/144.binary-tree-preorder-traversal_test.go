package interview

import "testing"

func TestPreorder144(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	result := Preorder144(root)
	expected := []int{1, 2, 4, 5, 3}
	if len(result) != len(expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestPreorder144Iter(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	result := Preorder144Iter(root)
	expected := []int{1, 2, 4, 5, 3}
	if len(result) != len(expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}
