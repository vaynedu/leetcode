package interview

// 494. 目标和
// 转化：正数和 - 负数和 = S，记正数和为 P，负数和为 N
// P - N = S, P + N = sum => P = (sum + S) / 2
// 问题变成：从 nums 中选若干元素凑成 P
func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if sum < 0 {
        sum = -sum
    }
    if target < 0 {
        target = -target
    }
    if target > sum || (sum+target)%2 != 0 {
        return 0
    }
    P := (sum + target) / 2
    dp := make([]int, P+1)
    dp[0] = 1
    for _, v := range nums {
        for i := P; i >= v; i-- {
            dp[i] += dp[i-v]
        }
    }
    return dp[P]
}
