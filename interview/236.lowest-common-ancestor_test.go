package interview

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	//      3
	//    /   \
	//   5     1
	//  / \   / \
	// 6   2  0   8
	//    / \
	//    7  4
	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	p := root.Left // 5
	q := root.Left.Right.Right // 4
	lca := LowestCommonAncestor(root, p, q)
	if lca.Val != 5 {
		t.Errorf("LCA of 5 and 4: got %d, want 5", lca.Val)
	}
}

func TestLowestCommonAncestorRoot(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	lca := LowestCommonAncestor(root, root, root.Left)
	if lca != root {
		t.Error("root and child: LCA should be root")
	}
}
