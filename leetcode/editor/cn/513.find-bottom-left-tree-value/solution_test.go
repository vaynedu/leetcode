package leetcode

import (
	"fmt"
	"testing"
)

// 513.找树左下角的值

//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。
//
//
//
// 示例 1:
//
//
//
//
//输入: root = [2,1,3]
//输出: 1
//
//
// 示例 2:
//
//
//
//
//输入: [1,2,3,4,null,5,6,null,null,7]
//输出: 7
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [1,10⁴]
//
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 627 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func helper(root *TreeNode, depth int, maxDepth *int, maxNode **TreeNode) {
	if root == nil {
		return
	}

	if depth > *maxDepth {
		*maxDepth = depth
		*maxNode = root
	}

	helper(root.Left, depth+1, maxDepth, maxNode)
	helper(root.Right, depth+1, maxDepth, maxNode)
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 思路，寻找每个层最左边的节点的值，并且记录深度,一旦发现有更大的深度，就记录当前节点的值，并且更新最大深度
	maxDepth := 0
	var maxNode *TreeNode = root
	helper(root, 0, &maxDepth, &maxNode)
	return maxNode.Val
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestFindBottomLeftTreeValue(t *testing.T) {
	fmt.Println("come on")

	// Test cases
	testCases := []struct {
		root     *TreeNode
		expected int
	}{
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
			expected: 4,
		},
		{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: 1,
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
			expected: 5,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: 1,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
							Left: &TreeNode{
								Val: 8,
							},
						},
					},
				},
			},
			expected: 8,
		},
	}

	for _, tc := range testCases {
		result := findBottomLeftValue(tc.root)
		if result != tc.expected {
			t.Errorf("For root = %v, expected %v, got %v", tc.root, tc.expected, result)
		}
	}
}
