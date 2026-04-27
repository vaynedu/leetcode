package interview

import "testing"

func TestBuildTree(t *testing.T) {
    pre := []int{3, 9, 20, 15, 7}
    in := []int{9, 3, 15, 20, 7}
    root := buildTree(pre, in)
    if root == nil || root.Val != 3 {
        t.Errorf("buildTree() root.Val = %d, want 3", root.Val)
    }
    // 简单验证
    if root.Left != nil && root.Left.Val != 9 {
        t.Errorf("root.Left.Val = %d, want 9", root.Left.Val)
    }
    if root.Right != nil && root.Right.Val != 20 {
        t.Errorf("root.Right.Val = %d, want 20", root.Right.Val)
    }

    // 单节点
    root2 := buildTree([]int{1}, []int{1})
    if root2 == nil || root2.Val != 1 {
        t.Errorf("buildTree([1], [1]) failed")
    }
}
