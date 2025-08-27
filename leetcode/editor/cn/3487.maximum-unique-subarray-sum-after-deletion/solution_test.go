package leetcode

// 3487.删除后的最大子数组元素和

import (
	"fmt"
	"sort"
	"testing"
)

/**
给你一个整数数组 nums 。

 你可以从数组 nums 中删除任意数量的元素，但不能将其变为 空 数组。执行删除操作后，选出 nums 中满足下述条件的一个子数组：


 子数组中的所有元素 互不相同 。
 最大化 子数组的元素和。


 返回子数组的 最大元素和 。
子数组 是数组的一个连续、
非空 的元素序列。



 示例 1：


 输入：nums = [1,2,3,4,5]


 输出：15

 解释：

 不删除任何元素，选中整个数组得到最大元素和。

 示例 2：


 输入：nums = [1,1,0,1,1]


 输出：1

 解释：

 删除元素 nums[0] == 1、nums[1] == 1、nums[2] == 0 和 nums[3] == 1 。选中整个数组 [1] 得到最大元素和。


 示例 3：


 输入：nums = [1,2,-1,-2,1,0,-1]


 输出：3

 解释：

 删除元素 nums[2] == -1 和 nums[3] == -2 ，从 [1, 2, 1, 0, -1] 中选中子数组 [2, 1] 以获得最大元素和。




 提示：


 1 <= nums.length <= 100
 -100 <= nums[i] <= 100


 Related Topics 贪心 数组 哈希表 👍 10 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	index := 0
	for i := n - 1; i >= 0; i-- {
		if nums[i] <= 0 {
			index = i
			break
		}

		for i > 0 {
			if nums[i] != nums[i-1] {
				break
			}
			i--
		}

		if i >= 0 {
			sum += nums[i]
		}
	}
	if sum == 0 {
		// 遇到了所有元素都小于0，取第一个负数值，即使最小值
		return nums[index]
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumUniqueSubarraySumAfterDeletion(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
	tests := []struct {
		input  []int
		expect int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			expect: 15,
		},
		{
			input:  []int{1, 1, 0, 1, 1},
			expect: 1,
		},
		{
			input:  []int{1, 2, -1, -2, 1, 0, -1},
			expect: 3,
		},
		{
			input:  []int{-17, -15},
			expect: -15,
		},
		{
			input:  []int{-17, -15, 0, 1, 2},
			expect: 3,
		},
		{
			input:  []int{-17, -15, 0},
			expect: 0,
		},
	}
	for _, test := range tests {
		actual := maxSum(test.input)
		if actual != test.expect {
			t.Errorf("input: %v, expect: %v, actual: %v", test.input, test.expect, actual)
		}
	}
}
