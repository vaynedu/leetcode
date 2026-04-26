package interview

/*
572. 另一棵树的子树

给你两棵二叉树 root 和 subRoot，判断 subRoot 是否是 root 的子树

子树：subRoot 的结构和值与 root 的某个子树完全相同

时间：O(m*n) | 空间：O(m) m=root节点数
*/

// IsSubtree 双重递归
func IsSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	// 如果当前节点开始匹配，返回 true
	if isSameTree(root, subRoot) {
		return true
	}
	// 否则检查左右子树
	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

// isSameTree 判断两棵树是否相同
func isSameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val &&
		isSameTree(a.Left, b.Left) &&
		isSameTree(a.Right, b.Right)
}
