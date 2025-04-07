package leetcode

// 783.二叉搜索树节点最小距离

import (
	"math"
	"testing"
)

/**
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

 差值是一个正数，其数值等于两值之差的绝对值。





 示例 1：


输入：root = [4,2,6,1,3]
输出：1




 示例 2：


输入：root = [1,0,48,null,null,12,49]
输出：1




 提示：


 树中节点的数目范围是 [2, 100]
 0 <= Node.val <= 10⁵




 注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
 相同

 Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 二叉树 👍 295 👎 0

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

// 全局变量用于记录前一个节点的值和最小差值
var prev int
var minDiff int

// 中序遍历辅助函数
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	// 递归遍历左子树
	inorder(root.Left)
	// 如果 prev 不为初始值，计算当前节点与前一个节点的差值
	if prev != -1 {
		minDiff = int(math.Min(float64(minDiff), float64(root.Val-prev)))
	}
	// 更新 prev 为当前节点的值
	prev = root.Val
	// 递归遍历右子树
	inorder(root.Right)
}

func minDiffInBST(root *TreeNode) int {
	// 初始化 prev 为 -1，表示还没有前一个节点
	prev = -1
	// 初始化最小差值为最大值
	minDiff = math.MaxInt32
	// 调用中序遍历函数
	inorder(root)
	return minDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// TestMinimumDistanceBetweenBstNodes 测试 minDiffInBST 函数
func TestMinimumDistanceBetweenBstNodes(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			expected: 1,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 48,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 49,
					},
				},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minDiffInBST(tt.root)
			if result != tt.expected {
				t.Errorf("minDiffInBST() = %d, want %d", result, tt.expected)
			}
		})
	}
}
