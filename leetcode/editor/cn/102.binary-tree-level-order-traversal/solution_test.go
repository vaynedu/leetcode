package leetcode

// 二叉树的层序遍历

import (
	"reflect"
	"testing"
)

/**
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]


 示例 2：


输入：root = [1]
输出：[[1]]


 示例 3：


输入：root = []
输出：[]




 提示：


 树中节点数目在范围 [0, 2000] 内
 -1000 <= Node.val <= 1000


 Related Topics 树 广度优先搜索 二叉树 👍 2114 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
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

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected [][]int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: [][]int{{1}},
		},
		{
			root:     nil,
			expected: [][]int{},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
				},
			},
			expected: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
	}

	for _, test := range tests {
		actual := levelOrder(test.root)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("For root %v, expected %v, but got %v", test.root, test.expected, actual)
		}
	}
}
