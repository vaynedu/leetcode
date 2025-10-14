package leetcode

import "testing"

// 45.跳跃游戏 II

/**
给定一个长度为 n 的 0 索引整数数组 nums。初始位置在下标 0。

 每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在索引 i 处，你可以跳转到任意 (i + j) 处：


 0 <= j <= nums[i] 且
 i + j < n


 返回到达 n - 1 的最小跳跃次数。测试用例保证可以到达 n - 1。



 示例 1:


输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。


 示例 2:


输入: nums = [2,3,0,1,4]
输出: 2




 提示:


 1 <= nums.length <= 10⁴
 0 <= nums[i] <= 1000
 题目保证可以到达 n - 1


 Related Topics 贪心 数组 动态规划 👍 2903 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	jumps := 0      // 跳跃次数
	currentEnd := 0 // 当前跳跃能到达的最远位置
	farthest := 0   // 下一步能到达的最远位置

	// 注意：不需要考虑最后一个位置，因为我们只需要到达它，不需要从它再跳出去
	for i := 0; i < len(nums)-1; i++ {
		// 更新在当前跳跃范围内，下一步能到达的最远位置
		farthest = max(farthest, i+nums[i])

		// 如果已经到达当前跳跃的边界
		if i == currentEnd {
			jumps++
			currentEnd = farthest

			// 如果已经能到达最后位置，提前结束
			if currentEnd >= len(nums)-1 {
				break
			}
		}
	}

	return jumps
}

// leetcode submit region end(Prohibit modification and deletion)

// 测试函数实现
func TestJumpGameIi(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{1, 1, 1, 1}, 3},
		{[]int{1}, 0},
		{[]int{2, 1}, 1},
	}

	for i, test := range tests {
		result := jump(test.nums)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %d, got %d", i+1, test.expected, result)
		}
	}
}
