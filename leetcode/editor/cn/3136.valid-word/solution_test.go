package leetcode

// 3136.有效单词

import (
	"fmt"
	"testing"
)

/**
有效单词 需要满足以下几个条件：


 至少 包含 3 个字符。
 由数字 0-9 和英文大小写字母组成。（不必包含所有这类字符。）
 至少 包含一个 元音字母 。
 至少 包含一个 辅音字母 。


 给你一个字符串 word 。如果 word 是一个有效单词，则返回 true ，否则返回 false 。

 注意：


 'a'、'e'、'i'、'o'、'u' 及其大写形式都属于 元音字母 。
 英文中的 辅音字母 是指那些除元音字母之外的字母。




 示例 1：


 输入：word = "234Adas"


 输出：true

 解释：

 这个单词满足所有条件。

 示例 2：


 输入：word = "b3"


 输出：false

 解释：

 这个单词的长度少于 3 且没有包含元音字母。

 示例 3：


 输入：word = "a3$e"


 输出：false

 解释：

 这个单词包含了 '$' 字符且没有包含辅音字母。



 提示：


 1 <= word.length <= 20
 word 由英文大写和小写字母、数字、'@'、'#' 和 '$' 组成。


 Related Topics 字符串 👍 10 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	var hasVowel = false
	var hasConsonant = false
	var count = 0

	for i := 0; i < len(word); i++ {
		c := word[i]

		// 检查字符是否合法（数字或大小写字母）
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') {
			return false
		}

		// 统一处理成小写
		lowerC := c
		if lowerC >= 'A' && lowerC <= 'Z' {
			lowerC += 'a' - 'A'
		}

		// 判断是否为元音
		if lowerC == 'a' || lowerC == 'e' || lowerC == 'i' || lowerC == 'o' || lowerC == 'u' {
			hasVowel = true
		} else if lowerC >= 'a' && lowerC <= 'z' { // 剩下的英文字母就是辅音
			hasConsonant = true
		}

		count++
	}

	// 必须满足：至少3个字符、至少一个元音、至少一个辅音
	return hasVowel && hasConsonant && count >= 3
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidWord(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志

	tests := []struct {
		input string
		want  bool
	}{
		{"234Adas", true},
		{"#zwI", false},
		{"uHBiL#", false},
	}
	for _, tt := range tests {
		got := isValid(tt.input)
		if got != tt.want {
			t.Errorf("isValid(%s) = %v, want %v", tt.input, got, tt.want)
		}
	}

}
