package interview

// 62. 不同路径
// 动态规划：dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 空间优化：一行数组
func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    for j := 0; j < n; j++ {
        dp[j] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[j] += dp[j-1]
        }
    }
    return dp[n-1]
}
