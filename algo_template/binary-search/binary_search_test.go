package binary_search

import "testing"

// ============================================================
// 测试 1: 标准二分查找（704. 二分查找）
// ============================================================

func TestBinarySearch(t *testing.T) {
    tests := []struct {
        name   string
        nums   []int
        target int
        want   int
    }{
        {"正常-中间存在", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
        {"正常-首元素", []int{-1, 0, 3, 5, 9, 12}, -1, 0},
        {"正常-尾元素", []int{-1, 0, 3, 5, 9, 12}, 12, 5},
        {"不存在-大于所有", []int{-1, 0, 3, 5, 9, 12}, 15, -1},
        {"不存在-小于所有", []int{-1, 0, 3, 5, 9, 12}, -10, -1},
        {"不存在-在中间", []int{-1, 0, 3, 5, 9, 12}, 4, -1},
        {"单元素-存在", []int{5}, 5, 0},
        {"单元素-不存在", []int{5}, 3, -1},
        {"两个元素-找到左", []int{1, 3}, 1, 0},
        {"两个元素-找到右", []int{1, 3}, 3, 1},
        {"两个元素-不存在", []int{1, 3}, 2, -1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := BinarySearch(tt.nums, tt.target)
            if got != tt.want {
                t.Errorf("BinarySearch(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
            }
        })
    }
}

// ============================================================
// 测试 2: 左侧边界/插入位置（35. 搜索插入位置）
// ============================================================

func TestLowerBound(t *testing.T) {
    tests := []struct {
        name   string
        nums   []int
        target int
        want   int
    }{
        {"正常-中间插入", []int{1, 3, 5, 6}, 5, 2},
        {"正常-头部插入", []int{1, 3, 5, 6}, 0, 0},
        {"正常-尾部插入", []int{1, 3, 5, 6}, 7, 4},
        {"正常-中间插入2", []int{1, 3, 5, 6}, 2, 1},
        {"不存在-插中间", []int{1, 3, 5, 6}, 4, 2},
        {"重复-第一个", []int{1, 1, 1, 1}, 1, 0},
        {"单元素-存在", []int{1}, 1, 0},
        {"单元素-小", []int{1}, 0, 0},
        {"单元素-大", []int{1}, 2, 1},
        {"空数组", []int{}, 1, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := LowerBound(tt.nums, tt.target)
            if got != tt.want {
                t.Errorf("LowerBound(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
            }
        })
    }
}

// ============================================================
// 测试 3: 第一个错误版本（278. 第一个错误的版本）
// ============================================================

func TestFirstBadVersion(t *testing.T) {
    tests := []struct {
        name string
        n    int
        want int
    }{
        {"正常-版本5是第一个错误", 10, 5},
        {"刚好到错误版本", 5, 5},
        {"只有第一个版本且正确", 1, 1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := FirstBadVersion(tt.n)
            if got != tt.want {
                t.Errorf("FirstBadVersion(%d) = %d, want %d", tt.n, got, tt.want)
            }
        })
    }
}

// ============================================================
// 测试 4: x 的平方根（69. x 的平方根）
// ============================================================

func TestMySqrt(t *testing.T) {
    tests := []struct {
        name string
        x    int
        want int
    }{
        {"正常4", 4, 2},
        {"正常8", 8, 2},
        {"正常9", 9, 3},
        {"正常0", 0, 0},
        {"正常1", 1, 1},
        {"正常16", 16, 4},
        {"正常100", 100, 10},
        {"正常2147395599", 2147395599, 46339},
        {"正常2147483647", 2147483647, 46340},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := MySqrt(tt.x)
            if got != tt.want {
                t.Errorf("MySqrt(%d) = %d, want %d", tt.x, got, tt.want)
            }
        })
    }
}

// ============================================================
// 测试 5: 有效的完全平方数（367. 有效的完全平方数）
// ============================================================

func TestValidPerfectSquare(t *testing.T) {
    tests := []struct {
        name string
        num  int
        want bool
    }{
        {"16", 16, true},
        {"1", 1, true},
        {"0", 0, true},
        {"4", 4, true},
        {"9", 9, true},
        {"14", 14, false},
        {"15", 15, false},
        {"2", 2, false},
        {"3", 3, false},
        {"5", 5, false},
        {"100", 100, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := ValidPerfectSquare(tt.num)
            if got != tt.want {
                t.Errorf("ValidPerfectSquare(%d) = %v, want %v", tt.num, got, tt.want)
            }
        })
    }
}

// ============================================================
// 测试 6: 搜索旋转排序数组（33. 搜索旋转排序数组）
// ============================================================

func TestSearchRotated(t *testing.T) {
    tests := []struct {
        name   string
        nums   []int
        target int
        want   int
    }{
        {"正常-中间", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
        {"正常-首", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
        {"正常-尾", []int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
        {"不存在", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
        {"未旋转-正常", []int{1, 2, 3, 4, 5}, 3, 2},
        {"未旋转-不存在", []int{1, 2, 3, 4, 5}, 6, -1},
        {"两个元素-找到", []int{3, 1}, 1, 1},
        {"两个元素-不存在", []int{3, 1}, 2, -1},
        {"单元素-存在", []int{1}, 1, 0},
        {"单元素-不存在", []int{1}, 0, -1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := SearchRotated(tt.nums, tt.target)
            if got != tt.want {
                t.Errorf("SearchRotated(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
            }
        })
    }
}

// ============================================================
// 测试 7: 寻找旋转排序数组中的最小值（153. 寻找旋转排序数组中的最小值）
// ============================================================

func TestFindMin(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"正常", []int{3, 4, 5, 1, 2}, 1},
        {"正常2", []int{4, 5, 6, 7, 0, 1, 2}, 0},
        {"未旋转", []int{1, 2, 3, 4, 5}, 1},
        {"两个元素", []int{2, 1}, 1},
        {"两个元素2", []int{1, 2}, 1},
        {"单元素", []int{1}, 1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := FindMin(tt.nums)
            if got != tt.want {
                t.Errorf("FindMin(%v) = %d, want %d", tt.nums, got, tt.want)
            }
        })
    }
}

// ============================================================
// 测试 8: 寻找峰值（162. 寻找峰值）
// ============================================================

func TestFindPeakElement(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"正常", []int{1, 2, 3, 1}, 2},
        {"递增", []int{1, 2, 3, 4, 5}, 4},
        {"递减", []int{5, 4, 3, 2, 1}, 0},
        {"两个元素-升", []int{1, 2}, 1},
        {"两个元素-降", []int{2, 1}, 0},
        {"两个元素-相等", []int{1, 1}, 0},
        {"单元素", []int{1}, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := FindPeakElement(tt.nums)
            if got != tt.want {
                t.Errorf("FindPeakElement(%v) = %d, want %d", tt.nums, got, tt.want)
            }
        })
    }
}
