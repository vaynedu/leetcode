package interview

import "testing"

func TestFlatten(t *testing.T) {
    // 构造 1->(2,5)->(3,4)->6
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 5}
    root.Left.Left = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 4}
    root.Right.Right = &TreeNode{Val: 6}

    flatten(root)

    // 验证前序：1->2->3->4->5->6
    vals := []int{}
    cur := root
    for cur != nil {
        vals = append(vals, cur.Val)
        if cur.Right == nil && cur.Left != nil {
            t.Errorf("flatten: left not nil after flatten")
        }
        cur = cur.Right
    }
    want := []int{1, 2, 3, 4, 5, 6}
    if !intEq(vals, want) {
        t.Errorf("flatten() = %v, want %v", vals, want)
    }
}
