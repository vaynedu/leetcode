package leetcode

// 227.基本计算器 II

import (
	"testing"
)

/**
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

 整数除法仅保留整数部分。

 你可以假设给定的表达式总是有效的。所有中间结果将在 [-2³¹, 2³¹ - 1] 的范围内。

 注意：不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。



 示例 1：


输入：s = "3+2*2"
输出：7


 示例 2：


输入：s = " 3/2 "
输出：1


 示例 3：


输入：s = " 3+5 / 2 "
输出：5




 提示：


 1 <= s.length <= 3 * 10⁵
 s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
 s 表示一个 有效表达式
 表达式中的所有整数都是非负整数，且在范围 [0, 2³¹ - 1] 内
 题目数据保证答案是一个 32-bit 整数


 Related Topics 栈 数学 字符串 👍 863 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/*
这道题要求实现一个支持四则运算（加减乘除）的计算器，不使用内置函数。关键点在于处理运算符优先级：
1.乘除法优先级高于加减法，需要先计算
2.使用栈来处理运算符优先级问题
3.遇到数字时根据前一个运算符决定如何处理
4.加减法直接入栈，乘除法立即计算后再入栈
*/

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

func TestBasicCalculatorIi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"3+2*2", 7},
		{" 3/2 ", 1},
		{" 3+5 / 2 ", 5},
		{"1*2-3/4+5*6-7*8+9/10", -24}, // 更复杂的测试用例
		{"14/3*2", 8},                 // 除法后乘法
		{"1+2*3+4", 11},               // 连续乘法
		{"42", 42},                    // 单个数字
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
