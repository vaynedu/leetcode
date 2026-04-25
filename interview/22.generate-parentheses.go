package interview

// 22. 括号生成
// 回溯：track 中 '(' 比 ')' 多时才加 ')'
func generateParenthesis(n int) []string {
    res := []string{}
    track := make([]byte, 0, 2*n)

    var backtrack func(left, right int)
    backtrack = func(left, right int) {
        if left == n && right == n {
            res = append(res, string(track))
            return
        }
        // 加左括号（一定可以加）
        if left < n {
            track = append(track, '(')
            backtrack(left+1, right)
            track = track[:len(track)-1]
        }
        // 加右括号（需要左边有剩余）
        if right < left {
            track = append(track, ')')
            backtrack(left, right+1)
            track = track[:len(track)-1]
        }
    }

    backtrack(0, 0)
    return res
}
