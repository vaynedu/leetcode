package interview

// 93. 复原 IP 地址
// 回溯：分段，每段 1-3 位，检查有效性
func restoreIpAddresses(s string) []string {
    result := []string{}
    if len(s) < 4 || len(s) > 12 {
        return result
    }
    var backtrack func(s string, parts int, path []string)
    backtrack = func(s string, parts int, path []string) {
        if parts == 4 {
            if len(s) == 0 {
                result = append(result, join(path))
            }
            return
        }
        for i := 1; i <= 3 && i <= len(s); i++ {
            seg := s[:i]
            if validSegment(seg) {
                backtrack(s[i:], parts+1, append(path, seg))
            }
        }
    }
    backtrack(s, 0, []string{})
    return result
}

func validSegment(s string) bool {
    if len(s) > 1 && s[0] == '0' {
        return false
    }
    v := 0
    for _, c := range s {
        v = v*10 + int(c-'0')
    }
    return v <= 255
}

func join(parts []string) string {
    result := parts[0]
    for i := 1; i < len(parts); i++ {
        result += "." + parts[i]
    }
    return result
}
