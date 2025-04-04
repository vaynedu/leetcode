package leetcode

// 二叉树的层序遍历 II

import (
	"container/list"
	"reflect"
	"testing"
)

/**
给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[15,7],[9,20],[3]]


 示例 2：


输入：root = [1]
输出：[[1]]


 示例 3：


输入：root = []
输出：[]




 提示：


 树中节点数目在范围 [0, 2000] 内
 -1000 <= Node.val <= 1000


 Related Topics 树 广度优先搜索 二叉树 👍 832 👎 0

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
func levelOrderBottom(root *TreeNode) [][]int {
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
			// ans = append(ans, levelAns)
			ans = append([][]int{levelAns}, ans...)
		}
	}

	//// 反转ans二维数组
	//for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
	//	ans[i], ans[j] = ans[j], ans[i]
	//}

	return ans

}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestBinaryTreeLevelOrderTraversalIi(t *testing.T) {
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
			expected: [][]int{{15, 7}, {9, 20}, {3}},
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
			expected: [][]int{{4, 5, 6, 7}, {2, 3}, {1}},
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
			expected: [][]int{{5}, {4}, {3}, {2}, {1}},
		},
	}

	for _, test := range tests {
		actual := levelOrderBottom(test.root)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("For root %v, expected %v, but got %v", test.root, test.expected, actual)
		}
	}
}
