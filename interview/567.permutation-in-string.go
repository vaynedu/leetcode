package interview

// 567. 字符串的排列（Permutation in String）
// 变长窗口（固定长度）：双 HashMap + valid 计数

func checkInclusion(s1 string, s2 string) bool {
    need := make(map[rune]int)
    window := make(map[rune]int)

    for _, c := range s1 {
        need[c]++
    }

    L, R := 0, 0
    valid := 0
    needLen := len(need)

    for R < len(s2) {
        c := rune(s2[R])
        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }
        R++

        // 窗口长度等于 s1 时，检查并收缩
        if R-L == len(s1) {
            if valid == needLen {
                return true
            }
            d := rune(s2[L])
            if _, ok := need[d]; ok {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
            L++
        }
    }

    return false
}
