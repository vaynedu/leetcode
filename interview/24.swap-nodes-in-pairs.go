package interview

// 24. 两两交换链表中的节点
// 迭代：三指针，每轮交换相邻两个
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    prev := dummy
    for prev.Next != nil && prev.Next.Next != nil {
        first := prev.Next
        second := first.Next
        // 交换
        first.Next = second.Next
        second.Next = first
        prev.Next = second
        // 移动到下一组
        prev = first
    }
    return dummy.Next
}
