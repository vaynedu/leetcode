package interview

// 162. 寻找峰值
// 峰值是比相邻元素都大的元素，数组两端可以看作负无穷
// 时间 O(log N)，空间 O(1)
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
