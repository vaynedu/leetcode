package leetcode

// 121.买卖股票的最佳时机

import (
	"fmt"
	"testing"
)

/**
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。



 示例 1：


输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。


 示例 2：


输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。




 提示：


 1 <= prices.length <= 10⁵
 0 <= prices[i] <= 10⁴


 Related Topics 数组 动态规划 👍 3888 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0] // 记录到目前为止的最低价格
	maxProfits := 0       // 记录到目前为止的最大利润

	for i := 1; i < len(prices); i++ {
		// 如果当前价格更低，更新最低价格
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// 否则计算当前卖出可以获得的利润
			profit := prices[i] - minPrice
			if profit > maxProfits {
				maxProfits = profit
			}
		}
	}

	return maxProfits
}

func maxProfitV1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// dp[i] 表示到第i天能获得的最大利润
	minPrice := prices[0] // 记录到目前为止的最低买入价格
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// 状态转移方程：今天卖出或之前卖出的最大利润
		maxProfit = max(maxProfit, prices[i]-minPrice)
		// 更新到目前为止的最低价格
		minPrice = min(minPrice, prices[i])
	}

	return maxProfit
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
