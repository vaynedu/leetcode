package interview

import "testing"

func TestMySqrt(t *testing.T) {
    tests := []struct {
        name string
        x    int
        want int
    }{
        {"0", 0, 0},
        {"1", 1, 1},
        {"4", 4, 2},
        {"8", 8, 2},
        {"9", 9, 3},
        {"16", 16, 4},
        {"100", 100, 10},
        {"INT_MAX", 2147483647, 46340},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := mySqrt(tt.x)
            if got != tt.want {
                t.Errorf("mySqrt(%d) = %d, want %d", tt.x, got, tt.want)
            }
        })
    }
}
