package leetcode

// 在每个树行中找最大值

import (
	"fmt"
	"math"
	"testing"
)

/**
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。



 示例1：




输入: root = [1,3,2,5,3,null,9]
输出: [1,3,9]


 示例2：


输入: root = [1,2,3]
输出: [1,3]




 提示：


 二叉树的节点个数的范围是 [0,10⁴]

 -2³¹ <= Node.val <= 2³¹ - 1




 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 395 👎 0

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
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		levelMax := math.MinInt32
		for size > 0 {
			value := queue[0]
			if value.Left != nil {
				queue = append(queue, value.Left)
			}
			if value.Right != nil {
				queue = append(queue, value.Right)
			}
			queue = queue[1:]
			levelMax = max(levelMax, value.Val)
			size--
		}
		ans = append(ans, levelMax)

	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestFindLargestValueInEachTreeRow(t *testing.T) {
	fmt.Println("come on baby !!!")
}
