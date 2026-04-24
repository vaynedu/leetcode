package interview

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 二叉树层序遍历 - BFS + 队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// 1. 记录当前层节点数
		size := len(queue)

		// 2. 遍历当前层的所有节点
		level := []int{}
		for i := 0; i < size; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			// 左右子节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 3. 当前层结果加入
		result = append(result, level)
	}

	return result
}
