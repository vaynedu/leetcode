package interview

import "math"

// 322. 零钱兑换
// 完全背包：dp[i] = 最少硬币数凑成金额 i
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = math.MaxInt
    }
    dp[0] = 0
    for _, c := range coins {
        for i := c; i <= amount; i++ {
            if dp[i-c] != math.MaxInt {
                dp[i] = min(dp[i], dp[i-c]+1)
            }
        }
    }
    if dp[amount] == math.MaxInt {
        return -1
    }
    return dp[amount]
}
