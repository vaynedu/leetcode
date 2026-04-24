package interview

import "testing"

func TestSearch704(t *testing.T) {
    tests := []struct {
        name   string
        nums   []int
        target int
        want   int
    }{
        {"存在-中间", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
        {"存在-首", []int{-1, 0, 3, 5, 9, 12}, -1, 0},
        {"存在-尾", []int{-1, 0, 3, 5, 9, 12}, 12, 5},
        {"不存在-大", []int{-1, 0, 3, 5, 9, 12}, 15, -1},
        {"不存在-小", []int{-1, 0, 3, 5, 9, 12}, -10, -1},
        {"不存在-中", []int{-1, 0, 3, 5, 9, 12}, 4, -1},
        {"单元素-存在", []int{5}, 5, 0},
        {"单元素-不存在", []int{5}, 3, -1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := search704(tt.nums, tt.target)
            if got != tt.want {
                t.Errorf("search704(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
            }
        })
    }
}
