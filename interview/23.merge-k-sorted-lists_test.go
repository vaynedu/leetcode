package interview

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	fmt.Println("=============================================")
	fmt.Println("       23. 合并 K 个升序链表")
	fmt.Println("=============================================")
	fmt.Println()

	// 构建测试链表 [1->4->5]
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 5}

	// 构建测试链表 [1->3->4]
	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	// 构建测试链表 [2->6]
	l3 := &ListNode{Val: 2}
	l3.Next = &ListNode{Val: 6}

	lists := []*ListNode{l1, l2, l3}
	result := mergeKLists(lists)

	// 验证结果
	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}
	for _, v := range expected {
		if result == nil || result.Val != v {
			t.Errorf("mergeKLists failed, got %v", result)
		}
		if result != nil {
			result = result.Next
		}
	}

	// 空数组测试
	if mergeKLists([]*ListNode{}) != nil {
		t.Errorf("empty lists should return nil")
	}

	// 全 nil 测试
	if mergeKLists([]*ListNode{nil, nil}) != nil {
		t.Errorf("all nil should return nil")
	}

	fmt.Println("✅ mergeKLists - 分治归并 - PASS")
}
