package interview

import "testing"

func TestIsSubtree(t *testing.T) {
	// root:
	//    3
	//  4   5
	// 1   2
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 2}},
	}
	// subRoot: 4(1)
	subRoot := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}}
	if !IsSubtree(root, subRoot) {
		t.Error("subtree should match")
	}
}

func TestIsSubtreeFalse(t *testing.T) {
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 2}},
	}
	subRoot := &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}
	if IsSubtree(root, subRoot) {
		t.Error("should not be subtree")
	}
}

func TestIsSubtreeEmpty(t *testing.T) {
	root := &TreeNode{Val: 1}
	if IsSubtree(root, nil) {
		t.Error("nil subRoot should return false")
	}
}
