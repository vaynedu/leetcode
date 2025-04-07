package leetcode

// 108.将有序数组转换为二叉搜索树

import (
	"fmt"
	"testing"
)

/**
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。



 示例 1：


输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：



 示例 2：


输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。




 提示：


 1 <= nums.length <= 10⁴
 -10⁴ <= nums[i] <= 10⁴
 nums 按 严格递增 顺序排列


 Related Topics 树 二叉搜索树 数组 分治 二叉树 👍 1641 👎 0

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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return dfs(nums, 0, len(nums))
}

func dfs(nums []int, start int, end int) *TreeNode {
	if start >= end {
		return nil
	}

	mix := (start + end) / 2

	var node *TreeNode = new(TreeNode)
	node.Val = nums[mix]

	left := dfs(nums, start, mix)
	right := dfs(nums, mix+1, end)

	node.Left = left
	node.Right = right

	return node
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestConvertSortedArrayToBinarySearchTree(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
