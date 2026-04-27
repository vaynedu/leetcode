package interview

// 230. 二叉搜索树中第K小的元素
// 中序遍历：左根右，第 k-1 个即为第 k 小
func kthSmallest(root *TreeNode, k int) int {
    count := 0
    result := 0
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        count++
        if count == k {
            result = node.Val
            return
        }
        inorder(node.Right)
    }
    inorder(root)
    return result
}
