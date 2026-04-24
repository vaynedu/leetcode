package interview

// 1004. 最大连续1的个数 III（Max Consecutive Ones III）
// 变长窗口：记录 0 的数量，超过 K 时收缩

func longestOnes(nums []int, k int) int {
    L := 0
    zeroCount := 0
    maxLen := 0

    for R := 0; R < len(nums); R++ {
        if nums[R] == 0 {
            zeroCount++
        }

        for zeroCount > k {
            if nums[L] == 0 {
                zeroCount--
            }
            L++
        }

        if R-L+1 > maxLen {
            maxLen = R - L + 1
        }
    }

    return maxLen
}
