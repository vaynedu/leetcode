package interview

// 367. 有效的完全平方数
// 二分查找：mid*mid 与 num 比较，注意溢出用除法
func isPerfectSquare(num int) bool {
    if num < 2 {
        return true
    }
    left, right := 2, num/2
    for left <= right {
        mid := left + (right-left)/2
        if mid == num/mid && num%mid == 0 {
            return true
        } else if mid > num/mid {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return false
}
