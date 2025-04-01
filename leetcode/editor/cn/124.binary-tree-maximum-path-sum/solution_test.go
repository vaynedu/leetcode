package leetcode

// 二叉树中的最大路径和

import (
	"fmt"
	"math"
	"testing"
)

//二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定
//经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
//
// 示例 2：
//
//
//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
//
//
//
// 提示：
//
//
// 树中节点数目范围是 [1, 3 * 10⁴]
// -1000 <= Node.val <= 1000
//
//
// Related Topics 树 深度优先搜索 动态规划 二叉树 👍 2379 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := max(0, dfs(root.Left, ans))
	right := max(0, dfs(root.Right, ans))

	*ans = max(*ans, root.Val+left+right)
	return max(root.Val+max(left, right), 0)
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	dfs(root, &ans)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestBinaryTreeMaximumPathSum(t *testing.T) {
	fmt.Println("come on baby !!!")
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  5,
			Left: nil,
		},
	}
	fmt.Println(maxPathSum(root)) // Expected output: 16 (1 + 2 + 4 + 5 + 3)

	// Test case 1: Single node
	root1 := &TreeNode{
		Val: -10,
	}
	fmt.Println(maxPathSum(root1)) // Expected output: -10

	// Test case 2: Example from the problem statement
	root2 := &TreeNode{
		Val: -10,
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
	}
	fmt.Println(maxPathSum(root2)) // Expected output: 42 (15 + 20 + 7)

	// Test case 3: All negative values
	root3 := &TreeNode{
		Val: -3,
		Left: &TreeNode{
			Val: -9,
			Left: &TreeNode{
				Val: -15,
			},
			Right: &TreeNode{
				Val: -20,
			},
		},
		Right: &TreeNode{
			Val: -7,
			Left: &TreeNode{
				Val: -15,
			},
			Right: &TreeNode{
				Val: -7,
			},
		},
	}
	fmt.Println(maxPathSum(root3)) // Expected output: -3

	// Test case 4: Mixed positive and negative values
	root4 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: -1,
		},
		Right: &TreeNode{
			Val: -2,
		},
	}
	fmt.Println(maxPathSum(root4)) // Expected output: 2

	// Test case 5: Degenerate tree (linked list)
	root5 := &TreeNode{
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
	}
	fmt.Println(maxPathSum(root5)) // Expected output: 15 (1 + 2 + 3 + 4 + 5)
}
