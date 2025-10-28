package leetcode

// 3354.使数组元素等于零

import (
	"testing"
)

/**
给你一个整数数组 nums 。

 开始时，选择一个满足 nums[curr] == 0 的起始位置 curr ，并选择一个移动 方向 ：向左或者向右。

 此后，你需要重复下面的过程：


 如果 curr 超过范围 [0, n - 1] ，过程结束。
 如果 nums[curr] == 0 ，沿当前方向继续移动：如果向右移，则 递增 curr ；如果向左移，则 递减 curr 。
 如果 nums[curr] > 0:

 将 nums[curr] 减 1 。
 反转 移动方向（向左变向右，反之亦然）。
 沿新方向移动一步。



 如果在结束整个过程后，nums 中的所有元素都变为 0 ，则认为选出的初始位置和移动方向 有效 。

 返回可能的有效选择方案数目。



 示例 1：


 输入：nums = [1,0,2,0,3]


 输出：2

 解释：

 可能的有效选择方案如下：


 选择 curr = 3 并向左移动。



 [1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,1,0,3] -> [1,0,1,0,3] -> [1,0,1,0,2] -> [1,0
,1,0,2] -> [1,0,0,0,2] -> [1,0,0,0,2] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,
1] -> [1,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] ->
 [0,0,0,0,0].


 选择 curr = 3 并向右移动。

 [1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,2,0,2] -> [1,0,2,0,2] -> [1,0,1,0,2] -> [1,0
,1,0,2] -> [1,0,1,0,1] -> [1,0,1,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,
0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [0,0,0,0,0].




 示例 2：


 输入：nums = [2,3,4,0,4,1,0]


 输出：0

 解释：

 不存在有效的选择方案。



 提示：


 1 <= nums.length <= 100
 0 <= nums[i] <= 100
 至少存在一个元素 i 满足 nums[i] == 0 。


 Related Topics 数组 前缀和 模拟 👍 29 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func countValidSelections(nums []int) int {
	n := len(nums)
	count := 0

	// 遍历每个可能的起始位置
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			// 尝试向左移动
			if canReachZero(append([]int{}, nums...), i, -1) {
				count++
			}
			// 尝试向右移动
			if canReachZero(append([]int{}, nums...), i, 1) {
				count++
			}
		}
	}

	return count
}

// 模拟移动过程
func canReachZero(nums []int, start, direction int) bool {
	n := len(nums)
	curr := start

	for curr >= 0 && curr < n {
		if nums[curr] == 0 {
			// 继续当前方向移动
			curr += direction
		} else {
			// nums[curr] > 0
			nums[curr]--
			// 反转方向
			direction = -direction
			// 沿新方向移动一步
			curr += direction
		}
	}

	// 检查是否所有元素都为0
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMakeArrayElementsEqualToZero(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 0, 2, 0, 3}, 2},
		{[]int{2, 3, 4, 0, 4, 1, 0}, 0},
		{[]int{0, 0, 0, 0}, 8}, // 每个0位置可以向左或向右
		{[]int{1, 0, 1}, 2},
	}

	for i, test := range tests {
		result := countValidSelections(test.nums)
		if result != test.expected {
			t.Errorf("Test %d failed: expected %d, got %d. Input: %v",
				i+1, test.expected, result, test.nums)
		}
	}
}
