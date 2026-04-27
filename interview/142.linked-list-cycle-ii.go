package interview

// 142. 环形链表 II
// 快慢指针 + 数学推导：相遇后从头同步走，再次相遇即为入口
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    slow := head
    fast := head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            // 相遇后，从头同步走，再次相遇即入口
            ptr := head
            for ptr != slow {
                ptr = ptr.Next
                slow = slow.Next
            }
            return ptr
        }
    }
    return nil
}
