package interview

// 96. 不同的二叉搜索树
// 动态规划：G[n] = sum(G[i-1] * G[n-i])，i 为根节点值
// G[n] = 第 n 个卡特兰数
func numTrees(n int) int {
    if n <= 1 {
        return 1
    }
    dp := make([]int, n+1)
    dp[0] = 1 // 空树
    dp[1] = 1 // 单节点

    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            dp[i] += dp[j-1] * dp[i-j]
        }
    }
    return dp[n]
}
