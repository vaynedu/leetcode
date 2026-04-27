package interview

// 198. 房屋盗贼
// 动态规划：dp[i] = max(dp[i-1], dp[i-2] + nums[i])
// 空间优化：只保留前两个状态
func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    prev2, prev1 := nums[0], max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        cur := max(prev1, prev2+nums[i])
        prev2 = prev1
        prev1 = cur
    }
    return prev1
}
