package leetcode

// 611.有效三角形的个数

import (
	"sort"
	"testing"
)

/**
给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。



 示例 1:


输入: nums = [2,2,3,4]
输出: 3
解释:有效的组合是:
2,3,4 (使用第一个 2)
2,3,4 (使用第二个 2)
2,2,3


 示例 2:


输入: nums = [4,2,3,4]
输出: 4



 提示:


 1 <= nums.length <= 1000
 0 <= nums[i] <= 1000


 Related Topics 贪心 数组 双指针 二分查找 排序 👍 667 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	ans := 0

	// 固定最大边，用双指针找其他两条边
	for i := 2; i < len(nums); i++ {
		left := 0
		right := i - 1

		// 使用双指针技术
		for left < right {
			// 如果较小的两个边之和大于最大边，说明可以构成三角形
			if nums[left]+nums[right] > nums[i] {
				// 由于数组已排序，left到right之间的所有元素与right和i都能构成三角形
				ans += right - left
				right--
			} else {
				// 否则增大较小的边
				left++
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidTriangleNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 4}, 3},
		{[]int{4, 2, 3, 4}, 4},
		{[]int{1, 1, 1}, 1},
		{[]int{1, 2, 3}, 0},
		{[]int{0, 0, 0}, 0},
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
	}

	for _, test := range tests {
		result := triangleNumber(test.input)
		if result != test.expected {
			t.Errorf("Input: %v, Expected: %d, Got: %d", test.input, test.expected, result)
		}
	}
}
