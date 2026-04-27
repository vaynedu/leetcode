package interview

// 162. 寻找峰值
// 二分查找：找局部峰值，必有一侧递增
func findPeakElement(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] < nums[mid+1] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}
