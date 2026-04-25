package interview

import "testing"

func TestMergeKLists(t *testing.T) {
	// 构造测试链表
	toList := func(vals []int) *ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &ListNode{Val: vals[0]}
		curr := head
		for i := 1; i < len(vals); i++ {
			curr.Next = &ListNode{Val: vals[i]}
			curr = curr.Next
		}
		return head
	}
	toSlice := func(head *ListNode) []int {
		var r []int
		for head != nil {
			r = append(r, head.Val)
			head = head.Next
		}
		return r
	}

	tests := []struct {
		name   string
		lists  [][]int
		expect []int
	}{
		{"两个链表", [][]int{{1, 4, 5}, {1, 3, 4}}, []int{1, 1, 3, 4, 4, 5}},
		{"三个链表", [][]int{{1, 3, 5}, {1, 4, 6}, {2}}, []int{1, 1, 2, 3, 4, 5, 6}},
		{"空链表", [][]int{{}, {}}, []int{}},
		{"含nil", [][]int{{1, 2}, {}, {3}}, []int{1, 2, 3}},
		{"一个链表", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lists []*ListNode
			for _, vals := range tt.lists {
				lists = append(lists, toList(vals))
			}
			got := mergeKLists(lists)
			gotSlice := toSlice(got)
			if !intSliceEqual(gotSlice, tt.expect) {
				t.Errorf("mergeKLists() = %v, want %v", gotSlice, tt.expect)
			}
		})
	}
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
