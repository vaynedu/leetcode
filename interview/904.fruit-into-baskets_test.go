package interview

import "testing"

func TestTotalFruit(t *testing.T) {
    tests := []struct {
        name     string
        fruits   []int
        expected int
    }{
        {"正常-官方示例1", []int{1, 2, 1}, 3},
        {"正常-官方示例2", []int{1, 2, 3, 2, 2}, 4},
        {"正常-官方示例3", []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
        {"边界-单元素", []int{1}, 1},
        {"边界-两元素", []int{1, 2}, 2},
        {"边界-空数组", []int{}, 0},
        {"极端-全相同", []int{1, 1, 1, 1}, 4},
        {"极端-交替", []int{1, 2, 1, 2, 1, 2}, 6},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := totalFruit(tt.fruits)
            if got != tt.expected {
                t.Errorf("totalFruit(%v) = %d, want %d", tt.fruits, got, tt.expected)
            }
        })
    }
}
