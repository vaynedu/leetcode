package interview

import "testing"

func TestFindAnagrams(t *testing.T) {
    tests := []struct {
        name     string
        s        string
        p        string
        expected []int
    }{
        {"正常-官方示例1", "cbaebabacd", "abc", []int{0, 6}},
        {"正常-官方示例2", "abab", "ab", []int{0, 1, 2}},
        {"边界-s为空", "", "abc", []int{}},
        {"边界-p为空", "abc", "", []int{}},
        {"边界-p更长", "abc", "abcd", []int{}},
        {"正常-单字符", "a", "a", []int{0}},
        {"正常-无匹配", "abcdef", "z", []int{}},
        {"正常-连续重复", "aaaaaaaaaaaaaa", "a", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := findAnagrams(tt.s, tt.p)
            if len(got) != len(tt.expected) {
                t.Errorf("findAnagrams(%q, %q) len = %d, want %d", tt.s, tt.p, len(got), len(tt.expected))
                return
            }
            for i := range got {
                if got[i] != tt.expected[i] {
                    t.Errorf("findAnagrams(%q, %q)[%d] = %d, want %d", tt.s, tt.p, i, got[i], tt.expected[i])
                }
            }
        })
    }
}
