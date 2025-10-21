package leetcode

// 面试题 16.26.计算器

import (
	"fmt"
	"testing"
)

/**
给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。

 表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格 。 整数除法仅保留整数部分。

 示例 1：


输入："3+2*2"
输出：7


 示例 2：


输入：" 3/2 "
输出：1

 示例 3：


输入：" 3+5 / 2 "
输出：5


 说明：


 你可以假设所给定的表达式都是有效的。
 请不要使用内置的库函数 eval。


 Related Topics 栈 数学 字符串 👍 116 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
	stack := []int{}
	num := 0
	operator := '+' // 初始操作符设为'+'

	for i := 0; i < len(s); i++ {
		// 如果是数字，则构建完整数字
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}

		// 如果遇到非数字字符（除了空格）或者到达字符串末尾，处理前面的数字
		if (s[i] < '0' || s[i] > '9') && s[i] != ' ' || i == len(s)-1 {
			// 根据前一个操作符处理数字
			switch operator {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				// 立即计算乘法
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top*num)
			case '/':
				// 立即计算除法
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top/num)
			}

			// 更新操作符
			if i < len(s)-1 {
				operator = rune(s[i])
			}
			num = 0
		}
	}

	// 将栈中所有数字相加得到最终结果
	result := 0
	for _, val := range stack {
		result += val
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCalculatorLcci(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
