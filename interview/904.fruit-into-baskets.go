package interview

// 904. 水果成篮（Fruit Into Baskets）
// 变长窗口：窗口内最多两种水果，map 统计数量

func totalFruit(fruits []int) int {
    window := make(map[int]int)
    L := 0
    maxLen := 0

    for R := 0; R < len(fruits); R++ {
        window[fruits[R]]++

        // 种类超过 2，收缩左边界
        for len(window) > 2 {
            window[fruits[L]]--
            if window[fruits[L]] == 0 {
                delete(window, fruits[L])
            }
            L++
        }

        if R-L+1 > maxLen {
            maxLen = R - L + 1
        }
    }

    return maxLen
}
