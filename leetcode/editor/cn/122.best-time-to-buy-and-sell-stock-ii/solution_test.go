package leetcode

// 122.买卖股票的最佳时机 II

import (
	"testing"
)

/**
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。然而，你可以在 同一天 多次买卖该股票，但要确保你持有的股票不超过一股。


 返回 你能获得的 最大 利润 。



 示例 1：


输入：prices = [7,1,5,3,6,4]
输出：7
解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
最大总利润为 4 + 3 = 7 。

 示例 2：


输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
最大总利润为 4 。

 示例 3：


输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0。



 提示：


 1 <= prices.length <= 3 * 10⁴
 0 <= prices[i] <= 10⁴


 Related Topics 贪心 数组 动态规划 👍 2787 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/*

解题思路
这是一个经典的贪心算法问题。核心思想是：只要明天的价格比今天高，就在今天买入明天卖出，这样可以获得所有可能的利润。
有两种理解方式：
贪心策略：每次价格上涨时都进行交易
峰谷法：在每个局部最低点买入，在局部最高点卖出
实际上这两种方法是等价的，因为每次价格上涨都进行交易就相当于在连续上涨的区间内完成了最优交易。
*/
func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBestTimeToBuyAndSellStockIi(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 1, 3}, 3},
		{[]int{1}, 0},
		{[]int{1, 1, 1, 1}, 0},
	}

	for i, test := range tests {
		result := maxProfit(test.prices)
		if result != test.expected {
			t.Errorf("Test %d failed: expected %d, got %d. Input: %v",
				i+1, test.expected, result, test.prices)
		}
	}
}
