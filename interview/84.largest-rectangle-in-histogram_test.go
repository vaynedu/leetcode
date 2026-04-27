package interview

import "testing"

func TestLargestRectangleArea(t *testing.T) {
    tests := []struct {
        name string
        heights []int
        want int
    }{
        {"基本", []int{2,1,5,6,2,3}, 10},
        {"全升", []int{1,2,3,4,5}, 9},
        {"全降", []int{5,4,3,2,1}, 5},
        {"高个", []int{2,4}, 4},
        {"单柱", []int{1}, 1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := largestRectangleArea(tt.heights)
            if got != tt.want {
                t.Errorf("largestRectangleArea(%v) = %d, want %d", tt.heights, got, tt.want)
            }
        })
    }
}
