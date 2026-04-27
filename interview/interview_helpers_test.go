package interview

// 测试辅助函数（同一 package 下所有 _test.go 共享）

func intEq(a, b []int) bool {
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

func intSliceEq(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !intEq(a[i], b[i]) {
			return false
		}
	}
	return true
}

// TreeNode 辅助函数
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// NewTree 从层序数组构建二叉树（nil 表示空节点）
func NewTree(vals []any) *TreeNode {
    if len(vals) == 0 || vals[0] == nil {
        return nil
    }
    root := &TreeNode{Val: vals[0].(int)}
    queue := []*TreeNode{root}
    i := 1
    for len(queue) > 0 && i < len(vals) {
        node := queue[0]
        queue = queue[1:]
        if i < len(vals) && vals[i] != nil {
            node.Left = &TreeNode{Val: vals[i].(int)}
            queue = append(queue, node.Left)
        }
        i++
        if i < len(vals) && vals[i] != nil {
            node.Right = &TreeNode{Val: vals[i].(int)}
            queue = append(queue, node.Right)
        }
        i++
    }
    return root
}

// ListNode 辅助函数
type ListNode struct {
    Val  int
    Next *ListNode
}

func NewList(vals ...int) *ListNode {
    if len(vals) == 0 {
        return nil
    }
    dummy := &ListNode{}
    cur := dummy
    for _, v := range vals {
        cur.Next = &ListNode{Val: v}
        cur = cur.Next
    }
    return dummy.Next
}

func listToSlice(head *ListNode) []int {
    var result []int
    for head != nil {
        result = append(result, head.Val)
        head = head.Next
    }
    return result
}

func listEqual(head *ListNode, vals []int) bool {
    slice := listToSlice(head)
    return intEq(slice, vals)
}

func strEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if !m[v] {
			return false
		}
	}
	return true
}
