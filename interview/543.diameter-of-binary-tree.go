package interview

// 543. 二叉树的直径
// 后序遍历：直径 = 左深度 + 右深度，维护全局最大
func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := dfs(node.Left)
        right := dfs(node.Right)
        diameter = max(diameter, left+right)
        return max(left, right) + 1
    }
    dfs(root)
    return diameter
}
