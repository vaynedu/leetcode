package interview

/*
94. 二叉树中序遍历

给你二叉树的根节点 root，返回中序遍历结果（左→根→右）
*/

// Inorder94 中序遍历递归版
func Inorder94(root *TreeNode) []int {
	var result []int
	inorderRec(root, &result)
	return result
}

func inorderRec(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inorderRec(node.Left, result)
	*result = append(*result, node.Val)
	inorderRec(node.Right, result)
}

// Inorder94Iter 中序遍历迭代版
func Inorder94Iter(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}
	return result
}
