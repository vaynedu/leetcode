package leetcode

// 3396.使数组元素互不相同所需的最少操作次数

import (
	"fmt"
	"testing"
)

/**
给你一个整数数组 nums，你需要确保数组中的元素 互不相同 。为此，你可以执行以下操作任意次：


 从数组的开头移除 3 个元素。如果数组中元素少于 3 个，则移除所有剩余元素。


 注意：空数组也视作为数组元素互不相同。返回使数组元素互不相同所需的 最少操作次数 。






 示例 1：


 输入： nums = [1,2,3,4,2,3,3,5,7]


 输出： 2

 解释：


 第一次操作：移除前 3 个元素，数组变为 [4, 2, 3, 3, 5, 7]。
 第二次操作：再次移除前 3 个元素，数组变为 [3, 5, 7]，此时数组中的元素互不相同。


 因此，答案是 2。

 示例 2：


 输入： nums = [4,5,6,4,4]


 输出： 2

 解释：


 第一次操作：移除前 3 个元素，数组变为 [4, 4]。
 第二次操作：移除所有剩余元素，数组变为空。


 因此，答案是 2。

 示例 3：


 输入： nums = [6,7,8,9]


 输出： 0

 解释：

 数组中的元素已经互不相同，因此不需要进行任何操作，答案是 0。



 提示：


 1 <= nums.length <= 100
 1 <= nums[i] <= 100


 Related Topics 数组 哈希表 👍 12 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)

// minimumOperations 计算使数组元素互不相同所需的最少操作次数
func minimumOperations(nums []int) int {
	count := 0
	for len(nums) > 0 {
		if !hasDuplicates(nums) {
			break
		}
		if len(nums) < 3 {
			nums = nums[:0]
		} else {
			nums = nums[3:]
		}
		count++
	}
	return count
}

// hasDuplicates 检查数组中是否存在重复元素
func hasDuplicates(nums []int) bool {
	numSet := make(map[int]bool)
	for _, num := range nums {
		if numSet[num] {
			return true
		}
		numSet[num] = true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumNumberOfOperationsToMakeElementsInArrayDistinct(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
	tests := []struct {
		input  []int
		expect int
	}{
		{
			input:  []int{5, 5},
			expect: 1,
		},
		{
			input:  []int{1},
			expect: 0,
		},
		{
			input:  []int{1, 2},
			expect: 0,
		},
		{
			input:  []int{3, 7, 7, 3},
			expect: 1,
		},
		{
			input:  []int{1, 2, 2},
			expect: 1,
		},
		{
			input:  []int{1, 2, 3},
			expect: 0,
		},
		{
			input:  []int{1, 2, 3, 4, 2, 3, 3, 5, 7},
			expect: 2,
		},
		{
			input:  []int{4, 5, 6, 4, 4},
			expect: 2,
		},
		{
			input:  []int{6, 7, 8, 9},
			expect: 0,
		},
		{
			input:  []int{1, 1, 1},
			expect: 1,
		},
		{
			input:  []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			expect: 3,
		},
		{
			input:  []int{8, 10, 7, 1, 5, 1, 8},
			expect: 2,
		},
	}
	for _, test := range tests {
		actual := minimumOperations(test.input)
		if actual != test.expect {
			t.Errorf("input: %v, expect: %v, actual: %v", test.input, test.expect, actual)
		}
	}
}
