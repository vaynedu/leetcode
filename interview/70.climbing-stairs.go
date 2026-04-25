package interview

// 70. 爬楼梯
// 动态规划：dp[i] = dp[i-1] + dp[i-2]
// 斐波那契数列，空间优化到 O(1)
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    prev2, prev1 := 1, 2
    for i := 3; i <= n; i++ {
        curr := prev1 + prev2
        prev2, prev1 = prev1, curr
    }
    return prev1
}
