package interview

// 64. 最小路径和
// 动态规划：dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
// 原地修改第一行/列
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])

    // 第一行：只能从左边来
    for j := 1; j < n; j++ {
        grid[0][j] += grid[0][j-1]
    }
    // 第一列：只能从上边来
    for i := 1; i < m; i++ {
        grid[i][0] += grid[i-1][0]
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            grid[i][j] += min(grid[i-1][j], grid[i][j-1])
        }
    }
    return grid[m-1][n-1]
}
