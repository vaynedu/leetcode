package leetcode

// 3375.使数组的值全部为 K 的最少操作次数

import (
	"fmt"
	"testing"
)

/**
给你一个整数数组 nums 和一个整数 k 。

 如果一个数组中所有 严格大于 h 的整数值都 相等 ，那么我们称整数 h 是 合法的 。

 比方说，如果 nums = [10, 8, 10, 8] ，那么 h = 9 是一个 合法 整数，因为所有满足 nums[i] > 9 的数都等于 10 ，但
是 5 不是 合法 整数。

 你可以对 nums 执行以下操作：


 选择一个整数 h ，它对于 当前 nums 中的值是合法的。
 对于每个下标 i ，如果它满足 nums[i] > h ，那么将 nums[i] 变为 h 。


 你的目标是将 nums 中的所有元素都变为 k ，请你返回 最少 操作次数。如果无法将所有元素都变 k ，那么返回 -1 。



 示例 1：


 输入：nums = [5,2,5,4,5], k = 2


 输出：2

 解释：

 依次选择合法整数 4 和 2 ，将数组全部变为 2 。

 示例 2：


 输入：nums = [2,1,2], k = 2


 输出：-1

 解释：

 没法将所有值变为 2 。

 示例 3：


 输入：nums = [9,7,5,3], k = 1


 输出：4

 解释：

 依次选择合法整数 7 ，5 ，3 和 1 ，将数组全部变为 1 。



 提示：


 1 <= nums.length <= 100
 1 <= nums[i] <= 100
 1 <= k <= 100


 Related Topics 数组 哈希表 👍 9 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums []int, k int) int {
	minValue := 100
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] < minValue {
			minValue = nums[i]
		}
	}
	if k > minValue {
		return -1
	}
	// 说明当前数组一定有解法，找当前的重复数字就行
	duplicateNumMap := make(map[int]struct{})
	for i := 0; i < n; i++ {
		if nums[i] != k {
			duplicateNumMap[nums[i]] = struct{}{}
		}
	}
	return len(duplicateNumMap)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumOperationsToMakeArrayValuesEqualToK(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
	tests := []struct {
		input  []int
		k      int
		expect int
	}{
		{
			input:  []int{5, 2, 5, 4, 5},
			k:      2,
			expect: 2,
		},
		{
			input:  []int{2, 1, 2},
			k:      2,
			expect: -1,
		},
		{
			input:  []int{9, 7, 5, 3},
			k:      1,
			expect: 4,
		},
	}
	for _, test := range tests {
		actual := minOperations(test.input, test.k)
		if actual != test.expect {
			t.Errorf("input: %v, k: %v, expect: %v, actual: %v", test.input, test.k, test.expect, actual)
		}
	}
}
