package leetcode

// 125.验证回文串

import (
	"testing"
)

/**
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

 字母和数字都属于字母数字字符。

 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。



 示例 1：


输入: s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。


 示例 2：


输入：s = "race a car"
输出：false
解释："raceacar" 不是回文串。


 示例 3：


输入：s = " "
输出：true
解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
由于空字符串正着反着读都一样，所以是回文串。




 提示：


 1 <= s.length <= 2 * 10⁵
 s 仅由可打印的 ASCII 字符组成


 Related Topics 双指针 字符串 👍 843 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)

func isPalindrome(s string) bool {

	// 移除字母和数字
	sNew := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') ||
			(s[i] >= 'A' && s[i] <= 'Z') {
			// 转换大写字母为小写
			if s[i] >= 'A' && s[i] <= 'Z' {
				sNew = append(sNew, s[i]+'a'-'A')
			} else {
				sNew = append(sNew, s[i])
			}
		}
	}

	// 判断是否是回文
	for i := 0; i < len(sNew)/2; i++ {
		if sNew[i] != sNew[len(sNew)-i-1] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"", true},
		{"a", true},
		{"Madam", true},
		{"No 'x' in Nixon", true},
		{"0P", false},
		{"1a2", false},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("Input: '%s', Expected: %t, Got: %t", test.input, test.expected, result)
		}
	}
}
