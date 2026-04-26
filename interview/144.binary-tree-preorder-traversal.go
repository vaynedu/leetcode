package interview

/*
144. 二叉树前序遍历

给你二叉树的根节点 root，返回前序遍历结果（根→左→右）
*/

// Preorder144 前序遍历递归版
func Preorder144(root *TreeNode) []int {
	var result []int
	preorderRec(root, &result)
	return result
}

func preorderRec(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	preorderRec(node.Left, result)
	preorderRec(node.Right, result)
}

// Preorder144Iter 前序遍历迭代版
func Preorder144Iter(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}
