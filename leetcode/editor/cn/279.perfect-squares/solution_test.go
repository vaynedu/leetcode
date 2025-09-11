package leetcode

// 279.完全平方数

import (
	"fmt"
	"testing"
)

/**
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。



 示例 1：


输入：n = 12
输出：3
解释：12 = 4 + 4 + 4

 示例 2：


输入：n = 13
输出：2
解释：13 = 4 + 9



 提示：


 1 <= n <= 10⁴


 👍 2222 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	// dp[i] 表示和为 i 的完全平方数的最少数量
	dp := make([]int, n+1)

	// 初始化：最坏情况是i个1相加
	// 例如：dp[3] = 3 表示 1+1+1
	for i := 1; i <= n; i++ {
		dp[i] = i // 初始设为最坏情况

		// 尝试减去所有可能的完全平方数 j²
		// j从1开始，直到j² > i为止
		for j := 1; j*j <= i; j++ {
			square := j * j // 当前完全平方数

			// 状态转移方程：
			// dp[i] = min(dp[i], dp[i-j²] + 1)
			// 其中 dp[i-j²] + 1 表示：
			//   使用了1个j²，剩下的(i-j²)需要dp[i-j²]个完全平方数
			if dp[i-square]+1 < dp[i] {
				dp[i] = dp[i-square] + 1
			}
		}
	}

	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPerfectSquares(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
