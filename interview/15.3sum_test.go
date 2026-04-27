package interview

import "testing"

func TestThreeSum(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int // count of triplets
    }{
        {"基本", []int{-1, 0, 1, 2, -1, -4}, 2},
        {"全零", []int{0, 0, 0}, 1},
        {"无解", []int{1, 2, 3}, 0},
        {"两负一正", []int{-2, -1, -1, 0, 1, 2}, 3},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := threeSum(tt.nums)
            if len(got) != tt.want {
                t.Errorf("threeSum(%v) = %v, want %d triplets", tt.nums, got, tt.want)
            }
        })
    }
}
