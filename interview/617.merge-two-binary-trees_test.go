package interview

import "testing"

func TestMergeTrees(t *testing.T) {
	// t1: 1(3(5),2)   t2: 2(1(null),3(null))
	// 合并: 3(4(5),5)
	t1 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 2},
	}
	t2 := &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	result := MergeTrees(t1, t2)
	if result.Val != 3 {
		t.Errorf("root val: got %d, want 3", result.Val)
	}
	if result.Left.Val != 4 {
		t.Errorf("left val: got %d, want 4", result.Left.Val)
	}
	if result.Left.Left.Val != 5 {
		t.Errorf("left.left val: got %d, want 5", result.Left.Left.Val)
	}
}

func TestMergeTreesNil(t *testing.T) {
	result := MergeTrees(nil, nil)
	if result != nil {
		t.Error("both nil should return nil")
	}
}
