package binarytree

import "container/list"

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ==================== 递归遍历 ====================

// Preorder 前序遍历：根 → 左 → 右
func Preorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	Preorder(root.Left, result)
	Preorder(root.Right, result)
}

// Inorder 中序遍历：左 → 根 → 右
func Inorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	Inorder(root.Left, result)
	*result = append(*result, root.Val)
	Inorder(root.Right, result)
}

// Postorder 后序遍历：左 → 右 → 根
func Postorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	Postorder(root.Left, result)
	Postorder(root.Right, result)
	*result = append(*result, root.Val)
}

// ==================== 迭代遍历 ====================

// PreorderIter 前序迭代：根 → 左 → 右（模拟栈递归）
func PreorderIter(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		// 先压右子节点，再压左子节点（这样左子节点先出栈）
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

// InorderIter 中序迭代：左 → 根 → 右
func InorderIter(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		// 一路压左子树
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 弹出，访问根，再处理右子树
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}
	return result
}

// PostorderIter 后序迭代：左 → 右 → 根
// 技巧：前序逆序。先做 根→右→左，再 reverse
func PostorderIter(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		// 先压左子节点，再压右子节点（出栈顺序变成 根→右→左）
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// reverse：变成 左→右→根
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// PreorderList 前序遍历（用 container/list）
func PreorderList(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		node := l.Remove(l.Back()).(*TreeNode)
		result = append(result, node.Val)
		if node.Right != nil {
			l.PushBack(node.Right)
		}
		if node.Left != nil {
			l.PushBack(node.Left)
		}
	}
	return result
}
