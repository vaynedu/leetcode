package interview

import "testing"

func TestHasPathSum(t *testing.T) {
	//    5
	//  4   8
	// 11  13  4
	// 7  2       1
	root := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}},
	}
	if !HasPathSum(root, 22) {
		t.Error("path sum 22 should exist (5+4+11+2)")
	}
	if HasPathSum(root, 21) {
		t.Error("path sum 21 should not exist")
	}
	if HasPathSum(nil, 0) {
		t.Error("nil tree should return false")
	}
}
