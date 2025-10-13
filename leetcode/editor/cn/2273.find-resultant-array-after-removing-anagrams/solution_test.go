package leetcode

// 2273.移除字母异位词后的结果数组

import (
	"testing"
)

/**
给你一个下标从 0 开始的字符串 words ，其中 words[i] 由小写英文字符组成。

 在一步操作中，需要选出任一下标 i ，从 words 中 删除 words[i] 。其中下标 i 需要同时满足下述两个条件：


 0 < i < words.length
 words[i - 1] 和 words[i] 是 字母异位词 。


 只要可以选出满足条件的下标，就一直执行这个操作。

 在执行所有操作后，返回 words 。可以证明，按任意顺序为每步操作选择下标都会得到相同的结果。

 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。例如，"dacb" 是 "abdc" 的一个字母异位词。



 示例 1：

 输入：words = ["abba","baba","bbaa","cd","cd"]
输出：["abba","cd"]
解释：
获取结果数组的方法之一是执行下述步骤：
- 由于 words[2] = "bbaa" 和 words[1] = "baba" 是字母异位词，选择下标 2 并删除 words[2] 。
  现在 words = ["abba","baba","cd","cd"] 。
- 由于 words[1] = "baba" 和 words[0] = "abba" 是字母异位词，选择下标 1 并删除 words[1] 。
  现在 words = ["abba","cd","cd"] 。
- 由于 words[2] = "cd" 和 words[1] = "cd" 是字母异位词，选择下标 2 并删除 words[2] 。
  现在 words = ["abba","cd"] 。
无法再执行任何操作，所以 ["abba","cd"] 是最终答案。

 示例 2：

 输入：words = ["a","b","c","d","e"]
输出：["a","b","c","d","e"]
解释：
words 中不存在互为字母异位词的两个相邻字符串，所以无需执行任何操作。



 提示：


 1 <= words.length <= 100
 1 <= words[i].length <= 10
 words[i] 由小写英文字母组成


 Related Topics 数组 哈希表 字符串 排序 👍 41 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func removeAnagrams(words []string) []string {
	result := []string{words[0]} // 第一个单词总是保留

	for i := 1; i < len(words); i++ {
		if !isAnagram(result[len(result)-1], words[i]) {
			result = append(result, words[i])
		}
	}

	return result
}

// 判断两个字符串是否为字母异位词
func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// 统计每个字符出现的次数
	count := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		count[s1[i]-'a']++
		count[s2[i]-'a']--
	}

	// 检查是否所有字符计数都为0
	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindResultantArrayAfterRemovingAnagrams(t *testing.T) {
	tests := []struct {
		words    []string
		expected []string
	}{
		{
			words:    []string{"abba", "baba", "bbaa", "cd", "cd"},
			expected: []string{"abba", "cd"},
		},
		{
			words:    []string{"a", "b", "c", "d", "e"},
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			words:    []string{"a", "b", "a"},
			expected: []string{"a", "b", "a"},
		},
	}

	for i, test := range tests {
		result := removeAnagrams(test.words)
		if !equal(result, test.expected) {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, test.expected, result)
		}
	}
}

// 辅助函数：比较两个字符串切片是否相等
func equal(a, b []string) bool {
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
