package interview

// 32. 最长有效括号
// 栈解法：栈底存"最后一个未匹配的右括号下标"，遇到左括号入栈，遇到右括号出栈
// 时间：O(n) | 空间：O(n)
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLen := 0
	stack := []int{-1} // 栈底：最后一个不匹配的右括号下标

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 左括号下标入栈
			stack = append(stack, i)
		} else {
			// 右括号，出栈
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 栈空：没有匹配的左括号，把当前右括号下标记为新的"不匹配基准"
				stack = append(stack, i)
			} else {
				// 有效长度 = 当前下标 - 栈顶（栈顶是最后一个未匹配的左括号下标）
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}
