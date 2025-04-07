package leetcode

// 109.有序链表转换二叉搜索树

import (
	"fmt"
	"testing"
)

/**
给定一个单链表的头节点 head ，其中的元素 按升序排序 ，将其转换为 平衡 二叉搜索树。



 示例 1:




输入: head = [-10,-3,0,5,9]
输出: [0,-3,9,-10,null,5]
解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。


 示例 2:


输入: head = []
输出: []




 提示:


 head 中的节点数在[0, 2 * 10⁴] 范围内
 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 二叉搜索树 链表 分治 二叉树 👍 939 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
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

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return helper(nums, 0, len(nums)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}
type ListNode struct {
	Next *ListNode
	Val  int
}

func TestConvertSortedListToBinarySearchTree(t *testing.T) {
	fmt.Println("come on baby !!!")
}
