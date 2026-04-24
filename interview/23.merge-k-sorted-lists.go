package interview

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists 合并K个升序链表 - 分治归并
// 时间：O(N log K)，空间：O(log K) 递归栈
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeSort(lists, 0, len(lists)-1)
}

func mergeSort(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	// 防止整数溢出
	mid := left + (right-left)/2
	// 递归左右两半
	l1 := mergeSort(lists, left, mid)
	l2 := mergeSort(lists, mid+1, right)
	// 合并两个有序链表
	return mergeTwoLists(l1, l2)
}

// mergeTwoLists 合并两个有序链表
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}
