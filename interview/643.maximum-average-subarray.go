package interview

// 643. 子数组最大平均数（Maximum Average Subarray I）
// 固定窗口：先算首窗和，右移减左加右

func findMaxAverage(nums []int, k int) float64 {
    // 第一个窗口的和
    sum := 0
    for i := 0; i < k; i++ {
        sum += nums[i]
    }
    maxSum := sum

    // 滑动窗口：右移时减左加右
    for i := k; i < len(nums); i++ {
        sum += nums[i] - nums[i-k]
        if sum > maxSum {
            maxSum = sum
        }
    }

    return float64(maxSum) / float64(k)
}
