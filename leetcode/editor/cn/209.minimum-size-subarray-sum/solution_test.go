package leetcode

// 209.长度最小的子数组

import (
	"math"
	"testing"
)

/**
给定一个含有 n 个正整数的数组和一个正整数 target 。

 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度
。如果不存在符合条件的子数组，返回 0 。



 示例 1：


输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。


 示例 2：


输入：target = 4, nums = [1,4,4]
输出：1


 示例 3：


输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0




 提示：


 1 <= target <= 10⁹
 1 <= nums.length <= 10⁵
 1 <= nums[i] <= 10⁴




 进阶：


 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。


 Related Topics 数组 二分查找 前缀和 滑动窗口 👍 2530 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLenV1OverLimit(target int, nums []int) int {
	// 思路， 前缀和
	// 遍历，查找最短数组

	sums := make([]int, len(nums)+1)
	sums[0] = 0
	for i := 0; i < len(nums); i++ {
		sums[i+1] += sums[i] + nums[i]
	}

	if sums[len(sums)-1] < target {
		return 0
	}

	ans := math.MaxInt32 // 按照题目中要求最大值赋值
	for i := len(sums) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if sums[i]-sums[j] >= target {
				ans = min(ans, i-j)
				break
			}
		}
	}
	return ans
}

func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口解法 - 时间复杂度 O(n)
	left := 0
	sum := 0
	minLength := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		// 扩大窗口，将右边界元素加入窗口
		sum += nums[right]

		// 当窗口内元素和满足条件时，尝试缩小窗口
		for sum >= target {
			// 更新最小长度
			minLength = min(minLength, right-left+1)
			// 缩小窗口，移除左边界元素
			sum -= nums[left]
			left++
		}
	}

	// 如果没有找到满足条件的子数组，返回0
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumSizeSubarraySum(t *testing.T) {
	tests := []struct {
		target   int
		nums     []int
		expected int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{15, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}, 2},
	}

	for i, test := range tests {
		result := minSubArrayLen(test.target, test.nums)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %d, got %d", i+1, test.expected, result)
		} else {
			t.Logf("Test case %d passed: target=%d, nums=%v, result=%d",
				i+1, test.target, test.nums, result)
		}
	}
}
