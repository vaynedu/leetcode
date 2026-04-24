package interview

// 153. 寻找旋转排序数组中的最小值
// 右半边有序时 nums[mid] > nums[right]，最小值在左半边
func findMin153(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return nums[left]
}
