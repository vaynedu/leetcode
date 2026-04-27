package interview

// 416. 分割等和子集
// 子集背包：是否能凑出 sum/2
func canPartition(nums []int) bool {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if sum%2 != 0 {
        return false
    }
    target := sum / 2
    dp := make([]bool, target+1)
    dp[0] = true
    for _, v := range nums {
        for i := target; i >= v; i-- {
            dp[i] = dp[i] || dp[i-v]
        }
    }
    return dp[target]
}
