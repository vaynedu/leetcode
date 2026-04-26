package interview

/*
145. 二叉树后序遍历

给你二叉树的根节点 root，返回后序遍历结果（左→右→根）
*/

// Postorder145 后序遍历递归版
func Postorder145(root *TreeNode) []int {
	var result []int
	postorderRec(root, &result)
	return result
}

func postorderRec(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	postorderRec(node.Left, result)
	postorderRec(node.Right, result)
	*result = append(*result, node.Val)
}

// Postorder145Iter 后序遍历迭代版
func Postorder145Iter(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// reverse: 根→右→左 → 左→右→根
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
