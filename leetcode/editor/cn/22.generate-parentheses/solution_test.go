package leetcode

// 22.括号生成

import (
	"fmt"
	"testing"
)

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



 示例 1：


输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]


 示例 2：


输入：n = 1
输出：["()"]




 提示：


 1 <= n <= 8


 Related Topics 字符串 动态规划 回溯 👍 3918 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func dfs(n int, res *[]string, path []byte, left int, right int) {

	if left > n || right > left {
		return
	}

	if n*2 == len(path) {
		*res = append(*res, string(path))
		return
	}

	dfs(n, res, []byte(string(path)+"("), left+1, right)
	dfs(n, res, []byte(string(path)+")"), left, right+1)

}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dfs(n, &res, []byte{}, 0, 0)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGenerateParentheses(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
