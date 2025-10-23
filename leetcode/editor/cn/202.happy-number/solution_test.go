package leetcode

// 202.快乐数

import (
	"testing"
)

/**
编写一个算法来判断一个数 n 是不是快乐数。

 「快乐数」 定义为：


 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
 如果这个过程 结果为 1，那么这个数就是快乐数。


 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。



 示例 1：


输入：n = 19
输出：true
解释：
1² + 9² = 82
8² + 2² = 68
6² + 8² = 100
1² + 0² + 0² = 1


 示例 2：


输入：n = 2
输出：false




 提示：


 1 <= n <= 2³¹ - 1


 Related Topics 哈希表 数学 双指针 👍 1755 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func isHappy(n int) bool {
	// 使用map记录已经出现过的数字，用于检测循环
	seen := make(map[int]bool)

	for n != 1 && !seen[n] {
		seen[n] = true
		n = getSumOfSquares(n)
	}

	return n == 1
}

// 计算一个数各位数字的平方和
func getSumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

func TestHappyNumber(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{19, true},
		{2, false},
		{1, true},
		{7, true},
		{10, true},
		{1111111, true},
	}

	for _, test := range tests {
		result := isHappy(test.input)
		if result != test.expected {
			t.Errorf("isHappy(%d) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
