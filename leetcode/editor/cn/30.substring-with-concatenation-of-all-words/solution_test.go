package leetcode

// 30.串联所有单词的子串

import (
	"testing"
)

/**
给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。

 s 中的 串联子串 是指一个包含 words 中所有字符串以任意顺序排列连接起来的子串。


 例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，
"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。


 返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。



 示例 1：


输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
输出顺序无关紧要。返回 [9,0] 也是可以的。


 示例 2：


输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]
解释：因为 words.length == 4 并且 words[i].length == 4，所以串联子串的长度必须为 16。
s 中没有子串长度为 16 并且等于 words 的任何顺序排列的连接。
所以我们返回一个空数组。


 示例 3：


输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出：[6,9,12]
解释：因为 words.length == 3 并且 words[i].length == 3，所以串联子串的长度必须为 9。
子串 "foobarthe" 开始位置是 6。它是 words 中以 ["foo","bar","the"] 顺序排列的连接。
子串 "barthefoo" 开始位置是 9。它是 words 中以 ["bar","the","foo"] 顺序排列的连接。
子串 "thefoobar" 开始位置是 12。它是 words 中以 ["the","foo","bar"] 顺序排列的连接。



 提示：


 1 <= s.length <= 10⁴
 1 <= words.length <= 5000
 1 <= words[i].length <= 30
 words[i] 和 s 由小写英文字母组成


 Related Topics 哈希表 字符串 滑动窗口 👍 1274 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

/*
首先分析问题：
1.给定字符串 s 和单词数组 words
2.所有单词长度相同
3.需要找出 s 中所有由 words 中所有单词以任意顺序连接而成的子串的起始索引
*/
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordLen * wordCount
	sLen := len(s)

	if sLen < totalLen {
		return []int{}
	}

	// 统计words中每个单词出现的次数
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	result := []int{}

	// 对每个可能的起始位置进行检查
	for i := 0; i < wordLen; i++ {
		left := i
		right := i
		currentMap := make(map[string]int)
		count := 0

		for right+wordLen <= sLen {
			// 获取当前窗口右边的单词
			word := s[right : right+wordLen]
			right += wordLen

			if _, exists := wordMap[word]; !exists {
				// 如果单词不在words中，重置窗口
				currentMap = make(map[string]int)
				count = 0
				left = right
				continue
			}

			// 更新当前窗口中单词的计数
			currentMap[word]++
			count++

			// 如果当前单词出现次数超过要求，移动左边界
			for currentMap[word] > wordMap[word] {
				leftWord := s[left : left+wordLen]
				currentMap[leftWord]--
				count--
				left += wordLen
			}

			// 如果窗口中单词数量等于words数量，找到一个解
			if count == wordCount {
				result = append(result, left)
			}
		}
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSubstringWithConcatenationOfAllWords(t *testing.T) {
	// 测试用例
	tests := []struct {
		s      string
		words  []string
		expect []int
	}{
		{
			s:      "barfoothefoobarman",
			words:  []string{"foo", "bar"},
			expect: []int{0, 9},
		},
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "word"},
			expect: []int{},
		},
		{
			s:      "barfoofoobarthefoobarman",
			words:  []string{"bar", "foo", "the"},
			expect: []int{6, 9, 12},
		},
	}

	for i, test := range tests {
		result := findSubstring(test.s, test.words)
		if !equal(result, test.expect) {
			t.Errorf("Test %d failed. Expected %v, got %v", i, test.expect, result)
		}
	}
}

// 辅助函数：比较两个整数切片是否相等
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
