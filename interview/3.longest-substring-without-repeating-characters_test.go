package interview

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
    tests := []struct {
        name     string
        s        string
        expected int
    }{
        {"正常-官方示例1", "abcabcbb", 3},
        {"正常-官方示例2", "bbbbb", 1},
        {"正常-官方示例3", "pwwkew", 3},
        {"正常-abba", "abba", 2},
        {"边界-空字符串", "", 0},
        {"边界-单字符空格", " ", 1},
        {"边界-单字符", "a", 1},
        {"极端-26字母", "abcdefghijklmnopqrstuvwxyz", 26},
        {"极端-全相同", "aaaaaaaaaa", 1},
        {"正常-混合", "abcdeffedcba", 6},
        {"正常-内部重复", "abcdab", 4},
        {"正常-末尾重复", "abcdabc", 4},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := lengthOfLongestSubstring(tt.s)
            if got != tt.expected {
                t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.expected)
            }
        })
    }
}
