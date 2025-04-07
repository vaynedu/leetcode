package leetcode

// 94.二叉树的中序遍历

import (
	"fmt"
	"testing"
)

/**
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。



 示例 1：


输入：root = [1,null,2,3]
输出：[1,3,2]


 示例 2：


输入：root = []
输出：[]


 示例 3：


输入：root = [1]
输出：[1]




 提示：


 树中节点数目在范围 [0, 100] 内
 -100 <= Node.val <= 100




 进阶: 递归算法很简单，你可以通过迭代算法完成吗？

 Related Topics 栈 树 深度优先搜索 二叉树 👍 2228 👎 0

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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		ans = append(ans, node.Val)
		traversal(node.Right)
	}

	traversal(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func TestBinaryTreeInorderTraversal(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
