package interview

import "testing"

func TestMinWindow(t *testing.T) {
    tests := []struct {
        name     string
        s        string
        t        string
        expected string
    }{
        {"正常-官方示例", "ADOBECODEBANC", "ABC", "BANC"},
        {"正常-单字符", "a", "a", "a"},
        {"正常-包含所有", "AB", "AB", "AB"},
        {"不存在", "a", "aa", ""},
        {"边界-s为空", "", "A", ""},
        {"边界-t为空", "A", "", ""},
        {"正常-重复字符", "AAABBB", "AB", "AB"},
        {"正常-头尾匹配", "XYZABCXYZ", "XYZ", "XYZ"},
        {"正常-中间匹配", "XXXYYYZZZ", "YZ", "YZ"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := minWindow(tt.s, tt.t)
            if got != tt.expected {
                t.Errorf("minWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.expected)
            }
        })
    }
}
