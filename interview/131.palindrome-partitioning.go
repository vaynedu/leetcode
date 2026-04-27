package interview

// 131. 分割回文串
// 回溯：枚举每个分割点，判断是否为回文
func partitionPalindrome(s string) [][]string {
    result := [][]string{}
    path := []string{}
    var backtrack func(start int)
    backtrack = func(start int) {
        if start == len(s) {
            tmp := make([]string, len(path))
            copy(tmp, path)
            result = append(result, tmp)
            return
        }
        for i := start + 1; i <= len(s); i++ {
            if isPalindrome(s[start:i]) {
                path = append(path, s[start:i])
                backtrack(i)
                path = path[:len(path)-1]
            }
        }
    }
    backtrack(0)
    return result
}

func isPalindrome(s string) bool {
    i, j := 0, len(s)-1
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}
