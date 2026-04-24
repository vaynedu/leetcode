package interview

// 1456. 定长子串中元音的最大数目
// 固定窗口：元音判断，滑动统计

func maxVowels(s string, k int) int {
    vowel := map[rune]bool{
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
    }

    // 第一个窗口
    count := 0
    for i := 0; i < k && i < len(s); i++ {
        if vowel[rune(s[i])] {
            count++
        }
    }
    maxCount := count

    // 滑动窗口：右移减左加右
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
