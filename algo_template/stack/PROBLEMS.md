# 栈与单调栈 · 题目索引

> 学完 `stack.md` 后，用这些题训练“栈里保留的未决状态是什么”。

## 模板要点

```text
普通栈：保存最近还没匹配的元素
下标栈：保存边界位置，常用于长度/宽度计算
单调栈：栈内按值单调，弹栈时结算被弹元素
哨兵：在开头或结尾补一个边界，统一清算逻辑
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 20 | [有效括号](./../../interview/20.valid-parentheses.go) | 简单 | ★★★ | 普通栈/匹配 | `go test -run TestIsValid` |
| 32 | [最长有效括号](./../../interview/32.longest-valid-parentheses.go) | 困难 | ★★★ | 栈下标/base | `go test -run TestLongestValidParentheses` |
| 84 | [柱状图最大矩形](./../../interview/84.largest-rectangle-in-histogram.go) | 困难 | ★★★ | 单调递增栈 | `go test -run TestLargestRectangleArea` |
| 42 | [接雨水](./../../interview/42.trapping-rain-water.go) | 困难 | ★★★ | 单调栈/低洼区 | `go test -run TestTrap` |

## 学习路径建议

```text
第一步：20（最近未匹配）
  ↓
第二步：32（栈下标和合法段 base）
  ↓
第三步：84（弹栈时确定左右边界）
  ↓
第四步：42（单调栈结算水量）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./interview/ -run 'Test(IsValid|LongestValidParentheses|LargestRectangleArea|Trap)' -count=1
```
