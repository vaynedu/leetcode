package interview

// 367. 有效的完全平方数
// 使用 x/mid 判断，避免溢出
func isPerfectSquare(num int) bool {
    if num < 2 {
        return true
    }
    left, right := 1, num/2
    for left <= right {
        mid := left + (right-left)/2
        if mid <= num/mid {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left*left == num || right*right == num
}
