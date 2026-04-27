package interview

// 278. 第一个错误的版本
// 二分查找：isBadVersion(mid) 为 true 则答案在左，否则在右
// 注意：isBadVersion 由 LeetCode 平台提供，这里用 APIVersion 模拟
func firstBadVersion(n int, badVersion int) int {
    left, right := 1, n
    for left < right {
        mid := left + (right-left)/2
        if mid >= badVersion {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}
