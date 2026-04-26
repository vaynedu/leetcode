package interview

import "testing"

func TestLowestCommonAncestorII(t *testing.T) {
	//       3
	//    /    \
	//   5      1
	//  / \    / \
	// 6   2  0    8
	//    / \
	//    7  4
	n6 := &TreeNodeWithParent{Val: 6}
	n7 := &TreeNodeWithParent{Val: 7}
	n4 := &TreeNodeWithParent{Val: 4}
	n2 := &TreeNodeWithParent{Val: 2, Left: n7, Right: n4}
	n0 := &TreeNodeWithParent{Val: 0}
	n8 := &TreeNodeWithParent{Val: 8}
	n5 := &TreeNodeWithParent{Val: 5, Left: n6, Right: n2}
	n1 := &TreeNodeWithParent{Val: 1, Left: n0, Right: n8}
	root := &TreeNodeWithParent{Val: 3, Left: n5, Right: n1}

	// 设置父指针
	n6.Parent, n7.Parent, n4.Parent, n2.Parent = n5, n2, n2, n5
	n0.Parent, n8.Parent, n5.Parent, n1.Parent = n1, n1, root, root

	// LCA of 6 and 4: 路径 6->5, 4->2->5, 最先交汇是 5
	lca := LowestCommonAncestorII(root, n6, n4)
	if lca.Val != 5 {
		t.Errorf("LCA of 6 and 4: got %d, want 5", lca.Val)
	}

	// LCA of 6 and 0: 路径 6->5->3, 0->1->3, LCA=3
	lca = LowestCommonAncestorII(root, n6, n0)
	if lca.Val != 3 {
		t.Errorf("LCA of 6 and 0: got %d, want 3", lca.Val)
	}
}

func TestLowestCommonAncestorIINil(t *testing.T) {
	result := LowestCommonAncestorII(nil, nil, nil)
	if result != nil {
		t.Error("nil inputs should return nil")
	}
}
