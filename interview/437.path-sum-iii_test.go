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
	// 路径: 5->2->-2=5, 5->3=8, 5->3->3=11, -3->11=8, 10->5->3=18, 10->5->3->3=21...
	// 按 target=8: [5,3],[5,3,3],[-3,11],[10,5,3],[-2+10],[10,-3,11]
	result := PathSumIII(root, 8)
	if result == 0 {
		t.Errorf("expected > 0 paths for target 8, got %d", result)
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
