package interview

import (
    "testing"
)

func TestGenerateParenthesis(t *testing.T) {
    tests := []struct {
        name string
        n    int
        want int // 数量用于验证
    }{
        {name: "n=1", n: 1, want: 1},
        {name: "n=2", n: 2, want: 2},
        {name: "n=3", n: 3, want: 5},
        {name: "n=4", n: 4, want: 14},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := generateParenthesis(tt.n)
            if len(got) != tt.want {
                t.Errorf("generateParenthesis(%d) returned %d combos, want %d", tt.n, len(got), tt.want)
            }
            // 验证每个括号串的有效性
            for _, s := range got {
                if !isValidParen(s) {
                    t.Errorf("generateParenthesis(%d) produced invalid string: %s", tt.n, s)
                }
            }
        })
    }
}

func isValidParen(s string) bool {
    cnt := 0
    for _, c := range s {
        if c == '(' {
            cnt++
        } else {
            cnt--
        }
        if cnt < 0 {
            return false
        }
    }
    return cnt == 0
}
