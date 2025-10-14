package leetcode

// 3349.检测相邻递增子数组 I

import (
	"testing"
)

/**
给你一个由 n 个整数组成的数组 nums 和一个整数 k，请你确定是否存在 两个 相邻 且长度为 k 的 严格递增 子数组。具体来说，需要检查是否存在从下标
a 和 b (a < b) 开始的 两个 子数组，并满足下述全部条件：


 这两个子数组 nums[a..a + k - 1] 和 nums[b..b + k - 1] 都是 严格递增 的。
 这两个子数组必须是 相邻的，即 b = a + k。


 如果可以找到这样的 两个 子数组，请返回 true；否则返回 false。

 子数组 是数组中的一个连续 非空 的元素序列。



 示例 1：


 输入：nums = [2,5,7,8,9,2,3,4,3,1], k = 3


 输出：true

 解释：


 从下标 2 开始的子数组为 [7, 8, 9]，它是严格递增的。
 从下标 5 开始的子数组为 [2, 3, 4]，它也是严格递增的。
 两个子数组是相邻的，因此结果为 true。


 示例 2：


 输入：nums = [1,2,3,4,4,4,4,5,6,7], k = 5


 输出：false



 提示：


 2 <= nums.length <= 100
 1 <= 2 * k <= nums.length
 -1000 <= nums[i] <= 1000


 Related Topics 数组 👍 11 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)

	// 需要至少 2*k 个元素才能有两个相邻的长度为 k 的子数组
	if n < 2*k {
		return false
	}

	// 检查每一对相邻的长度为 k 的子数组
	for i := 0; i <= n-2*k; i++ {
		// 检查从 i 开始的第一个子数组是否严格递增
		if isArrayIncrease(nums, i, i+k-1) &&
			// 检查从 i+k 开始的第二个子数组是否严格递增
			isArrayIncrease(nums, i+k, i+2*k-1) {
			return true
		}
	}

	return false
}

func isArrayIncrease(nums []int, start, end int) bool {
	for i := start + 1; i <= end; i++ { // 修改为 i <= end
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func TestAdjacentIncreasingSubarraysDetectionI(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3, true},
		{[]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, 5, false},
		{[]int{1, 2, 3, 4, 5, 6}, 3, true},
		{[]int{1, 2, 1, 2, 3, 4}, 2, true},
		{[]int{1, 2, 3}, 2, false},
		{[]int{-15, 19}, 1, true},
		{[]int{-3, -19, -8, -16}, 2, false},
	}

	for i, test := range tests {
		result := hasIncreasingSubarrays(test.nums, test.k)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, test.expected, result)
		}
	}
}
