package interview

// 63. 不同路径 II
// 动态规划：有障碍物时 dp[j] = 0（障碍在当前位置）
// 处理第一行/列时遇到障碍则后续都是0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    dp := make([]int, n)

    for j := 0; j < n; j++ {
        if obstacleGrid[0][j] == 1 {
            dp[j] = 0
        } else if j == 0 {
            dp[j] = 1
        } else {
            dp[j] = dp[j-1]
        }
    }

    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                dp[j] = 0
            } else if j > 0 {
                dp[j] += dp[j-1]
            }
        }
    }
    return dp[n-1]
}
