package interview

import (
	"slices"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	//    1
	//  2   3
	//  5
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	result := BinaryTreePaths(root)
	expected := []string{"1->2->5", "1->3"}
	if len(result) != len(expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
	for _, e := range expected {
		if !slices.Contains(result, e) {
			t.Errorf("missing %s", e)
		}
	}
}

func TestBinaryTreePathsSingle(t *testing.T) {
	root := &TreeNode{Val: 1}
	result := BinaryTreePaths(root)
	if len(result) != 1 || result[0] != "1" {
		t.Errorf("expected [1], got %v", result)
	}
}
