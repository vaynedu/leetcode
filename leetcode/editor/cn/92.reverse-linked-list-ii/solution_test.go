package leetcode

// 92.反转链表 II

import (
	. "leetcode/model"
	"testing"
)

/**
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节
点，返回 反转后的链表 。



 示例 1：


输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]


 示例 2：


输入：head = [5], left = 1, right = 1
输出：[5]




 提示：


 链表中节点数目为 n
 1 <= n <= 500
 -500 <= Node.val <= 500
 1 <= left <= right <= n




 进阶： 你可以使用一趟扫描完成反转吗？

 Related Topics 链表 👍 2023 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 创建虚拟头节点，简化边界情况处理
	dummy := &ListNode{Next: head}

	// 找到反转区间的前一个节点
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// current指向反转区间的第一个节点
	current := prev.Next

	// 进行反转操作
	// 将left到right之间的节点逐个插入到prev后面
	for i := 0; i < right-left; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedListIi(t *testing.T) {
	// 测试用例1: [1,2,3,4,5], left=2, right=4 => [1,4,3,2,5]
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = &ListNode{Val: 4}
	head1.Next.Next.Next.Next = &ListNode{Val: 5}

	expected1 := []int{1, 4, 3, 2, 5}
	result1 := reverseBetween(head1, 2, 4)

	if !isEqualList(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, listToSlice(result1))
	}

	// 测试用例2: [5], left=1, right=1 => [5]
	head2 := &ListNode{Val: 5}
	expected2 := []int{5}
	result2 := reverseBetween(head2, 1, 1)

	if !isEqualList(result2, expected2) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected2, listToSlice(result2))
	}

	// 测试用例3: [1,2,3], left=1, right=3 => [3,2,1]
	head3 := &ListNode{Val: 1}
	head3.Next = &ListNode{Val: 2}
	head3.Next.Next = &ListNode{Val: 3}
	expected3 := []int{3, 2, 1}
	result3 := reverseBetween(head3, 1, 3)

	if !isEqualList(result3, expected3) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected3, listToSlice(result3))
	}
}

// 辅助函数：将链表转换为切片
func listToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

// 辅助函数：比较链表和切片是否相等
func isEqualList(head *ListNode, expected []int) bool {
	current := head
	for _, val := range expected {
		if current == nil || current.Val != val {
			return false
		}
		current = current.Next
	}
	return current == nil
}
