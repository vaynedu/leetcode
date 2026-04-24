package interview

// 438. 找到所有字母异位词（Find All Anagrams in a String）
// 变长窗口（固定长度）：双 HashMap + valid 计数，记录所有满足条件的位置

func findAnagrams(s string, p string) []int {
    need := make(map[rune]int)
    window := make(map[rune]int)

    for _, c := range p {
        need[c]++
    }

    L, R := 0, 0
    valid := 0
    result := []int{}

    for R < len(s) {
        c := rune(s[R])
        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }
        R++

        // 窗口长度等于 p 时，检查并收缩
        if R-L == len(p) {
            if valid == len(need) {
                result = append(result, L)
            }
            d := rune(s[L])
            if _, ok := need[d]; ok {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
            L++
        }
    }

    return result
}
