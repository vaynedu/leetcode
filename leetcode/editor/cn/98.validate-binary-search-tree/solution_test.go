package leetcode

// 98.验证二叉搜索树

import (
	"math"
	"testing"
)

/**
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

 有效 二叉搜索树定义如下：


 节点的左子树只包含 小于 当前节点的数。
 节点的右子树只包含 大于 当前节点的数。
 所有左子树和右子树自身必须也是二叉搜索树。




 示例 1：


输入：root = [2,1,3]
输出：true


 示例 2：


输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。




 提示：


 树中节点数目范围在[1, 10⁴] 内
 -2³¹ <= Node.val <= 2³¹ - 1


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 2552 👎 0

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
// isValidBST 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

// helper 辅助函数，用于递归验证二叉搜索树
func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	// 如果当前节点值不在上下界范围内，返回 false
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	// 递归验证左子树，更新上界为当前节点值
	if !helper(root.Left, lower, root.Val) {
		return false
	}
	// 递归验证右子树，更新下界为当前节点值
	if !helper(root.Right, root.Val, upper) {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TestValidateBinarySearchTree 测试 isValidBST 函数
func TestValidateBinarySearchTree(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidBST(tt.root)
			if result != tt.expected {
				t.Errorf("isValidBST() = %v, want %v", result, tt.expected)
			}
		})
	}
}
