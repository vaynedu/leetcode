package leetcode

// 290.单词规律

import (
	"fmt"
	"strings"
	"testing"
)

/**
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。具体来说：


 pattern 中的每个字母都 恰好 映射到 s 中的一个唯一单词。
 s 中的每个唯一单词都 恰好 映射到 pattern 中的一个字母。
 没有两个字母映射到同一个单词，也没有两个单词映射到同一个字母。




 示例1:


输入: pattern = "abba", s = "dog cat cat dog"
输出: true

 示例 2:


输入:pattern = "abba", s = "dog cat cat fish"
输出: false

 示例 3:


输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false



 提示:


 1 <= pattern.length <= 300
 pattern 只包含小写英文字母
 1 <= s.length <= 3000
 s 只包含小写英文字母和 ' '
 s 不包含 任何前导或尾随对空格
 s 中每个单词都被 单个空格 分隔


 Related Topics 哈希表 字符串 👍 721 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	wTop := make(map[string]byte)
	pTow := make(map[byte]string)

	for i := 0; i < len(pattern); i++ {
		if val, ok := pTow[pattern[i]]; ok {
			if val != words[i] {
				return false
			}
		} else {
			pTow[pattern[i]] = words[i]
		}

		if val, ok := wTop[words[i]]; ok {
			if val != pattern[i] {
				return false
			}
		} else {
			wTop[words[i]] = pattern[i]
		}
	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)

func TestWordPattern(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
