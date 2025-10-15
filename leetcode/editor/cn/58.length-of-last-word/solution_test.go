package leetcode

// 58.最后一个单词的长度

import (
	"testing"
)

/**
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。



 示例 1：


输入：s = "Hello World"
输出：5
解释：最后一个单词是“World”，长度为 5。


 示例 2：


输入：s = "   fly me   to   the moon  "
输出：4
解释：最后一个单词是“moon”，长度为 4。


 示例 3：


输入：s = "luffy is still joyboy"
输出：6
解释：最后一个单词是长度为 6 的“joyboy”。




 提示：


 1 <= s.length <= 10⁴
 s 仅有英文字母和空格 ' ' 组成
 s 中至少存在一个单词


 Related Topics 字符串 👍 777 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLastWord(s string) int {
	// 从字符串末尾开始
	i := len(s) - 1

	// 跳过末尾的空格
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// 计算最后一个单词的长度
	length := 0
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLengthOfLastWord(t *testing.T) {
	// 测试用例
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"a", 1},
		{"a ", 1},
		{" a", 1},
	}

	for _, test := range tests {
		result := lengthOfLastWord(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %d, Got: %d", test.input, test.expected, result)
		}
	}
}
