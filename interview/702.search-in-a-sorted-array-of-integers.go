package interview

// 702. 搜索有序数组（未知长度）
// 先用双倍指针找到上界，再二分
func search(reader ArrayReader, target int) int {
    // 找上界：翻倍直到超出或找到目标
    if reader.Get(0) == target {
        return 0
    }
    left, right := 0, 1
    for reader.Get(right) < target {
        left = right
        right *= 2
    }
    // 在 [left, right] 内二分
    for left <= right {
        mid := left + (right-left)/2
        val := reader.Get(mid)
        if val == target {
            return mid
        } else if val < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

// ArrayReader 接口（题目提供）
type ArrayReader interface {
    Get(i int) int
}
