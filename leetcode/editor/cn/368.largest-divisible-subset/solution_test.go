package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// 368.最大整除子集

//给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[
//j]) 都应当满足：
//
//
// answer[i] % answer[j] == 0 ，或
// answer[j] % answer[i] == 0
//
//
// 如果存在多个有效解子集，返回其中任何一个均可。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[1,2]
//解释：[1,3] 也会被视为正确答案。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,4,8]
//输出：[1,2,4,8]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 2 * 10⁹
// nums 中的所有整数 互不相同
//
//
// Related Topics 数组 数学 动态规划 排序 👍 610 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	n := len(nums)

	// Step 1 排序
	sort.Ints(nums)

	// Step 2 创建dp数组，dp[i]表示以nums[i]结尾的子集
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	// Step 3 创建pre数组，pre[i]表示以nums[i]结尾的子集的前一个元素
	pre := make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = -1
	}

	// Step 4 遍历nums，计算dp数组
	maxIndex := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				pre[i] = j
			}
			if dp[i] > dp[maxIndex] {
				maxIndex = i
			}
		}
	}

	// Step 5 根据pre数组构建结果集
	res := make([]int, 0)
	for i := maxIndex; i >= 0; i = pre[i] {
		res = append([]int{nums[i]}, res...)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLargestDivisibleSubset(t *testing.T) {
	fmt.Println("come on")

	// Test cases
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2}},
		{[]int{1, 2, 4, 8}, []int{1, 2, 4, 8}},
		{[]int{1, 3, 6, 24}, []int{1, 3, 6, 24}},
		{[]int{5, 9, 18, 36}, []int{5, 18, 36}},
		{[]int{1}, []int{1}},
	}

	for _, tc := range testCases {
		result := largestDivisibleSubset(tc.nums)
		fmt.Printf("For nums = %v, expected %v, got %v\n", tc.nums, tc.expected, result)
	}
}
