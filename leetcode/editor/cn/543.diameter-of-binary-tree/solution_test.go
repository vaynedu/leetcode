package leetcode

// 543.二叉树的直径

import (
	"fmt"
	. "leetcode/model"
	"testing"
)

/**
给你一棵二叉树的根节点，返回该树的 直径 。

 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

 两节点之间路径的 长度 由它们之间边数表示。



 示例 1：


输入：root = [1,2,3,4,5]
输出：3
解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。


 示例 2：


输入：root = [1,2]
输出：1




 提示：


 树中节点数目在范围 [1, 10⁴] 内
 -100 <= Node.val <= 100


 👍 1792 👎 0

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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 全局变量记录最大直径
var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0 // 初始化最大直径
	depth(root)
	return maxDiameter
}

// depth 计算以root为根的子树的最大深度，并更新直径
func depth(root *TreeNode) int {
	// 空节点深度为0
	if root == nil {
		return 0
	}

	// 递归计算左右子树的最大深度
	leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)

	// 当前节点的直径等于左子树深度+右子树深度
	currentDiameter := leftDepth + rightDepth

	// 更新全局最大直径
	if currentDiameter > maxDiameter {
		maxDiameter = currentDiameter
	}

	// 返回当前节点的最大深度
	return max(leftDepth, rightDepth) + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDiameterOfBinaryTree(t *testing.T) {
	// 测试用例1: [1,2,3,4,5]
	//     1
	//    / \
	//   2   3
	//  / \
	// 4   5
	// 期望输出: 3 (路径 4->2->1->3 或 5->2->1->3)
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}

	result1 := diameterOfBinaryTree(root1)
	if result1 != 3 {
		t.Errorf("测试用例1失败: 期望3, 实际%d", result1)
	}

	// 测试用例2: [1,2]
	//   1
	//  /
	// 2
	// 期望输出: 1 (路径 2->1)
	root2 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}

	result2 := diameterOfBinaryTree(root2)
	if result2 != 1 {
		t.Errorf("测试用例2失败: 期望1, 实际%d", result2)
	}

	fmt.Println("二叉树直径测试通过!")
}
