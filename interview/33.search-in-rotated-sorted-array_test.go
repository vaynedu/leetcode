package interview

import "testing"

func TestSearch33(t *testing.T) {
    tests := []struct {
        name   string
        nums   []int
        target int
        want   int
    }{
        {"正常-找0", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
        {"正常-找4", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
        {"正常-找2", []int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
        {"不存在", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
        {"未旋转", []int{1, 2, 3, 4, 5}, 3, 2},
        {"未旋转-不存在", []int{1, 2, 3, 4, 5}, 6, -1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := search33(tt.nums, tt.target)
            if got != tt.want {
                t.Errorf("search33(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
            }
        })
    }
}
