package interview

/*
236. 二叉树的最近公共祖先

给定二叉树 root 和两个节点 p、q
返回它们的最近公共祖先（LCA）

时间：O(n) | 空间：O(h)
*/

// LowestCommonAncestor 递归
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	// 在左右子树中查找
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	// 如果左右都找到了，说明当前节点是 LCA
	if left != nil && right != nil {
		return root
	}
	// 否则返回非空的那个
	if left != nil {
		return left
	}
	return right
}
