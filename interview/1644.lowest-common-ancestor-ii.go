package interview

/*
1644. 二叉树的最近公共祖先 II

与 236 区别：树中节点有指向父节点的指针 parent
可以利用 parent 从节点向上遍历找 LCA

时间：O(h) | 空间：O(h)
*/

// TreeNodeWithParent 带父指针的节点
type TreeNodeWithParent struct {
	Val       int
	Parent    *TreeNodeWithParent
	Left      *TreeNodeWithParent
	Right     *TreeNodeWithParent
}

// LowestCommonAncestorII 利用 parent 指针
func LowestCommonAncestorII(root, p, q *TreeNodeWithParent) *TreeNodeWithParent {
	// 收集 p 到根的路径
	pathP := map[*TreeNodeWithParent]bool{}
	for p != nil {
		pathP[p] = true
		p = p.Parent
	}
	// 从 q 向上找，第一个出现在 pathP 中的就是 LCA
	for q != nil {
		if pathP[q] {
			return q
		}
		q = q.Parent
	}
	return nil
}
