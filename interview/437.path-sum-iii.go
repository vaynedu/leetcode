package interview

/*
437. 路径总和 III

给定二叉树和一个目标值，返回所有路径（起点可以是任意节点，不一定是根）
路径方向必须向下（父子关系）

时间：O(n) | 空间：O(h)
*/

// PathSumIII 前缀和解法
func PathSumIII(root *TreeNode, target int) int {
	// 前缀和计数：key=前缀和，value=出现次数
	prefix := map[int]int{0: 1}
	return pathSumIIIHelper(root, 0, target, prefix)
}

// 返回满足条件的路径数
// currSum: 从根到当前节点的路径和
func pathSumIIIHelper(node *TreeNode, currSum, target int, prefix map[int]int) int {
	if node == nil {
		return 0
	}
	// 当前节点的前缀和
	currSum += node.Val
	// 找有多少个前缀和等于 currSum - target
	count := prefix[currSum-target]
	// 更新当前前缀和计数
	prefix[currSum]++

	// 递归左右子树
	count += pathSumIIIHelper(node.Left, currSum, target, prefix)
	count += pathSumIIIHelper(node.Right, currSum, target, prefix)

	// 回溯：删除当前节点的前缀和贡献
	prefix[currSum]--

	return count
}
