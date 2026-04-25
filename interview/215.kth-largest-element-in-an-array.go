package interview

// 215. 数组中的第K个最大元素
// 快速选择（Quick Select）：平均 O(n)，最坏 O(n²)
// 或者用堆：O(n log k) 建堆

// 快速选择实现
func findKthLargest(nums []int, k int) int {
	// 第 k 大 → 升序排名 n-k
	target := len(nums) - k
	return quickSelect(nums, 0, len(nums)-1, target)
}

func quickSelect(nums []int, left, right, target int) int {
	if left == right {
		return nums[left]
	}
	pivotIndex := partition(nums, left, right)
	if pivotIndex == target {
		return nums[pivotIndex]
	} else if pivotIndex < target {
		return quickSelect(nums, pivotIndex+1, right, target)
	} else {
		return quickSelect(nums, left, pivotIndex-1, target)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
