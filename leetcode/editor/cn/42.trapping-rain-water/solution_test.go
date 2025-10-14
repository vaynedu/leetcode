package leetcode

// 42.接雨水

import (
	"testing"
)

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



 示例 1：




输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。


 示例 2：


输入：height = [4,2,0,3,2,5]
输出：9




 提示：


 n == height.length
 1 <= n <= 2 * 10⁴
 0 <= height[i] <= 10⁵


 Related Topics 栈 数组 双指针 动态规划 单调栈 👍 5920 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/*
接雨水问题可以通过多种方法解决，这里提供双指针法的实现：
使用两个指针分别从数组的两端向中间移动
维护左右两侧的最大高度
对于每个位置，能接的雨水量等于该位置左右两侧最大高度的较小值减去当前位置的高度
通过比较左右两侧最大高度，决定从哪一侧进行处理
*/
func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

// 测试函数实现
func TestTrappingRainWater(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{3, 0, 2, 0, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{1}, 0},
		{[]int{}, 0},
	}

	for i, test := range tests {
		result := trap(test.height)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %d, got %d", i+1, test.expected, result)
		}
	}
}
