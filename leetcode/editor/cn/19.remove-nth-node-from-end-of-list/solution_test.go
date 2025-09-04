package leetcode

// 19.删除链表的倒数第 N 个结点

import (
	"fmt"
	"testing"
)

/**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。



 示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]


 示例 2：


输入：head = [1], n = 1
输出：[]


 示例 3：


输入：head = [1,2], n = 1
输出：[1]




 提示：


 链表中结点的数目为 sz
 1 <= sz <= 30
 0 <= Node.val <= 100
 1 <= n <= sz




 进阶：你能尝试使用一趟扫描实现吗？

 Related Topics 链表 双指针 👍 3174 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getLength(head *ListNode) int {
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{Val: 0, Next: head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)
type ListNode struct {
	Val  int
	Next *ListNode
}

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
