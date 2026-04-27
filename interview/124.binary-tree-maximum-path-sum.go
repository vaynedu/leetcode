package interview

// 124. 二叉树中的最大路径和
// 后序遍历：每个节点计算"向下"的最大贡献，返回 max(0, 左)+根+max(0, 右)
func maxPathSum(root *TreeNode) int {
    maxSum := -1 << 63 // 负无穷
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := dfs(node.Left)
        right := dfs(node.Right)
        // 以当前节点为拐点的路径
        pathSum := node.Val + max(0, left) + max(0, right)
        maxSum = max(maxSum, pathSum)
        // 返回"向下"的最大贡献（只能选一边或都不选）
        return node.Val + max(0, max(left, right))
    }
    dfs(root)
    return maxSum
}
