package interview

import "testing"

func TestPartition(t *testing.T) {
    tests := []struct {
        name string
        s string
        want int
    }{
        {"基本", "aab", 2},
        {"单字符", "a", 1},
        {"全回文", "aaa", 4},
        {"无回文", "abc", 1},
        {"长", "aabb", 4},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := partitionPalindrome(tt.s)
            if len(got) != tt.want {
                t.Errorf("partition(%s) returned %d, want %d", tt.s, len(got), tt.want)
            }
        })
    }
}
