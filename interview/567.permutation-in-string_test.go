package interview

import "testing"

func TestCheckInclusion(t *testing.T) {
    tests := []struct {
        name     string
        s1       string
        s2       string
        expected bool
    }{
        {"正常-有排列", "ab", "eidbaooo", true},
        {"正常-无排列", "ab", "eidboaooo", false},
        {"边界-s1为空", "", "abc", false},
        {"边界-s2为空", "a", "", false},
        {"正常-单字符匹配", "a", "a", true},
        {"正常-单字符不匹配", "a", "b", false},
        {"正常-abcd在xyzabcd", "abcd", "xyzabcd", true},
        {"正常-多个可能", "abc", "abcabc", true},
        {"极端-完全相同", "abc", "abc", true},
        {"正常-s1更长", "abcd", "ab", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := checkInclusion(tt.s1, tt.s2)
            if got != tt.expected {
                t.Errorf("checkInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.expected)
            }
        })
    }
}
