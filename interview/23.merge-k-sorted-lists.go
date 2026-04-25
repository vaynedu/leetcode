package interview

import "container/heap"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 23. 合并K个排序链表
// 最小堆：把所有链表头加入堆，每次弹出最小的接入结果

// MinHeapItem 是堆中存储的元素
type MinHeapItem struct {
	node  *ListNode
	heapIdx int // 堆中原始链表下标，用于值相同时分辨
}

// MinHeap 实现 container/heap 的接口
type MinHeap []*MinHeapItem

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].node.Val != h[j].node.Val {
		return h[i].node.Val < h[j].node.Val
	}
	return h[i].heapIdx < h[j].heapIdx
}
func (h *MinHeap) Swap(i, j int)       { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Push(x interface{})   { *h = append(*h, x.(*MinHeapItem)) }
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)

	// 初始化：把所有链表头加入堆（跳过 nil）
	for i, node := range lists {
		if node != nil {
			heap.Push(h, &MinHeapItem{node: node, heapIdx: i})
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for h.Len() > 0 {
		item := heap.Pop(h).(*MinHeapItem)
		curr.Next = item.node
		curr = curr.Next
		// 弹出的节点有后继的话，后继入堆
		if item.node.Next != nil {
			heap.Push(h, &MinHeapItem{node: item.node.Next, heapIdx: item.heapIdx})
		}
	}

	return dummy.Next
}
