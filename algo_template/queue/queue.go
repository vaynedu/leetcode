package queue

import "container/list"

type Queue []int

func (q *Queue) Push(x int) {
	*q = append(*q, x)
}
func (q *Queue) Pop() int {
	if len(*q) == 0 {
		return -1
	}
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
func (q *Queue) Top() int {
	if len(*q) == 0 {
		return -1
	}
	return (*q)[0]
}

// levelOrderList 层次遍历
func levelOrderList(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		levelAns := make([]int, 0)
		for size > 0 {
			value := queue.Remove(queue.Front()).(*TreeNode)
			if value.Left != nil {
				queue.PushBack(value.Left)
			}
			if value.Right != nil {
				queue.PushBack(value.Right)
			}
			levelAns = append(levelAns, value.Val)
			size--
		}
		if len(levelAns) > 0 {
			ans = append(ans, levelAns)
		}
	}
	return ans
}

// levelOrderArray 数组版本
func levelOrderArray(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for size > 0 {
			value := queue[0]
			if value.Left != nil {
				queue = append(queue, value.Left)
			}
			if value.Right != nil {
				queue = append(queue, value.Right)
			}
			queue = queue[1:]
			level = append(level, value.Val)
			size--
		}
		if len(level) > 0 {
			ans = append(ans, level)
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
