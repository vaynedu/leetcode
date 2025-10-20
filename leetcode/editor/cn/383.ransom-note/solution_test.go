package leetcode

// 383.赎金信

import (
	"testing"
)

/**
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

 如果可以，返回 true ；否则返回 false 。

 magazine 中的每个字符只能在 ransomNote 中使用一次。



 示例 1：


输入：ransomNote = "a", magazine = "b"
输出：false


 示例 2：


输入：ransomNote = "aa", magazine = "ab"
输出：false


 示例 3：


输入：ransomNote = "aa", magazine = "aab"
输出：true




 提示：


 1 <= ransomNote.length, magazine.length <= 10⁵
 ransomNote 和 magazine 由小写英文字母组成


 Related Topics 哈希表 字符串 计数 👍 991 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}

	charCount := make(map[rune]int)
	for _, ch := range magazine {
		charCount[ch]++
	}

	for _, ch := range ransomNote {
		if charCount[ch] <= 0 {
			return false
		}

		charCount[ch]--
	}

	return true

}

//leetcode submit region end(Prohibit modification and deletion)

func TestRansomNote(t *testing.T) {
	// 测试用例
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
		{"aab", "baa", true},
		{"", "abc", true},
		{"abc", "", false},
	}

	for i, test := range tests {
		result := canConstruct(test.ransomNote, test.magazine)
		if result != test.expected {
			t.Errorf("Test %d failed. Input: ransomNote=%s, magazine=%s. Expected %v, got %v",
				i, test.ransomNote, test.magazine, test.expected, result)
		}
	}
}
