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

	// Step 1: Sort the array
	sort.Ints(nums)

	// Step 2: Initialize dp and prev arrays
	n := len(nums)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}

	// Step 3: Dynamic Programming
	maxIndex := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[maxIndex] {
			maxIndex = i
		}
	}

	// Step 4: Reconstruct the largest divisible subset
	result := []int{}
	for maxIndex != -1 {
		result = append(result, nums[maxIndex])
		maxIndex = prev[maxIndex]
	}

	// Reverse the result to get the correct order
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
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
