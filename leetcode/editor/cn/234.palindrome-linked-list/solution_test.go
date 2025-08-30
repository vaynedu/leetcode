package leetcode

// 234.回文链表

import (
	"fmt"
	"testing"
)

/**
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。



 示例 1：


输入：head = [1,2,2,1]
输出：true


 示例 2：


输入：head = [1,2]
输出：false




 提示：


 链表中节点数目在范围[1, 10⁵] 内
 0 <= Node.val <= 9




 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

 Related Topics 栈 递归 链表 双指针 👍 2110 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	vals := make([]int, 0)
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	for i := 0; i < len(vals)/2; i++ {
		if vals[i] != vals[len(vals)-i-1] {
			return false
		}
	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestPalindromeLinkedList(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
