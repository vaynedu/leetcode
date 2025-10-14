package leetcode

// 238.除自身以外数组的乘积

import (
	"testing"
)

/**
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。

 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。



 示例 1:


输入: nums = [1,2,3,4]
输出: [24,12,8,6]


 示例 2:


输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]




 提示：


 2 <= nums.length <= 10⁵
 -30 <= nums[i] <= 30
 输入 保证 数组 answer[i] 在 32 位 整数范围内




 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）

 Related Topics 数组 前缀和 👍 2126 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// 第一步：计算每个位置左侧所有元素的乘积
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// 第二步：计算每个位置右侧所有元素的乘积，并与左侧乘积相乘
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] = result[i] * rightProduct
		rightProduct *= nums[i]
	}

	return result
}

// leetcode submit region end(Prohibit modification and deletion)

// 测试函数实现
func TestProductOfArrayExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{2, 3}, []int{3, 2}},
		{[]int{1, 0}, []int{0, 1}},
	}

	for i, test := range tests {
		result := productExceptSelf(test.nums)
		if !equal(result, test.expected) {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, test.expected, result)
		}
	}
}

// 辅助函数：比较两个整数切片是否相等
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
