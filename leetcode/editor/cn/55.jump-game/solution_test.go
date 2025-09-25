package leetcode

// 55.跳跃游戏

import (
	"testing"
)

/**
给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。



 示例 1：


输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。


 示例 2：


输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。




 提示：


 1 <= nums.length <= 10⁴
 0 <= nums[i] <= 10⁵


 Related Topics 贪心 数组 动态规划 👍 3135 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)

/*
正确解题思路
这是一个典型的贪心算法问题，有两种主要解法：
方法一：贪心算法（推荐）
维护一个变量 maxReach 表示当前能到达的最远位置，遍历数组不断更新这个值。
方法二：逆向思维
从终点开始，看是否能回推到起点。

*/

// leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	// 贪心算法：维护能到达的最远位置
	maxReach := 0 // 能到达的最远下标

	for i := 0; i < len(nums); i++ {
		// 如果当前位置超过了能到达的最远位置，说明无法到达
		if i > maxReach {
			return false
		}

		// 更新能到达的最远位置
		reach := i + nums[i]
		if reach > maxReach {
			maxReach = reach
		}

		// 如果已经能到达或超过最后一个下标，直接返回true
		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return maxReach >= len(nums)-1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestJumpGame(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},  // 示例1
		{[]int{3, 2, 1, 0, 4}, false}, // 示例2
		{[]int{0}, true},              // 只有一个元素
		{[]int{1, 0}, true},           // 两元素可到达
		{[]int{0, 1}, false},          // 第一步就无法跳跃
		{[]int{1, 1, 1, 0}, true},     // 最后一个是0但可以到达
		{[]int{1, 1, 0, 1}, false},    // 无法越过中间的0
		{[]int{2, 0, 0}, true},        // 可以跳过中间的0
	}

	for i, test := range tests {
		result := canJump(test.nums)
		if result != test.expected {
			t.Errorf("Test %d failed: input=%v, expected=%v, got=%v",
				i+1, test.nums, test.expected, result)
		}
	}
}
