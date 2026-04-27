package interview

// 104. 二叉树最大深度
// 递归：左右子树深度最大值 + 1
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    if left > right {
        return left + 1
    }
    return right + 1
}
