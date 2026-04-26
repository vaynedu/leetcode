package interview

/*
112. 路径总和

给你二叉树 root 和目标值 targetSum
判断是否存在从根到叶子的路径，使路径上所有节点值之和等于 targetSum

时间：O(n) | 空间：O(h)
*/

// HasPathSum 递归版
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 叶子节点：判断是否等于目标值
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	// 递归左右子树，目标值减去当前节点值
	return HasPathSum(root.Left, targetSum-root.Val) ||
		HasPathSum(root.Right, targetSum-root.Val)
}
