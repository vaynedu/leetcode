package leetcode

// 53.最大子数组和

import (
	"fmt"
	"testing"
)

/**
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

 子数组 是数组中的一个连续部分。



 示例 1：


输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。


 示例 2：


输入：nums = [1]
输出：1


 示例 3：


输入：nums = [5,4,-1,7,8]
输出：23




 提示：


 1 <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴




 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

 Related Topics 数组 分治 动态规划 👍 6996 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// nums = [-2,1,-3,4,-1,2,1,-5,4]
	// 连续子序列
	sum := nums[0]
	curSum := nums[0]
	for i := 1; i < n; i++ {
		curSum = max(curSum+nums[i], nums[i])
		sum = max(sum, curSum)
	}

	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumSubarray(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
	tests := []struct {
		input  []int
		expect int
	}{
		{
			input:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expect: 6,
		},
		{
			input:  []int{1},
			expect: 1,
		},
		{
			input:  []int{5, 4, -1, 7, 8},
			expect: 23,
		},
	}
	for _, test := range tests {
		actual := maxSubArray(test.input)
		if actual != test.expect {
			t.Errorf("input: %v, expect: %v, actual: %v", test.input, test.expect, actual)
		}
	}
}
