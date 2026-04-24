package interview

import "testing"

func TestMaxArea(t *testing.T) {
    tests := []struct {
        name     string
        height   []int
        expected int
    }{
        {"正常情况-官方示例", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
        {"正常情况-两端对称", []int{4, 3, 2, 1, 4}, 16},
        {"边界-只有两个柱子", []int{1, 1}, 1},
        {"边界-只有一个柱子", []int{1}, 0},
        {"极端-升序排列", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 25},
        {"极端-降序排列", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 25},
        {"极端-全部相同", []int{5, 5, 5, 5, 5}, 20},
        {"边界-首尾最高", []int{1, 2, 1}, 2},
        {"正常-中间高两边低", []int{1, 8, 6, 2, 5, 4, 8, 25}, 48},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := maxArea(tt.height)
            if got != tt.expected {
                t.Errorf("maxArea(%v) = %d, want %d", tt.height, got, tt.expected)
            }
        })
    }
}
