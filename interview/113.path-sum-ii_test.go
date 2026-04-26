package interview

import "testing"

func TestPathSumII(t *testing.T) {
	//    5
	//  4   8
	// 7  2
	root := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 8},
	}
	// 路径: [5,4,7]=16, [5,4,2]=11, [5,8]=13
	// target=11 → [5,4,2]
	result := PathSumII(root, 11)
	if len(result) != 1 {
		t.Errorf("expected 1 path for target=11, got %d: %v", len(result), result)
	}
	if len(result) > 0 && result[0][2] != 2 {
		t.Errorf("expected [5,4,2], got %v", result)
	}

	// target=9 → 无路径
	result = PathSumII(root, 9)
	if len(result) != 0 {
		t.Errorf("expected 0 paths for target=9, got %v", result)
	}
}

func TestPathSumIINone(t *testing.T) {
	root := &TreeNode{Val: 1}
	result := PathSumII(root, 0)
	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
}
