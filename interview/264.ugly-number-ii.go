package interview

// 264. 丑数 II
// DP + 三指针：生成有序丑数序列
func nthUglyNumber(n int) int {
    dp := make([]int, n)
    dp[0] = 1
    p2, p3, p5 := 0, 0, 0
    for i := 1; i < n; i++ {
        v2, v3, v5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
        dp[i] = min(v2, min(v3, v5))
        if dp[i] == v2 {
            p2++
        }
        if dp[i] == v3 {
            p3++
        }
        if dp[i] == v5 {
            p5++
        }
    }
    return dp[n-1]
}
