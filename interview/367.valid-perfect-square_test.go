package interview

import "testing"

func TestIsPerfectSquare(t *testing.T) {
    tests := []struct {
        name string
        num  int
        want bool
    }{
        {"1", 1, true},
        {"4", 4, true},
        {"9", 9, true},
        {"16", 16, true},
        {"0", 0, true},
        {"2", 2, false},
        {"3", 3, false},
        {"14", 14, false},
        {"100", 100, true},
        {"INT_MAX", 2147483647, false},
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
