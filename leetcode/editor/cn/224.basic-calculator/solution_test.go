package leetcode

// 224.基本计算器

import (
	"testing"
)

/**
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

 注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。



 示例 1：


输入：s = "1 + 1"
输出：2


 示例 2：


输入：s = " 2-1 + 2 "
输出：3


 示例 3：


输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23




 提示：


 1 <= s.length <= 3 * 10⁵
 s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
 s 表示一个有效的表达式
 '+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
 '-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
 输入中不存在两个连续的操作符
 每个数字和运行的计算将适合于一个有符号的 32位 整数


 Related Topics 栈 递归 数学 字符串 👍 1161 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/*
这个问题要求我们实现一个基本计算器，支持加减法和括号。主要难点在于处理括号和负数。我们可以使用栈来解决：
1.使用栈存储操作符，处理括号内的子表达式
2.遇到数字时进行累加计算
3.遇到'+'或'-'时更新当前操作符
4.遇到'('时将当前结果和操作符压入栈中
5.遇到')'时从栈中弹出结果和操作符进行计算

*/
func calculate(s string) int {
	result := 0
	sign := 1        // 当前符号，1表示正，-1表示负
	stack := []int{} // 栈用于保存括号前的结果和符号

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			// 忽略空格
			continue
		case '+':
			// 设置正号
			sign = 1
		case '-':
			// 设置负号
			sign = -1
		case '(':
			// 遇到左括号，将当前结果和符号压入栈中
			stack = append(stack, result)
			stack = append(stack, sign)
			// 重置结果和符号以计算括号内的表达式
			result = 0
			sign = 1
		case ')':
			// 遇到右括号，从栈中取出符号和之前的结果
			prevSign := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prevResult := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 计算括号内的结果并合并到之前的结果中
			result = prevResult + prevSign*result
		default:
			// 处理数字
			num := 0
			// 提取完整的数字
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			i-- // 因为for循环会自增，所以需要回退一位
			result += sign * num
		}
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBasicCalculator(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
		{"-1 + 1", 0},
		{"-(1 + 2)", -3},
		{"1-(     -2)", 3},
	}

	for _, test := range tests {
		result := calculate(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %d, Got: %d", test.input, test.expected, result)
		} else {
			t.Logf("Input: %s, Result: %d - PASS", test.input, result)
		}
	}
}
