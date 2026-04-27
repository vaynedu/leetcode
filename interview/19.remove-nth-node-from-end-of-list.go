package interview

// 19. 删除链表的倒数第 N 个节点
// 快慢指针：先走 n 步，然后一起走
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    fast := dummy
    slow := dummy
    // fast 先走 n+1 步
    for i := 0; i <= n; i++ {
        if fast == nil {
            return head
        }
        fast = fast.Next
    }
    // 一起走，直到 fast 到达末尾
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    // slow.Next 即为要删除的节点
    slow.Next = slow.Next.Next
    return dummy.Next
}
