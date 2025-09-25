package leetcode

// 120.三角形最小路径和

import (
	"testing"
)

/**
给定一个三角形 triangle ，找出自顶向下的最小路径和。

 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位
于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。



 示例 1：


输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。


 示例 2：


输入：triangle = [[-10]]
输出：-10




 提示：


 1 <= triangle.length <= 200
 triangle[0].length == 1
 triangle[i].length == triangle[i - 1].length + 1
 -10⁴ <= triangle[i][j] <= 10⁴




 进阶：


 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？


 Related Topics 数组 动态规划 👍 1486 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

/*
这是一个经典的动态规划问题。可以采用自顶向下或自底向上的方法：
自底向上动态规划（推荐）：
从倒数第二层开始向上计算
每个位置的最小路径和 = 当前值 + min(下一层相邻两个位置的最小路径和)
最终 triangle[0][0] 就是答案
空间优化：
只需要一个长度为 n 的数组存储每层的结果
从底向上更新数

1.状态定义：dp[j] 表示从当前层第 j 个位置到底层的最小路径和
2.状态转移方程：dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
3.初始化：最底层的 dp 值就是 triangle 最后一行的值
4.计算顺序：从倒数第二行向上计算到第 0 行
5.空间优化：复用一个一维数组，避免使用二维数组

*/

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			// 当前位置的最小路径和 = 当前值 + 下一层相邻两位置的较小值
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}
	return dp[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTriangle(t *testing.T) {
	tests := []struct {
		triangle [][]int
		expected int
	}{
		{
			triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			expected: 11,
		},
		{
			triangle: [][]int{{-10}},
			expected: -10,
		},
		{
			triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}},
			expected: -1,
		},
	}

	for i, test := range tests {
		result := minimumTotal(test.triangle)
		if result != test.expected {
			t.Errorf("Test %d failed: expected %d, got %d", i+1, test.expected, result)
		}
	}
}
