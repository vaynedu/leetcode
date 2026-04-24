package interview

import "testing"

func TestMaxSlidingWindow(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        k        int
        expected []int
    }{
        {"正常-官方示例", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
        {"正常-单元素", []int{1}, 1, []int{1}},
        {"正常-k=1", []int{1, 3, -1, -3, 5, 3, 6, 7}, 1, []int{1, 3, -1, -3, 5, 3, 6, 7}},
        {"边界-空数组", []int{}, 0, []int{}},
        {"正常-k等于长度", []int{1, 2, 3, 4, 5}, 5, []int{5}},
        {"正常-递减数组", []int{5, 4, 3, 2, 1}, 3, []int{5, 4, 3}},
        {"正常-递增数组", []int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5}},
        {"极端-全部相同", []int{1, 1, 1, 1, 1}, 3, []int{1, 1, 1}},
        {"正常-有负数", []int{-1, -3, -5, -7}, 2, []int{-1, -3, -5}},
        {"正常-正负混合", []int{5, -2, 10, 8, 3}, 3, []int{10, 10, 10}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := maxSlidingWindow(tt.nums, tt.k)
            if len(got) != len(tt.expected) {
                t.Errorf("maxSlidingWindow(%v, %d) len = %d, want %d", tt.nums, tt.k, len(got), len(tt.expected))
                return
            }
            for i := range got {
                if got[i] != tt.expected[i] {
                    t.Errorf("maxSlidingWindow(%v, %d)[%d] = %d, want %d", tt.nums, tt.k, i, got[i], tt.expected[i])
                }
            }
        })
    }
}
