package sliding_window

import "testing"

// ============================================================
// 测试 1: 固定窗口
// ============================================================

func TestMaxAverageSubarray(t *testing.T) {
    tests := []struct {
        nums     []int
        k        int
        expected float64
    }{
        {[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
        {[]int{5}, 1, 5.0},
        {[]int{1, 2, 3}, 3, 2.0},
        {[]int{1, 2, 3, 4, 5}, 3, 4.0},
    }
    for _, tt := range tests {
        got := MaxAverageSubarray(tt.nums, tt.k)
        if got != tt.expected {
            t.Errorf("MaxAverageSubarray(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.expected)
        }
    }
}

func TestMaxVowels(t *testing.T) {
    tests := []struct {
        s        string
        k        int
        expected int
    }{
        {"abciiidef", 3, 3},
        {"aeiou", 2, 2},
        {"aeiou", 5, 5},
        {"bcdfg", 3, 0},
    }
    for _, tt := range tests {
        got := MaxVowels(tt.s, tt.k)
        if got != tt.expected {
            t.Errorf("MaxVowels(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.expected)
        }
    }
}

// ============================================================
// 测试 2: 变长窗口 - 双 HashMap
// ============================================================

func TestMinWindow(t *testing.T) {
    tests := []struct {
        s, t, expected string
    }{
        {"ADOBECODEBANC", "ABC", "BANC"},
        {"a", "a", "a"},
        {"a", "aa", ""},
    }
    for _, tt := range tests {
        got := MinWindow(tt.s, tt.t)
        if got != tt.expected {
            t.Errorf("MinWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.expected)
        }
    }
}

func TestCheckInclusion(t *testing.T) {
    tests := []struct {
        s1, s2   string
        expected bool
    }{
        {"ab", "eidbaooo", true},
        {"ab", "eidboaooo", false},
        {"abc", "abc", true},
        {"abcd", "ab", false},
    }
    for _, tt := range tests {
        got := CheckInclusion(tt.s1, tt.s2)
        if got != tt.expected {
            t.Errorf("CheckInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.expected)
        }
    }
}

func TestFindAnagrams(t *testing.T) {
    tests := []struct {
        s, p     string
        expected []int
    }{
        {"cbaebabacd", "abc", []int{0, 6}},
        {"abab", "ab", []int{0, 1, 2}},
        {"abcdef", "z", []int{}},
    }
    for _, tt := range tests {
        got := FindAnagrams(tt.s, tt.p)
        if len(got) != len(tt.expected) {
            t.Errorf("FindAnagrams(%q, %q) len = %d, want %d", tt.s, tt.p, len(got), len(tt.expected))
            continue
        }
        for i := range got {
            if got[i] != tt.expected[i] {
                t.Errorf("FindAnagrams(%q, %q)[%d] = %d, want %d", tt.s, tt.p, i, got[i], tt.expected[i])
            }
        }
    }
}

// ============================================================
// 测试 3: 变长窗口 - 条件收缩
// ============================================================

func TestLengthOfLongestSubstring(t *testing.T) {
    tests := []struct {
        s        string
        expected int
    }{
        {"abcabcbb", 3},
        {"bbbbb", 1},
        {"pwwkew", 3},
        {"abba", 2},
        {"", 0},
        {"abcdefghijklmnopqrstuvwxyz", 26},
    }
    for _, tt := range tests {
        got := LengthOfLongestSubstring(tt.s)
        if got != tt.expected {
            t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.expected)
        }
    }
}

func TestTotalFruit(t *testing.T) {
    tests := []struct {
        fruits   []int
        expected int
    }{
        {[]int{1, 2, 1}, 3},
        {[]int{1, 2, 3, 2, 2}, 4},
        {[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
        {[]int{1, 1, 1, 1}, 4},
        {[]int{1, 2, 1, 2, 1, 2}, 6},
    }
    for _, tt := range tests {
        got := TotalFruit(tt.fruits)
        if got != tt.expected {
            t.Errorf("TotalFruit(%v) = %d, want %d", tt.fruits, got, tt.expected)
        }
    }
}

func TestLongestOnes(t *testing.T) {
    tests := []struct {
        nums     []int
        k        int
        expected int
    }{
        {[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
        {[]int{0, 0, 1, 1, 1, 0, 1}, 1, 5},
        {[]int{1, 0, 1, 1, 0}, 0, 2},
        {[]int{1, 1, 1, 1, 1}, 0, 5},
        {[]int{0, 0, 0}, 3, 3},
        {[]int{0, 0, 0}, 1, 1},
    }
    for _, tt := range tests {
        got := LongestOnes(tt.nums, tt.k)
        if got != tt.expected {
            t.Errorf("LongestOnes(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.expected)
        }
    }
}

// ============================================================
// 测试 4: 单调 deque
// ============================================================

func TestMaxSlidingWindow(t *testing.T) {
    tests := []struct {
        nums     []int
        k        int
        expected []int
    }{
        {[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
        {[]int{1}, 1, []int{1}},
        {[]int{1, 3, -1, -3, 5, 3, 6, 7}, 1, []int{1, 3, -1, -3, 5, 3, 6, 7}},
        {[]int{5, 4, 3, 2, 1}, 3, []int{5, 4, 3}},
        {[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5}},
        {[]int{1, 1, 1, 1, 1}, 3, []int{1, 1, 1}},
    }
    for _, tt := range tests {
        got := MaxSlidingWindow(tt.nums, tt.k)
        if len(got) != len(tt.expected) {
            t.Errorf("MaxSlidingWindow(%v, %d) len = %d, want %d", tt.nums, tt.k, len(got), len(tt.expected))
            continue
        }
        for i := range got {
            if got[i] != tt.expected[i] {
                t.Errorf("MaxSlidingWindow(%v, %d)[%d] = %d, want %d", tt.nums, tt.k, i, got[i], tt.expected[i])
            }
        }
    }
}

// ============================================================
// 测试 5: 双指针
// ============================================================

func TestMaxArea(t *testing.T) {
    tests := []struct {
        height   []int
        expected int
    }{
        {[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
        {[]int{4, 3, 2, 1, 4}, 16},
        {[]int{1, 1}, 1},
        {[]int{1}, 0},
        {[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 25},
    }
    for _, tt := range tests {
        got := MaxArea(tt.height)
        if got != tt.expected {
            t.Errorf("MaxArea(%v) = %d, want %d", tt.height, got, tt.expected)
        }
    }
}
