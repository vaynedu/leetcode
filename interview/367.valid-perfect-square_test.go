package interview

import "testing"

func TestIsPerfectSquare(t *testing.T) {
    tests := []struct {
        name string
        num int
        want bool
    }{
        {"16", 16, true},
        {"14", 14, false},
        {"1", 1, true},
        {"0", 0, true},
        {"100", 100, true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := isPerfectSquare(tt.num)
            if got != tt.want {
                t.Errorf("isPerfectSquare(%d) = %v, want %v", tt.num, got, tt.want)
            }
        })
    }
}
