package leetcode

// 完全二叉树的节点个数

import (
	"fmt"
	"testing"
)

/**
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

 完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第
 h 层（从第 0 层开始），则该层包含 1~ 2ʰ 个节点。



 示例 1：


输入：root = [1,2,3,4,5,6]
输出：6


 示例 2：


输入：root = []
输出：0


 示例 3：


输入：root = [1]
输出：1




 提示：


 树中节点的数目范围是[0, 5 * 10⁴]
 0 <= Node.val <= 5 * 10⁴
 题目数据保证输入的树是 完全二叉树




 进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？

 Related Topics 位运算 树 二分查找 二叉树 👍 1234 👎 0

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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := root.Left
	leftDepth := 0
	for left != nil {
		left = left.Left
		leftDepth++
	}
	right := root.Right
	rightDepth := 0
	for right != nil {
		right = right.Right
		rightDepth++
	}

	if leftDepth == rightDepth {
		return (1 << (leftDepth + 1)) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func TestCountCompleteTreeNodes(t *testing.T) {
	// 测试用例
	testCases := []struct {
		root     *TreeNode
		expected int
	}{
		{&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5}},
			Right: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 6}}},
			6},
		{nil, 0},
		{&TreeNode{Val: 1}, 1},
	}

	// 执行测试
	for _, tc := range testCases {
		result := countNodes(tc.root)
		if result != tc.expected {
			t.Errorf("对于输入 %v, 期望输出 %d, 实际输出 %d", tc.root, tc.expected, result)
		} else {
			fmt.Printf("测试通过: 输入 %v, 输出 %d\n", tc.root, result)
		}
	}
}
