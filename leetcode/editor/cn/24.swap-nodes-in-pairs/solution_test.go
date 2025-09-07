package leetcode

// 24.两两交换链表中的节点

import (
	"fmt"
	"testing"
)

/**
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。



 示例 1：


输入：head = [1,2,3,4]
输出：[2,1,4,3]


 示例 2：


输入：head = []
输出：[]


 示例 3：


输入：head = [1]
输出：[1]




 提示：


 链表中节点的数目在范围 [0, 100] 内
 0 <= Node.val <= 100


 Related Topics 递归 链表 👍 2468 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	q := head.Next

	// 1 2 3
	// p q
	p.Next = swapPairs(q.Next)
	q.Next = p

	return q
}

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestSwapNodesInPairs(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
