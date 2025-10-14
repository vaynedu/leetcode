package leetcode

// 135.分发糖果

import (
	"testing"
)

/**
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

 你需要按照以下要求，给这些孩子分发糖果：


 每个孩子至少分配到 1 个糖果。
 相邻两个孩子中，评分更高的那个会获得更多的糖果。


 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。



 示例 1：


输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。


 示例 2：


输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。



 提示：


 n == ratings.length
 1 <= n <= 2 * 10⁴
 0 <= ratings[i] <= 2 * 10⁴


 Related Topics 贪心 数组 👍 1717 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

/*
解题思路
这是一个经典的贪心算法问题。要满足以下条件：
每个孩子至少分配到 1 个糖果
相邻两个孩子中，评分更高的孩子会获得更多糖果
解决方案是使用两次遍历的贪心策略：
第一次从左到右遍历，确保右边评分高的孩子比左边的孩子获得更多糖果
第二次从右到左遍历，确保左边评分高的孩子比右边的孩子获得更多糖果
最后取两次遍历结果的最大值作为每个孩子的糖果数
*/
func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	candies := make([]int, n)
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	// 从左到右遍历，确保右边评分高的孩子比左边的孩子获得更多糖果
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	// 从右到左遍历，确保左边评分高的孩子比右边的孩子获得更多糖果
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	total := 0
	for _, v := range candies {
		total += v
	}

	return total

}

//leetcode submit region end(Prohibit modification and deletion)

// 测试函数实现
func TestCandy(t *testing.T) {
	tests := []struct {
		ratings  []int
		expected int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{1, 3, 2, 2, 1}, 7},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{5, 4, 3, 2, 1}, 15},
		{[]int{1}, 1},
		{[]int{1, 1, 1, 1}, 4},
	}

	for i, test := range tests {
		result := candy(test.ratings)
		if result != test.expected {
			t.Errorf("Test case %d failed: expected %d, got %d", i+1, test.expected, result)
		}
	}
}
