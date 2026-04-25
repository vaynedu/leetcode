package interview

// 17. 电话号码的字母组合
// 回溯：枚举每个数字对应的所有字母
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    letters := map[byte]string{
        '2': "abc", '3': "def", '4': "ghi",
        '5': "jkl", '6': "mno", '7': "pqrs",
        '8': "tuv", '9': "wxyz",
    }
    res := []string{}
    path := make([]byte, len(digits))

    var backtrack func(idx int)
    backtrack = func(idx int) {
        if idx == len(digits) {
            res = append(res, string(path))
            return
        }
        for _, c := range letters[digits[idx]] {
            path[idx] = byte(c)
            backtrack(idx + 1)
        }
    }

    backtrack(0)
    return res
}
