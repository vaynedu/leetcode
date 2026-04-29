package interview

import "testing"

func TestPostorder145(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	result := Postorder145(root)
	expected := []int{4, 5, 2, 3, 1}
	if !intEq(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestPostorder145Iter(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	result := Postorder145Iter(root)
	expected := []int{4, 5, 2, 3, 1}
	if !intEq(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}
