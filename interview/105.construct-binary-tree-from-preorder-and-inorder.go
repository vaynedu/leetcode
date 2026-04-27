package interview

// 105. 从前序与中序遍历序列构造二叉树
// 递归：前序首为根，在中序中定位，递归构建左右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{Val: preorder[0]}
    // 在中序中找根
    idx := 0
    for idx < len(inorder) && inorder[idx] != preorder[0] {
        idx++
    }
    root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
    root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
    return root
}
