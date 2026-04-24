package interview

import "testing"

func TestLongestOnes(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        k        int
        expected int
    }{
        {"正常-官方示例1", []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
        {"正常-官方示例2", []int{0, 0, 1, 1, 1, 0, 1}, 1, 5},
        {"边界-k=0", []int{1, 0, 1, 1, 0}, 0, 2},
        {"边界-k够用", []int{0, 0, 0}, 3, 3},
        {"极端-全1", []int{1, 1, 1, 1, 1}, 0, 5},
        {"极端-全0-k够", []int{0, 0, 0}, 3, 3},
        {"极端-全0-k不够", []int{0, 0, 0}, 1, 1},
        {"正常-中间翻转", []int{1, 0, 0, 0, 1, 1, 1}, 2, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := longestOnes(tt.nums, tt.k)
            if got != tt.expected {
                t.Errorf("longestOnes(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.expected)
            }
        })
    }
}
