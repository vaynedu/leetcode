package leetcode

// 438.找到字符串中所有字母异位词

import (
	"testing"
)

/**
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。



 示例 1:


输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。


 示例 2:


输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。




 提示:


 1 <= s.length, p.length <= 3 * 10⁴
 s 和 p 仅包含小写字母


 👍 1753 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	result := make([]int, 0)

	// 使用数组记录字符频次（假设只有小写字母）
	pCount := [26]int{}
	windowCount := [26]int{}

	// 统计字符串p中各字符的频次
	for _, ch := range p {
		pCount[ch-'a']++
	}

	// 滑动窗口的大小
	windowSize := len(p)

	// 初始化第一个窗口
	for i := 0; i < windowSize; i++ {
		windowCount[s[i]-'a']++
	}

	// 检查第一个窗口
	if pCount == windowCount {
		result = append(result, 0)
	}

	// 滑动窗口：从第二个窗口开始
	for i := windowSize; i < len(s); i++ {
		// 添加新字符到窗口右侧
		windowCount[s[i]-'a']++

		// 移除窗口左侧的字符
		windowCount[s[i-windowSize]-'a']--

		// 检查当前窗口是否为异位词
		if pCount == windowCount {
			result = append(result, i-windowSize+1)
		}
	}

	return result
}

//p的字符频次: a:1, b:1, c:1
//
//初始窗口: "cba" (索引 0-2)
//窗口字符频次: a:1, b:1, c:1  -> 匹配! 添加索引 0
//
//滑动到: "bae" (索引 1-3)
//窗口字符频次: a:1, b:1, e:1  -> 不匹配
//
//滑动到: "aeb" (索引 2-4)
//窗口字符频次: a:1, b:1, e:1  -> 不匹配
//
//滑动到: "eba" (索引 3-5)
//窗口字符频次: a:2, b:1, e:1  -> 不匹配
//
//滑动到: "bab" (索引 4-6)
//窗口字符频次: a:1, b:2        -> 不匹配
//
//滑动到: "aba" (索引 5-7)
//窗口字符频次: a:2, b:1        -> 不匹配
//
//滑动到: "bac" (索引 6-8)
//窗口字符频次: a:1, b:1, c:1  -> 匹配! 添加索引 6
//
//滑动到: "acd" (索引 7-9)
//窗口字符频次: a:1, c:2, d:1  -> 不匹配
//
//最终结果: [0, 6]
//leetcode submit region end(Prohibit modification and deletion)

func TestFindAllAnagramsInAString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected []int
	}{
		{
			name:     "示例1",
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6},
		},
		{
			name:     "示例2",
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2},
		},
		{
			name:     "相同字符串",
			s:        "abc",
			p:        "abc",
			expected: []int{0},
		},
		{
			name:     "无匹配",
			s:        "abc",
			p:        "def",
			expected: []int{},
		},
		{
			name:     "p比s长",
			s:        "abc",
			p:        "abcdef",
			expected: []int{},
		},
		{
			name:     "重复字符",
			s:        "baa",
			p:        "aa",
			expected: []int{1},
		},
		{
			name:     "复杂情况",
			s:        "abacbabc",
			p:        "abc",
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "单字符匹配",
			s:        "aaaaaaa",
			p:        "a",
			expected: []int{0, 1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAnagrams(tt.s, tt.p)

			if !equalIntSlices(result, tt.expected) {
				t.Errorf("findAnagrams(%q, %q) = %v, expected %v",
					tt.s, tt.p, result, tt.expected)
			}
		})
	}
}

// equalIntSlices 比较两个整数切片是否相等
func equalIntSlices(a, b []int) bool {
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
