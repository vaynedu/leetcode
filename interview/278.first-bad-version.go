package interview

// 278. 第一个错误的版本
// 二分查找：isBadVersion(mid) 为 true 则答案在左，否则在右
func firstBadVersion(n int) int {
    left, right := 1, n
    for left < right {
        mid := left + (right-left)/2
        if isBadVersion(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}
