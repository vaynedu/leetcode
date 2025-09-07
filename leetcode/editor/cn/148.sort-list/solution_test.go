package leetcode

// 148.排序链表

import (
	"fmt"
	"testing"
)

/**
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。






 示例 1：


输入：head = [4,2,1,3]
输出：[1,2,3,4]


 示例 2：


输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]


 示例 3：


输入：head = []
输出：[]




 提示：


 链表中节点的数目在范围 [0, 5 * 10⁴] 内
 -10⁵ <= Node.val <= 10⁵




 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

 Related Topics 链表 双指针 分治 排序 归并排序 👍 2606 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// sortList 使用归并排序对链表进行升序排序
func sortList(head *ListNode) *ListNode {
	// 如果链表为空或只有一个节点，则直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针找到链表中点，分割链表
	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 断开链表，分成两部分
	prev.Next = nil

	// 递归排序左右两部分
	left := sortList(head)
	right := sortList(slow)

	// 合并两个已排序的链表
	return merge(left, right)
}

// merge 合并两个有序链表
func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 哑节点，简化操作
	current := dummy

	// 比较两个链表的节点值，依次连接较小的节点
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// 将剩余节点连接到结果链表
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestSortList(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
