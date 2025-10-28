package leetcode

// 20.有效的括号

import (
	"fmt"
	"testing"
)

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

 有效字符串需满足：


 左括号必须用相同类型的右括号闭合。
 左括号必须以正确的顺序闭合。
 每个右括号都有一个对应的相同类型的左括号。




 示例 1：


 输入：s = "()"


 输出：true

 示例 2：


 输入：s = "()[]{}"


 输出：true

 示例 3：


 输入：s = "(]"


 输出：false

 示例 4：


 输入：s = "([])"


 输出：true

 示例 5：


 输入：s = "([)]"


 输出：false



 提示：


 1 <= s.length <= 10⁴
 s 仅由括号 '()[]{}' 组成


 Related Topics 栈 字符串 👍 4826 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	hash := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidParentheses(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
