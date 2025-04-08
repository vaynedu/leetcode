package leetcode

// 99.恢复二叉搜索树

import (
	"fmt"
	"testing"
)

/**
给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。



 示例 1：


输入：root = [1,3,null,null,2]
输出：[3,1,null,null,2]
解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。


 示例 2：


输入：root = [3,1,4,null,null,2]
输出：[2,1,4,null,null,3]
解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。



 提示：


 树上节点的数目在范围 [2, 1000] 内
 -2³¹ <= Node.val <= 2³¹ - 1




 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？

 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 995 👎 0

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

func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode

	// Helper function to perform in-order traversal
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil && node.Val < prev.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node
		inorder(node.Right)
	}

	// Perform in-order traversal to find the two nodes
	inorder(root)

	// Swap the values of the two nodes
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestRecoverBinarySearchTree(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
	tests := []struct {
		input  *TreeNode
		expect *TreeNode
	}{
		{
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			expect: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			expect: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	for _, test := range tests {
		recoverTree(test.input)
		if !isSameTree(test.input, test.expect) {
			t.Errorf("input: %v, expect: %v, actual: %v", test.input, test.expect, test.input)
		}
	}
}

// Helper function to check if two trees are the same
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
