package leetcode

// 14.最长公共前缀

import (
	"testing"
)

/**
编写一个函数来查找字符串数组中的最长公共前缀。

 如果不存在公共前缀，返回空字符串 ""。



 示例 1：


输入：strs = ["flower","flow","flight"]
输出："fl"


 示例 2：


输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。



 提示：


 1 <= strs.length <= 200
 0 <= strs[i].length <= 200
 strs[i] 如果非空，则仅由小写英文字母组成


 Related Topics 字典树 数组 字符串 👍 3409 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	// 边界条件检查
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准进行比较
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		// 检查其他字符串在位置i是否也有相同的字符
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度不够或者字符不匹配
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	// 如果第一个字符串完全匹配，则它本身就是公共前缀
	return strs[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"a"}, "a"},
		{[]string{"ab", "a"}, "a"},
		{[]string{}, ""},
		{[]string{"same", "same", "same"}, "same"},
	}

	for _, test := range tests {
		result := longestCommonPrefix(test.input)
		if result != test.expected {
			t.Errorf("Input: %v, Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}
