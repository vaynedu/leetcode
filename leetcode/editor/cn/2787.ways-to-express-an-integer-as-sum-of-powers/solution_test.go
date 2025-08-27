package leetcode

// 2787.将一个数字表示成幂的和的方案数

import (
	"testing"
)

/**
给你两个 正 整数 n 和 x 。

 请你返回将 n 表示成一些 互不相同 正整数的 x 次幂之和的方案数。换句话说，你需要返回互不相同整数 [n1, n2, ..., nk] 的集合数目，满足
n = n1ˣ + n2ˣ + ... + nkˣ 。

 由于答案可能非常大，请你将它对 10⁹ + 7 取余后返回。

 比方说，n = 160 且 x = 3 ，一个表示 n 的方法是 n = 2³ + 3³ + 5³ 。



 示例 1：

 输入：n = 10, x = 2
输出：1
解释：我们可以将 n 表示为：n = 3² + 1² = 10 。
这是唯一将 10 表达成不同整数 2 次方之和的方案。


 示例 2：

 输入：n = 4, x = 1
输出：2
解释：我们可以将 n 按以下方案表示：
- n = 4¹ = 4 。
- n = 3¹ + 1¹ = 4 。




 提示：


 1 <= n <= 300
 1 <= x <= 5


 Related Topics 动态规划 👍 49 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfWays(n int, x int) int {
	const mod = 1e9 + 7
	// 计算所有可能用到的x次幂值
	powers := make([]int, 0)
	for i := 1; ; i++ {
		power := pow(i, x)
		if power > n {
			break
		}
		powers = append(powers, power)
	}
	dp := make([]int, n+1)
	dp[0] = 1
	// 对每个x次幂值进行遍历
	for _, power := range powers {
		// 从后往前更新，避免重复使用同一个幂值
		for i := n; i >= power; i-- {
			dp[i] = (dp[i] + dp[i-power]) % mod
		}
	}
	return dp[n]
}

// 计算 base 的 exp 次幂
func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestWaysToExpressAnIntegerAsSumOfPowers(t *testing.T) {
	tests := []struct {
		n        int
		x        int
		expected int
	}{
		{10, 2, 1},
		{4, 1, 2},
		{160, 3, 1}, // 2³ + 3³ + 5³ = 8 + 27 + 125 = 160
	}

	for _, test := range tests {
		result := numberOfWays(test.n, test.x)
		if result != test.expected {
			t.Errorf("numberOfWays(%d, %d) = %d; expected %d", test.n, test.x, result, test.expected)
		} else {
			t.Logf("numberOfWays(%d, %d) = %d; passed", test.n, test.x, result)
		}
	}
}
