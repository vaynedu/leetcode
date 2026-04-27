package interview

// 98. 验证二叉搜索树
// 中序遍历：左 < 根 < 右，利用前驱节点比较
func isValidBST(root *TreeNode) bool {
    var prev *int
    var inorder func(node *TreeNode) bool
    inorder = func(node *TreeNode) bool {
        if node == nil {
            return true
        }
        if !inorder(node.Left) {
            return false
        }
        if prev != nil && *prev >= node.Val {
            return false
        }
        prev = &node.Val
        return inorder(node.Right)
    }
    return inorder(root)
}
