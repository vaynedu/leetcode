package leetcode

import (
	"fmt"
	"testing"
)

// 404.左叶子之和

//给定二叉树的根节点 root ，返回所有左叶子之和。
//
//
// 示例 1：
//
//
//
//输入: root = [3,9,20,null,null,15,7]
//输出: 24
//
//
// 示例 2:
//
//
//输入: root = [1]
//
//
//
//
// 提示:
// 提示:
//
// 节点数在 [1, 1000] 范围内
// -1000 <= Node.val <= 1000
//
//
//
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 767 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)

}

// leetcode submit region end(Prohibit modification and deletion)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestSumOfLeftLeaves(t *testing.T) {
	fmt.Println("come on")
	// 生成函数测试用例
	// 生成函数测试用例
}
