package leetcode

// 15.三数之和

import (
	"sort"
	"testing"
)

/**
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

 注意：答案中不可以包含重复的三元组。





 示例 1：


输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。


 示例 2：


输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。


 示例 3：


输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。




 提示：


 3 <= nums.length <= 3000
 -10⁵ <= nums[i] <= 10⁵


 Related Topics 数组 双指针 排序 👍 7705 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	// 排序数组
	sort.Ints(nums)
	n := len(nums)

	// 边界条件：如果数组长度小于3或全为正数/负数，则无解
	if n < 3 || nums[0] > 0 || nums[n-1] < 0 {
		return [][]int{}
	}

	res := make([][]int, 0)

	// 遍历数组，固定第一个数
	for k := 0; k < n-2; k++ {
		// 如果当前数字大于0，则三数之和不可能为0
		if nums[k] > 0 {
			break
		}

		// 跳过重复元素，避免重复三元组
		if k > 0 && nums[k-1] == nums[k] {
			continue
		}

		// 双指针寻找另外两个数
		i := k + 1
		j := n - 1

		for i < j {
			sum := nums[k] + nums[i] + nums[j]

			if sum == 0 {
				// 找到一个三元组
				res = append(res, []int{nums[k], nums[i], nums[j]})

				// 跳过重复元素
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}

				// 移动双指针
				i++
				j--
			} else if sum < 0 {
				// 和小于0，需要增大，移动左指针
				i++
			} else {
				// 和大于0，需要减小，移动右指针
				j--
			}
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{[]int{}, [][]int{}},
		{[]int{1, 2, 3}, [][]int{}},
	}

	for _, test := range tests {
		result := threeSum(test.input)
		if !equalTriplets(result, test.expected) {
			t.Errorf("Input: %v, Expected: %v, Got: %v", test.input, test.expected, result)
		}
	}
}

// 辅助函数：比较两个三元组切片是否相等
func equalTriplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// 由于结果可能顺序不同，需要排序后比较
	sort.Slice(a, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if a[i][k] != a[j][k] {
				return a[i][k] < a[j][k]
			}
		}
		return false
	})

	sort.Slice(b, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if b[i][k] != b[j][k] {
				return b[i][k] < b[j][k]
			}
		}
		return false
	})

	for i := range a {
		for j := 0; j < 3; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
