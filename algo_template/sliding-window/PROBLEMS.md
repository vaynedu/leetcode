# 滑动窗口 · 题目索引

> 学完 `sliding_window.go` 模板后，用这些题实战检验。

## 模板要点

```
三种类型：
1. 固定窗口（k已知）→ 先扩k，再滑
2. 变长窗口（极值目标）→ 双指针扩张，收缩条件
3. 变长窗口（可行性目标）→ 同上

记忆口诀：扩-算-缩 / 扩-更-缩
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 3 | [无重复最长子串](./../../interview/3.longest-substring-without-repeating-characters.go) | 中等 | ★★★ | 变长-极值 | `go test -run TestLongestSubstring` |
| 76 | [最小覆盖子串](./../../interview/76.minimum-window-substring.go) | 困难 | ★★★ | 变长-极值 | `go test -run TestMinimumWindow` |
| 239 | [滑动窗口最大值](./../../interview/239.sliding-window-maximum.go) | 困难 | ★★★ | 固定+单调队列 | `go test -run TestSlidingWindowMaximum` |
| 438 | [字母异位词](./../../interview/438.find-all-anagrams-in-a-string.go) | 中等 | ★★★ | 变长-极值 | `go test -run TestFindAnagrams` |
| 567 | [字符串排列](./../../interview/567.permutation-in-string.go) | 中等 | ★★★ | 变长-极值 | `go test -run TestCheckInclusion` |
| 643 | [最大平均子数组](./../../interview/643.maximum-average-subarray.go) | 简单 | ★★☆ | 固定窗口 | `go test -run TestMaximumAverage` |
| 904 | [水果成篮](./../../interview/904.fruit-into-baskets.go) | 中等 | ★★☆ | 变长-极值 | `go test -run TestTotalFruit` |
| 1456 | [定长子串元音数](./../../interview/1456.maximum-number-of-vowels-in-a-substring-of-given-length.go) | 中等 | ★★☆ | 固定窗口 | `go test -run TestMaxVowels` |

## 学习路径建议

```
第一步：643 / 1456（固定窗口，最简单）
  ↓
第二步：3 / 438 / 567（变长窗口-极值，hash计数）
  ↓
第三步：76（最小覆盖子串，hash计数+双指针）
  ↓
第四步：239（固定窗口+单调队列，hard）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./algo_template/sliding-window/ -v
```
