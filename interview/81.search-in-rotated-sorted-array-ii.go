package interview

// SearchRotatedII 81. 搜索旋转排序数组 II
// 难度：Medium | 标签：数组 | 二分查找
// 核心思路：与33类似，但需处理 nums[left]==nums[mid]==nums[right] 导致无法判断哪半有序的情况，此时 left++ 收缩边界
// 时间：O(n)最坏 | 空间：O(1)
func SearchRotatedII(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		// 无法判断哪半有序时，收缩左边界
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
		} else if nums[left] <= nums[mid] { // 左半边有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右半边有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
