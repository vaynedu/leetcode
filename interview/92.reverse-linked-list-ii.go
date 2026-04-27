package interview

// 92. 反转链表 II
// 头插法：遍历到 left-1，然后头插法反转 [left, right]
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    prev := dummy
    // 1. 找到 left 的前一个节点
    for i := 0; i < left-1; i++ {
        prev = prev.Next
    }
    // 2. 开始反转，用头插法
    cur := prev.Next
    for i := 0; i < right-left; i++ {
        next := cur.Next
        cur.Next = next.Next
        next.Next = prev.Next
        prev.Next = next
    }
    return dummy.Next
}
