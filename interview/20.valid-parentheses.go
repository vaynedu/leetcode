package interview

// 20. 有效括号（Valid Parentheses）
// 栈：左括号进栈，右括号匹配出栈，最终栈空为有效

func isValid(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

    for _, c := range s {
        if c == '(' || c == '[' || c == '{' {
            stack = append(stack, c) // 左括号进栈
        } else {
            // 右括号：栈空或不匹配 → 无效
            if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
                return false
            }
            stack = stack[:len(stack)-1] // 出栈
        }
    }
    return len(stack) == 0 // 栈空说明全部匹配
}
