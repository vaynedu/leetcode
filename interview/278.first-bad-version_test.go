package interview

import "testing"

var badVersion = 4

func isBadVersion(n int) bool {
    return n >= badVersion
}

func TestFirstBadVersion(t *testing.T) {
    tests := []struct {
        name string
        n int
        want int
    }{
        {"基本", 5, 4},
        {"第一个", 1, 1},
        {"全是", 3, 1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            badVersion = 4
            got := firstBadVersion(tt.n)
            if got != tt.want {
                t.Errorf("firstBadVersion(%d) = %d, want %d", tt.n, got, tt.want)
            }
        })
    }
}
