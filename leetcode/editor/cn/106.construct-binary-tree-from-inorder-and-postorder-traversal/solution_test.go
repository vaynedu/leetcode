package leetcode

// 106.从中序与后序遍历序列构造二叉树

import (
	"fmt"
	"testing"
)

/**
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回
这颗 二叉树 。



 示例 1:


输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]


 示例 2:


输入：inorder = [-1], postorder = [-1]
输出：[-1]




 提示:


 1 <= inorder.length <= 3000
 postorder.length == inorder.length
 -3000 <= inorder[i], postorder[i] <= 3000
 inorder 和 postorder 都由 不同 的值组成
 postorder 中每一个值都在 inorder 中
 inorder 保证是树的中序遍历
 postorder 保证是树的后序遍历


 Related Topics 树 数组 哈希表 分治 二叉树 👍 1326 👎 0

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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			root.Left = buildTree(inorder[:i], postorder[:i])
			root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
			break
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func TestConstructBinaryTreeFromInorderAndPostorderTraversal(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
