package interview

// FindMinII 154. 寻找旋转排序数组中的最小值 II
// 难度：Hard | 标签：数组 | 二分查找
// 核心思路：与153类似，但当 nums[mid] == nums[right] 时无法判断最小值在哪侧，此时 right-- 收缩
// 时间：O(n)最坏 | 空间：O(1)
func FindMinII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else { // nums[mid] == nums[right]，无法判断，收缩右边界
			right--
		}
	}
	return nums[left]
}
