package interview

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"标准", "(()", 2},
		{"嵌套", ")()())", 4},
		{"全匹配", "()()()", 6},
		{"空串", "", 0},
		{"左多", "((()))", 6},
		{"右多", "())(())", 4},
		{"不匹配", ")(", 0},
		{"单左", "(", 0},
		{"单右", ")", 0},
		{"长串", "(()())", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestValidParentheses(tt.s)
			if got != tt.want {
				t.Errorf("longestValidParentheses(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
