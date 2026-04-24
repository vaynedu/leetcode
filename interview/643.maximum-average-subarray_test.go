package interview

import "testing"

func TestFindMaxAverage(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        k        int
        expected float64
    }{
        {"正常-官方示例", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
        {"正常-单元素", []int{5}, 1, 5.0},
        {"正常-k等于长度", []int{1, 2, 3}, 3, 2.0},
        {"极端-全正数", []int{1, 2, 3, 4, 5}, 3, 4.0},
        {"极端-全负数", []int{-5, -4, -3, -2}, 2, -2.5},
        {"正常-有零", []int{0, 1, 1, 0, 0}, 2, 1.0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := findMaxAverage(tt.nums, tt.k)
            if got != tt.expected {
                t.Errorf("findMaxAverage(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.expected)
            }
        })
    }
}
