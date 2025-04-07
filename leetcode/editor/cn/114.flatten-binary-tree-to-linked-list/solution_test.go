package leetcode

// 114.二叉树展开为链表

import (
	"testing"
)

/**
给你二叉树的根结点 root ，请你将它展开为一个单链表：


 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
 展开后的单链表应该与二叉树 先序遍历 顺序相同。




 示例 1：


输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]


 示例 2：


输入：root = []
输出：[]


 示例 3：


输入：root = [0]
输出：[0]




 提示：


 树中结点数在范围 [0, 2000] 内
 -100 <= Node.val <= 100




 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？

 Related Topics 栈 树 深度优先搜索 链表 二叉树 👍 1822 👎 0

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
// 用于存储链表头节点
var prev *TreeNode

// flatten 将二叉树展开为链表
func flatten(root *TreeNode) {
	prev = nil
	flattenHelper(root)
}

// flattenHelper 辅助函数，用于递归展开二叉树
func flattenHelper(root *TreeNode) {
	if root == nil {
		return
	}
	// 递归处理右子树
	flattenHelper(root.Right)
	// 递归处理左子树
	flattenHelper(root.Left)
	// 将当前节点的右子树指向之前处理好的链表头
	root.Right = prev
	// 左子树置为 nil
	root.Left = nil
	// 更新链表头为当前节点
	prev = root
}

//leetcode submit region end(Prohibit modification and deletion)

//可以采用后序遍历的逆序来处理二叉树节点，具体步骤如下：
//先递归处理右子树。
//再递归处理左子树。
//把当前节点的右子树指向之前处理好的链表头，左子树置为 nil。
//更新链表头为当前节点。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TestFlattenBinaryTreeToLinkedList 测试 flatten 函数
func TestFlattenBinaryTreeToLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Example 2",
			root:     nil,
			expected: []int{},
		},
		{
			name: "Example 3",
			root: &TreeNode{
				Val: 0,
			},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.root)
			result := convertToList(tt.root)
			if !equal(result, tt.expected) {
				t.Errorf("flatten() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// convertToList 将展开后的链表转换为数组，方便测试
func convertToList(root *TreeNode) []int {
	var result []int
	for root != nil {
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}

// equal 比较两个整数切片是否相等
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
