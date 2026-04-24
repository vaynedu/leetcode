package interview

// 3. 无重复最长子串（Longest Substring Without Repeating Characters）
// 滑动窗口 + HashMap：右指针扩展，遇重复左指针收缩

func lengthOfLongestSubstring(s string) int {
    last := make(map[rune]int) // 记录字符上次出现的位置
    L := 0
    maxLen := 0

    for R, c := range s {
        // 如果字符在窗口内重复（pos >= L），L 收缩到重复字符之后
        if pos, exists := last[c]; exists && pos >= L {
            L = pos + 1
        }
        last[c] = R      // 更新字符位置
        if R-L+1 > maxLen {
            maxLen = R - L + 1
        }
    }
    return maxLen
}
