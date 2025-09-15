package leetcode

// 1935.可以输入的最大单词数

import (
	"fmt"
	"strings"
	"testing"
)

/**
键盘出现了一些故障，有些字母键无法正常工作。而键盘上所有其他键都能够正常工作。

 给你一个由若干单词组成的字符串 text ，单词间由单个空格组成（不含前导和尾随空格）；另有一个字符串 brokenLetters ，由所有已损坏的不同字母键
组成，返回你可以使用此键盘完全输入的 text 中单词的数目。



 示例 1：

 输入：text = "hello world", brokenLetters = "ad"
输出：1
解释：无法输入 "world" ，因为字母键 'd' 已损坏。


 示例 2：

 输入：text = "leet code", brokenLetters = "lt"
输出：1
解释：无法输入 "leet" ，因为字母键 'l' 和 't' 已损坏。


 示例 3：

 输入：text = "leet code", brokenLetters = "e"
输出：0
解释：无法输入任何单词，因为字母键 'e' 已损坏。




 提示：


 1 <= text.length <= 10⁴
 0 <= brokenLetters.length <= 26
 text 由若干用单个空格分隔的单词组成，且不含任何前导和尾随空格
 每个单词仅由小写英文字母组成
 brokenLetters 由 互不相同 的小写英文字母组成


 Related Topics 哈希表 字符串 👍 29 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func canBeTypedWords(text string, brokenLetters string) int {
	var brokenMap = make(map[byte]bool)
	for i := 0; i < len(brokenLetters); i++ {
		brokenMap[brokenLetters[i]] = true
	}

	strs := strings.Split(text, " ")
	count := 0
	for i := 0; i < len(strs); i++ {
		flag := true
		for j := 0; j < len(strs[i]); j++ {
			if brokenMap[strs[i][j]] {
				flag = false
				break
			}
		}
		if flag {
			count++
		}
	}
	return count

}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumNumberOfWordsYouCanType(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
