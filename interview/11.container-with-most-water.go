package interview

// 11. 装水容器（Container With Most Water）
// 双指针左右夹逼，短板决定高度，移动短边指针

func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxArea := 0

    for left < right {
        // 计算当前面积：宽度 × 高度（高度取较短的那根柱子）
        w := right - left
        h := height[left]
        if height[right] < h {
            h = height[right]
        }
        area := w * h
        if area > maxArea {
            maxArea = area
        }

        // 移动高度较小的指针（宽度必定减小，所以要找更高的柱子才可能增大面积）
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return maxArea
}
