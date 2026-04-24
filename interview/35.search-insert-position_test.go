package interview

import "testing"

func TestSearchInsert(t *testing.T) {
    tests := []struct {
        name   string
        nums   []int
        target int
        want   int
    }{
        {"插入-尾部", []int{1, 3, 5, 6}, 5, 2},
        {"插入-头部", []int{1, 3, 5, 6}, 0, 0},
        {"插入-新尾部", []int{1, 3, 5, 6}, 7, 4},
        {"插入-中间", []int{1, 3, 5, 6}, 2, 1},
        {"重复", []int{1, 1, 1, 1}, 1, 0},
        {"单元素-存在", []int{1}, 1, 0},
        {"单元素-小", []int{1}, 0, 0},
        {"单元素-大", []int{1}, 2, 1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := searchInsert(tt.nums, tt.target)
            if got != tt.want {
                t.Errorf("searchInsert(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
            }
        })
    }
}
