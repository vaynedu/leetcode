package tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func traverse(root *TreeNode) {
	// 前序遍历
	traverse(root.Left)
	// 中序遍历
	traverse(root.Right)
	// 后序遍历
}

// 深度和高度
// 深度是前序遍历，任意一个节点到根节点的距离
// 高度是后序遍历，叶子节点到任意节点的距离
// 两个刚好是相反的

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 后续遍历 左右中
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
