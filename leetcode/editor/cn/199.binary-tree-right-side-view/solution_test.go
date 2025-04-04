package leetcode

// 二叉树的右视图

import (
	"fmt"
	"testing"
)

/**
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。



 示例 1：


 输入：root = [1,2,3,null,5,null,4]


 输出：[1,3,4]

 解释：



 示例 2：


 输入：root = [1,2,3,4,null,null,null,5]


 输出：[1,3,4,5]

 解释：



 示例 3：


 输入：root = [1,null,3]


 输出：[1,3]

 示例 4：


 输入：root = []


 输出：[]



 提示:


 二叉树的节点个数的范围是 [0,100]

 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1188 👎 0

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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for size > 0 {
			value := queue[0]
			if value.Left != nil {
				queue = append(queue, value.Left)
			}
			if value.Right != nil {
				queue = append(queue, value.Right)
			}
			queue = queue[1:]
			level = append(level, value.Val)
			size--
		}
		if len(level) > 0 {
			ans = append(ans, level[len(level)-1])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestBinaryTreeRightSideView(t *testing.T) {
	fmt.Println("come on baby !!!")
}
