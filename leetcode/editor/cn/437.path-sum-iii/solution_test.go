package leetcode

// 437.路径总和 III

import (
	. "leetcode/model"
	"testing"
)

/**
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。



 示例 1：




输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
输出：3
解释：和等于 8 的路径有 3 条，如图所示。


 示例 2：


输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：3




 提示:


 二叉树的节点个数的范围是 [0,1000]

 -10⁹ <= Node.val <= 10⁹
 -1000 <= targetSum <= 1000


 👍 2132 👎 0

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
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathSumStartWithRoot(root, targetSum) +
		pathSum(root.Left, targetSum) +
		pathSum(root.Right, targetSum)
}

func pathSumStartWithRoot(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.Val == targetSum {
		count++
	}
	count += pathSumStartWithRoot(root.Left, targetSum-root.Val)
	count += pathSumStartWithRoot(root.Right, targetSum-root.Val)
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPathSumIii(t *testing.T) {
	// 构建测试用例1: [10,5,-3,3,2,null,11,3,-2,null,1]
	//       10
	//      /  \
	//     5   -3
	//    / \    \
	//   3   2   11
	//  / \ /
	// 3  -2 1
	root1 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: -2},
			},
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val:   -3,
			Right: &TreeNode{Val: 11},
		},
	}

	// 测试用例2: [5,4,8,11,null,13,4,7,2,null,null,5,1]
	//         5
	//       /  \
	//      4    8
	//     /    / \
	//    11   13  4
	//   / \      / \
	//  7   2    5   1
	root2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},
	}

	// 测试用例3: 空树
	var root3 *TreeNode

	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		expected  int
	}{
		{"示例1", root1, 8, 3},
		{"示例2", root2, 22, 3},
		{"空树", root3, 1, 0},
		{"单节点匹配", &TreeNode{Val: 1}, 1, 1},
		{"单节点不匹配", &TreeNode{Val: 1}, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pathSum(tt.root, tt.targetSum)
			if result != tt.expected {
				t.Errorf("pathSum(%v, %d) = %d, expected %d", tt.root, tt.targetSum, result, tt.expected)
			}
		})
	}
}
