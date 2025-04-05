package leetcode

// 找出所有子集的异或总和再求和

import (
	"fmt"
	"testing"
)

/**
一个数组的 异或总和 定义为数组中所有元素按位 XOR 的结果；如果数组为 空 ，则异或总和为 0 。


 例如，数组 [2,5,6] 的 异或总和 为 2 XOR 5 XOR 6 = 1 。


 给你一个数组 nums ，请你求出 nums 中每个 子集 的 异或总和 ，计算并返回这些值相加之 和 。

 注意：在本题中，元素 相同 的不同子集应 多次 计数。

 数组 a 是数组 b 的一个 子集 的前提条件是：从 b 删除几个（也可能不删除）元素能够得到 a 。



 示例 1：

 输入：nums = [1,3]
输出：6
解释：[1,3] 共有 4 个子集：
- 空子集的异或总和是 0 。
- [1] 的异或总和为 1 。
- [3] 的异或总和为 3 。
- [1,3] 的异或总和为 1 XOR 3 = 2 。
0 + 1 + 3 + 2 = 6


 示例 2：

 输入：nums = [5,1,6]
输出：28
解释：[5,1,6] 共有 8 个子集：
- 空子集的异或总和是 0 。
- [5] 的异或总和为 5 。
- [1] 的异或总和为 1 。
- [6] 的异或总和为 6 。
- [5,1] 的异或总和为 5 XOR 1 = 4 。
- [5,6] 的异或总和为 5 XOR 6 = 3 。
- [1,6] 的异或总和为 1 XOR 6 = 7 。
- [5,1,6] 的异或总和为 5 XOR 1 XOR 6 = 2 。
0 + 5 + 1 + 6 + 4 + 3 + 7 + 2 = 28


 示例 3：

 输入：nums = [3,4,5,6,7,8]
输出：480
解释：每个子集的全部异或总和值之和为 480 。




 提示：


 1 <= nums.length <= 12
 1 <= nums[i] <= 20


 Related Topics 位运算 数组 数学 回溯 组合数学 枚举 👍 179 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)

func subsetXORSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 1.计算当前数组的所有子集
	// 2.然后数组求和

	totalSum := 0
	var backTrack func(int, int)
	backTrack = func(index, currentXOR int) {
		if index == len(nums) {
			totalSum += currentXOR
			return
		}

		// 不选择当前数字
		backTrack(index+1, currentXOR)

		// 选择当前的数字
		backTrack(index+1, currentXOR^nums[index])

	}

	backTrack(0, 0)
	return totalSum

}

//leetcode submit region end(Prohibit modification and deletion)

// subsetXORSumTwo 计算数组所有子集的异或总和再求和
func subsetXORSumTwo(nums []int) int {
	// 输入验证
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// 初始化结果
	totalSum := 0

	// 使用位运算生成所有子集
	// 使用位运算中的 mask 来代表所有子集是一种高效且直观的方法。通过将每个子集用一个整数 mask
	// 我们可以利用二进制位来决定数组中的每个元素是否包含在当前子集中。
	n := len(nums)
	for mask := 0; mask < (1 << n); mask++ {
		xorSum := 0
		for i := 0; i < n; i++ {
			// 如果第 i 位被选中，则加入异或计算
			if mask&(1<<i) != 0 {
				xorSum ^= nums[i]
			}
		}
		totalSum += xorSum
	}

	return totalSum
}

func TestSumOfAllSubsetXorTotals(t *testing.T) {
	// 测试用例
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 3}, 6},
		{[]int{5, 1, 6}, 28},
		{[]int{3, 4, 5, 6, 7, 8}, 480},
		{[]int{}, 0},    // 空数组测试
		{[]int{10}, 10}, // 单元素数组测试
	}

	// 执行测试
	for _, tc := range testCases {
		result := subsetXORSum(tc.input)
		if result != tc.expected {
			t.Errorf("对于输入 %v, 期望输出 %d, 实际输出 %d", tc.input, tc.expected, result)
		} else {
			fmt.Printf("测试通过: 输入 %v, 输出 %d\n", tc.input, result)
		}
	}
}
