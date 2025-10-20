package leetcode

// 3.无重复字符的最长子串

import (
	"testing"
)

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。



 示例 1:


输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。


 示例 2:


输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。


 示例 3:


输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。




 提示：


 0 <= s.length <= 5 * 10⁴
 s 由英文字母、数字、符号和空格组成


 Related Topics 哈希表 字符串 滑动窗口 👍 11073 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	// 使用滑动窗口 + 哈希表优化
	// map记录字符最后出现的位置
	charIndex := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < len(s); right++ {
		// 如果字符已存在且在当前窗口内，则移动左指针
		if index, exist := charIndex[s[right]]; exist && index >= left {
			left = index + 1
		}
		// 更新字符最后出现位置
		charIndex[s[right]] = right
		// 更新最大长度
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

// leetcode submit region end(Prohibit modification and deletion)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdef", 6},
		{"aab", 2},
		{"dvdf", 3},
	}

	for i, test := range tests {
		result := lengthOfLongestSubstring(test.input)
		if result != test.expected {
			t.Errorf("Test case %d failed: input=%s, expected=%d, got=%d",
				i+1, test.input, test.expected, result)
		} else {
			t.Logf("Test case %d passed: input=%s, result=%d",
				i+1, test.input, result)
		}
	}
}
