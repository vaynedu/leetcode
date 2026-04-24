package interview

// SearchRange 34. 在排序数组中查找元素的第一个和最后一个位置
// 难度：Medium | 标签：数组 | 二分查找
// 核心思路：两次二分查找——第一次找左侧边界，第二次找右侧边界
// 时间：O(log n) | 空间：O(1)
func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 找左侧边界：第一个 >= target 的位置
	left := lowerBound(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	// 找右侧边界：第一个 > target 的位置 - 1
	right := lowerBound(nums, target+1) - 1
	return []int{left, right}
}

// lowerBound 返回第一个 >= target 的下标（左闭右开区间）
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
