package interview

// 141. 环形链表
// 快慢指针：若有环，快指针必追上慢指针
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    slow := head
    fast := head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
