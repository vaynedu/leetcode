package interview

import "testing"

func TestIsValid(t *testing.T) {
    tests := []struct {
        name     string
        s        string
        expected bool
    }{
        {"正常-简单配对", "()", true},
        {"正常-三种括号", "()[]{}", true},
        {"正常-嵌套", "{[]}", true},
        {"正常-多层嵌套", "((()))", true},
        {"正常-混合嵌套", "({}[])", true},
        {"边界-空字符串", "", true},
        {"边界-只有左括号", "(", false},
        {"边界-只有右括号", ")", false},
        {"错误-不匹配", "(]", false},
        {"错误-交叉匹配", "([)]", false},
        {"错误-交叉嵌套", "{[}]", false},
        {"错误-缺少右括号", "((())", false},
        {"错误-缺少左括号", "())", false},
        {"正常-长字符串", "()[]{}()[]{}", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := isValid(tt.s)
            if got != tt.expected {
                t.Errorf("isValid(%q) = %v, want %v", tt.s, got, tt.expected)
            }
        })
    }
}
