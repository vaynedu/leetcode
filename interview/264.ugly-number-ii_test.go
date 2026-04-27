package interview

import "testing"

func TestNthUglyNumber(t *testing.T) {
    tests := []struct {
        name string
        n int
        want int
    }{
        {"第1", 1, 1},
        {"第10", 10, 12},
        {"第11", 11, 15},
        {"前5", 5, 5},
        {"第15", 15, 24},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := nthUglyNumber(tt.n)
            if got != tt.want {
                t.Errorf("nthUglyNumber(%d) = %d, want %d", tt.n, got, tt.want)
            }
        })
    }
}
