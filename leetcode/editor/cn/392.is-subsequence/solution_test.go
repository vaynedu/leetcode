package leetcode

// 392.判断子序列

import (
	"testing"
)

/**
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而
"aec"不是）。

 进阶：

 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？


 致谢：

 特别感谢 @pbrother 添加此问题并且创建所有测试用例。



 示例 1：


输入：s = "abc", t = "ahbgdc"
输出：true


 示例 2：


输入：s = "axc", t = "ahbgdc"
输出：false




 提示：


 0 <= s.length <= 100
 0 <= t.length <= 10^4
 两个字符串都只由小写字符组成。


 Related Topics 双指针 字符串 动态规划 👍 1198 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func isSubsequence(s string, t string) bool {

	count := 0
	start := 0
	for i := 0; i < len(s); i++ {
		for j := start; j < len(t); j++ {
			if s[i] == t[j] {
				start = j + 1
				count++
				break
			}
		}
	}
	if count == len(s) {
		return true
	}
	return false
}

func isSubsequenceV1(s string, t string) bool {
	i := 0 // s的指针
	j := 0 // t的指针

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	// 如果i等于s的长度，说明s中所有字符都被匹配到了
	return i == len(s)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"", "ahbgdc", true},
		{"abc", "", false},
		{"", "", true},
		{"a", "b", false},
		{"a", "a", true},
		{"ace", "abcde", true},
		{"aec", "abcde", false},
	}

	for _, test := range tests {
		result := isSubsequence(test.s, test.t)
		if result != test.expected {
			t.Errorf("Input: s='%s', t='%s', Expected: %t, Got: %t",
				test.s, test.t, test.expected, result)
		}
	}
}
