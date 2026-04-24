package interview

// 76. 最小覆盖子串（Minimum Window Substring）
// 滑动窗口 + 双 HashMap：右扩左缩，need/window + valid 判覆盖

func minWindow(s string, t string) string {
    // t 为空，直接返回空
    if len(t) == 0 {
        return ""
    }

    need := make(map[rune]int)
    window := make(map[rune]int)

    // 初始化 need：t 中每个字符需要的数量
    for _, c := range t {
        need[c]++
    }

    L, R := 0, 0
    valid := 0                    // 已满足条件的字符种类数
    minLen := len(s) + 1          // 初始化为不可能的大值
    minStr := ""                  // 最小子串

    for R < len(s) {
        // 右指针扩展
        c := rune(s[R])
        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] { // 首次满足时 valid++
                valid++
            }
        }
        R++

        // 收缩窗口：窗口已覆盖 t 时，尽量左缩
        for valid == len(need) {
            // 更新最小窗口
            if R-L < minLen {
                minLen = R - L
                minStr = s[L:R]
            }

            // 左指针收缩
            d := rune(s[L])
            if _, ok := need[d]; ok {
                if window[d] == need[d] { // 收缩后不再满足，valid--
                    valid--
                }
                window[d]--
            }
            L++
        }
    }

    return minStr
}
