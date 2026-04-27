package interview

// 300. 最长递增子序列
// DP+二分：tails[i] = 长度为 i+1 的递增子序列的最小尾部值
func lengthOfLIS(nums []int) int {
    tails := make([]int, 0)
    for _, v := range nums {
        i := 0
        for i < len(tails) && tails[i] < v {
            i++
        }
        if i == len(tails) {
            tails = append(tails, v)
        } else {
            tails[i] = v
        }
    }
    return len(tails)
}
