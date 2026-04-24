package interview

import "testing"

func TestMaxVowels(t *testing.T) {
    tests := []struct {
        name     string
        s        string
        k        int
        expected int
    }{
        {"正常-官方示例", "abciiidef", 3, 3},
        {"正常-全元音", "aeiou", 2, 2},
        {"边界-k=1", "abc", 1, 1},
        {"边界-k等于长度", "aeiou", 5, 5},
        {"极端-全元音", "aeiouaeiou", 5, 5},
        {"极端-无元音", "bcdfg", 3, 0},
        {"正常-混合", "python", 2, 1},
        {"正常-多元音", "leetcode", 3, 2},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := maxVowels(tt.s, tt.k)
            if got != tt.expected {
                t.Errorf("maxVowels(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.expected)
            }
        })
    }
}
