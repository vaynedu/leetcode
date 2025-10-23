package leetcode

// 205.同构字符串

import (
	"testing"
)

/**
给定两个字符串 s 和 t ，判断它们是否是同构的。

 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。



 示例 1:


输入：s = "egg", t = "add"
输出：true


 示例 2：


输入：s = "foo", t = "bar"
输出：false

 示例 3：


输入：s = "paper", t = "title"
输出：true



 提示：





 1 <= s.length <= 5 * 10⁴
 t.length == s.length
 s 和 t 由任意有效的 ASCII 字符组成


 Related Topics 哈希表 字符串 👍 798 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/*

解决方案：
使用两个哈希表分别记录s到t和t到s的字符映射关系
遍历字符串，检查映射关系是否一致

*/
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if val, ok := sToT[s[i]]; ok {
			if val != t[i] {
				return false
			}
		} else {
			sToT[s[i]] = t[i]
		}

		// 检查t到s的映射
		if val, ok := tToS[t[i]]; ok {
			// 如果已经存在映射，但与当前字符不一致，则不是同构字符串
			if val != s[i] {
				return false
			}
		} else {
			// 建立新的映射关系
			tToS[t[i]] = s[i]
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIsomorphicStrings(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"ab", "aa", false},
		{"", "", true},
		{"a", "a", true},
	}

	for _, test := range tests {
		result := isIsomorphic(test.s, test.t)
		if result != test.expected {
			t.Errorf("isIsomorphic(%s, %s) = %v; expected %v", test.s, test.t, result, test.expected)
		}
	}
}
