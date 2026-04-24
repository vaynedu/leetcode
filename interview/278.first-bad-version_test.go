package interview

import "testing"

func TestFirstBadVersion(t *testing.T) {
    tests := []struct {
        name string
        n    int
        want int
    }{
        {"正常", 10, 5},
        {"刚好是错误", 5, 5},
        {"n=1正确", 1, 1},
        {"n=1错误", 1, 1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := firstBadVersion(tt.n)
            if got != tt.want {
                t.Errorf("firstBadVersion(%d) = %d, want %d", tt.n, got, tt.want)
            }
        })
    }
}
