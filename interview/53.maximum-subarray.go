package interview

// 53. 最大子数组和
// 动态规划：dp[i] = max(dp[i-1] + nums[i], nums[i])
// 空间优化：只保留前一个状态
func maxSubArray(nums []int) int {
    maxSum := nums[0]
    prev := nums[0]
    for i := 1; i < len(nums); i++ {
        prev = max(prev+nums[i], nums[i])
        maxSum = max(maxSum, prev)
    }
    return maxSum
}
