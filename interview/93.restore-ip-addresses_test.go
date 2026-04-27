package interview

import "testing"

func TestRestoreIpAddresses(t *testing.T) {
    tests := []struct {
        name string
        s string
        want int
    }{
        {"基本", "25525511135", 2},
        {"无解", "0000", 1},
        {"四段", "1111", 1},
        {"有解", "101023", 2},
        {"短了", "12", 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := restoreIpAddresses(tt.s)
            if len(got) != tt.want {
                t.Errorf("restoreIpAddresses(%s) returned %d, want %d", tt.s, len(got), tt.want)
            }
        })
    }
}
