package interview

/*
113. 路径总和 II

给你二叉树 root 和目标值 targetSum
返回所有从根到叶子节点且路径总和等于 targetSum 的路径
*/

// PathSumII 返回所有满足条件的路径
func PathSumII(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int
	pathSumRec(root, targetSum, path, &result)
	return result
}

func pathSumRec(node *TreeNode, target int, path []int, result *[][]int) {
	if node == nil {
		return
	}
	path = append(path, node.Val)
	// 叶子节点：判断是否等于目标值
	if node.Left == nil && node.Right == nil {
		if node.Val == target {
			// 复制 path 切片（防止后续修改影响已收集的路径）
			tmp := make([]int, len(path))
			copy(tmp, path)
			*result = append(*result, tmp)
		}
		return
	}
	pathSumRec(node.Left, target-node.Val, path, result)
	pathSumRec(node.Right, target-node.Val, path, result)
	// path 会自动回溯（因为是值拷贝，不是引用）
}
