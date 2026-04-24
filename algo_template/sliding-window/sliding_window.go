package sliding_window

// ============================================================
// 滑动窗口模板 - 通用实现
// ============================================================

// --------------------------------------------------
// 类型1: 固定窗口（窗口大小固定为 k）
// 适用于：子数组最大平均数、元音计数等
// --------------------------------------------------

// MaxAverageSubarray 643. 子数组最大平均数
// 固定窗口：先算首窗和，右移减左加右
func MaxAverageSubarray(nums []int, k int) float64 {
    sum := 0
    for i := 0; i < k; i++ {
        sum += nums[i]
    }
    maxSum := sum
    for i := k; i < len(nums); i++ {
        sum += nums[i] - nums[i-k]
        if sum > maxSum {
            maxSum = sum
        }
    }
    return float64(maxSum) / float64(k)
}

// MaxVowels 1456. 定长子串中元音的最大数目
// 固定窗口：元音 set，滑动统计
func MaxVowels(s string, k int) int {
    vowel := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
    count := 0
    for i := 0; i < k && i < len(s); i++ {
        if vowel[rune(s[i])] {
            count++
        }
    }
    maxCount := count
    for i := k; i < len(s); i++ {
        if vowel[rune(s[i])] {
            count++
        }
        if vowel[rune(s[i-k])] {
            count--
        }
        if count > maxCount {
            maxCount = count
        }
    }
    return maxCount
}

// --------------------------------------------------
// 类型2: 变长窗口 - 双 HashMap + valid 计数
// 适用于：最小覆盖子串、字符串排列、字母异位词
// --------------------------------------------------

// MinWindow 76. 最小覆盖子串
// 变长窗口：need/window 双 map，valid 判覆盖
func MinWindow(s string, t string) string {
    if len(t) == 0 {
        return ""
    }
    need := make(map[rune]int)
    window := make(map[rune]int)
    for _, c := range t {
        need[c]++
    }
    L, R := 0, 0
    valid := 0
    minLen := len(s) + 1
    minStr := ""
    for R < len(s) {
        c := rune(s[R])
        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }
        R++
        for valid == len(need) {
            if R-L < minLen {
                minLen = R - L
                minStr = s[L:R]
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
    return minStr
}

// CheckInclusion 567. 字符串的排列
// 变长窗口（固定长度）：找到立即返回
func CheckInclusion(s1 string, s2 string) bool {
    need := make(map[rune]int)
    window := make(map[rune]int)
    for _, c := range s1 {
        need[c]++
    }
    L, R := 0, 0
    valid := 0
    for R < len(s2) {
        c := rune(s2[R])
        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }
        R++
        if R-L == len(s1) {
            if valid == len(need) {
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

// FindAnagrams 438. 找到所有字母异位词
// 变长窗口（固定长度）：记录所有起始位置
func FindAnagrams(s string, p string) []int {
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

// --------------------------------------------------
// 类型3: 变长窗口 - 条件收缩
// 适用于：无重复最长子串、水果成篮、最大连续1
// --------------------------------------------------

// LengthOfLongestSubstring 3. 无重复最长子串
// 变长窗口：last map 记录字符上次位置，L = max(L, pos+1)
func LengthOfLongestSubstring(s string) int {
    last := make(map[rune]int)
    L := 0
    maxLen := 0
    for R, c := range s {
        if pos, exists := last[c]; exists && pos >= L {
            L = pos + 1
        }
        last[c] = R
        if R-L+1 > maxLen {
            maxLen = R - L + 1
        }
    }
    return maxLen
}

// TotalFruit 904. 水果成篮
// 变长窗口：map 统计种类，超过2收缩
func TotalFruit(fruits []int) int {
    window := make(map[int]int)
    L := 0
    maxLen := 0
    for R := 0; R < len(fruits); R++ {
        window[fruits[R]]++
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

// LongestOnes 1004. 最大连续1的个数 III
// 变长窗口：记录0的数量，超过K收缩
func LongestOnes(nums []int, k int) int {
    L := 0
    zeroCount := 0
    maxLen := 0
    for R := 0; R < len(nums); R++ {
        if nums[R] == 0 {
            zeroCount++
        }
        for zeroCount > k {
            if nums[L] == 0 {
                zeroCount--
            }
            L++
        }
        if R-L+1 > maxLen {
            maxLen = R - L + 1
        }
    }
    return maxLen
}

// --------------------------------------------------
// 类型4: 单调 deque
// 适用于：滑动窗口最大值
// --------------------------------------------------

// MaxSlidingWindow 239. 滑动窗口最大值
// 单调 deque：存下标保持递减，队首是当前窗口最大
func MaxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 {
        return []int{}
    }
    var deque []int
    result := []int{}
    for R := 0; R < len(nums); R++ {
        for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[R] {
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, R)
        if deque[0] <= R-k {
            deque = deque[1:]
        }
        if R >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }
    return result
}

// --------------------------------------------------
// 类型5: 双指针（本质是变长窗口特例）
// 适用于：装水容器
// --------------------------------------------------

// MaxArea 11. 装水容器
// 双指针：两端向中间收敛，移动短边
func MaxArea(height []int) int {
    left, right := 0, len(height)-1
    maxArea := 0
    for left < right {
        w := right - left
        h := height[left]
        if height[right] < h {
            h = height[right]
        }
        area := w * h
        if area > maxArea {
            maxArea = area
        }
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return maxArea
}
