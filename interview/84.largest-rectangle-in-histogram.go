package interview

import "math"

// 84. 柱状图中最大的矩形
// 单调栈：维护递增高度，弹出时计算宽度
func largestRectangleArea(heights []int) int {
    heights = append(heights, 0) // 末尾加哨兵
    stack := []int{}
    maxArea := 0
    for i, h := range heights {
        for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }
            maxArea = max(maxArea, height*width)
        }
        stack = append(stack, i)
    }
    return maxArea
}
