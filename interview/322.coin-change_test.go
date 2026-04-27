package interview

import "testing"

func TestCoinChange(t *testing.T) {
    tests := []struct {
        name string
        coins []int
        amount int
        want int
    }{
        {"基本", []int{1, 2, 5}, 11, 3},
        {"无解", []int{2}, 3, -1},
        {"全部1", []int{1}, 0, 0},
        {"多个", []int{1, 5, 11}, 15, 3},
        {"大金额", []int{1}, 100, 100},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := coinChange(tt.coins, tt.amount)
            if got != tt.want {
                t.Errorf("coinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
            }
        })
    }
}
