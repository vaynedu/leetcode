package interview

// 120. 三角形最小路径和
// 动态规划：从底部向上，dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
// 空间优化：从底部往上只用一行的空间
func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    dp := make([]int, n+1) // 多一个元素避免边界判断

    // 从最后一行开始往上
    for i := n - 1; i >= 0; i-- {
        for j := 0; j <= i; j++ {
            dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
        }
    }
    return dp[0]
}
