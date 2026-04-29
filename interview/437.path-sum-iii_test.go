package interview

import "testing"

func TestPathSumIII(t *testing.T) {
	//     10
	//    /  \
	//   5   -3
	//  / \    \
	// 3   2   11
	// \    \
	//  3 -2   1
	root := &TreeNode{Val: 10,
		Left: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}},
		},
		Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}},
	}
	result := PathSumIII(root, 8)
	if result != 3 {
		t.Errorf("PathSumIII(root, 8) = %d, want 3", result)
	}
}

func TestPathSumIIIEmpty(t *testing.T) {
	if PathSumIII(nil, 1) != 0 {
		t.Error("nil tree should return 0")
	}
}

func TestPathSumIIISingle(t *testing.T) {
	root := &TreeNode{Val: 1}
	if PathSumIII(root, 1) != 1 {
		t.Error("single node with value 1 should have 1 path")
	}
}
