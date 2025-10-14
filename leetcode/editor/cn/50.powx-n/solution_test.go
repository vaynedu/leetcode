package leetcode

// 50.Pow(x, n)

import (
	"testing"
)

/**
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xⁿ ）。



 示例 1：


输入：x = 2.00000, n = 10
输出：1024.00000


 示例 2：


输入：x = 2.10000, n = 3
输出：9.26100


 示例 3：


输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25




 提示：


 -100.0 < x < 100.0
 -231 <= n <= 231-1
 n 是一个整数
 要么 x 不为零，要么 n > 0 。
 -104 <= xⁿ <= 104


 Related Topics 递归 数学 👍 1499 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	// 处理负指数的情况
	if n < 0 {
		x = 1 / x
		n = -n
	}

	// 使用快速幂算法
	return fastPow(x, n)
}

// 快速幂算法实现
func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	// 递归计算 x^(n/2)
	half := fastPow(x, n/2)

	// 根据n的奇偶性计算结果
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

// leetcode submit region end(Prohibit modification and deletion)

// 测试函数实现
func TestPowxN(t *testing.T) {
	tests := []struct {
		x        float64
		n        int
		expected float64
	}{
		{2.00000, 10, 1024.00000},
		{2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
		{3.00000, 0, 1.00000},
		{0.00000, 1, 0.00000},
		{1.00000, 2147483647, 1.00000},
		{-2.00000, 2, 4.00000},
		{-2.00000, 3, -8.00000},
	}

	for i, test := range tests {
		result := myPow(test.x, test.n)
		// 允许一定的浮点数误差
		if abs(result-test.expected) > 1e-5 {
			t.Errorf("Test case %d failed: expected %f, got %f", i+1, test.expected, result)
		}
	}
}

// 辅助函数：计算绝对值
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
