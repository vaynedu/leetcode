package interview

// 278. 第一个错误的版本
// 左边界模板：找第一个 IsBadVersion(mid) == true
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

// isBadVersion 是官方提供的接口
func isBadVersion(version int) bool {
    // 模拟：假设版本 5 是第一个错误版本
    return version >= 5
}
