package interview

import "testing"

func TestTrap(t *testing.T) {
    tests := []struct {
        name string
        height []int
        want int
    }{
        {"基本", []int{0,1,0,2,1,0,1,3,2,1,2,1}, 6},
        {"全升", []int{1,2,3,4,5}, 0},
        {"全降", []int{5,4,3,2,1}, 0},
        {"盆形", []int{3,0,2,0,4}, 7},
        {"两个坑", []int{2,0,2}, 2},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := trap(tt.height)
            if got != tt.want {
                t.Errorf("trap(%v) = %d, want %d", tt.height, got, tt.want)
            }
        })
    }
}
