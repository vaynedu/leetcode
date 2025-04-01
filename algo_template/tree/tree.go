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
