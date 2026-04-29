# 堆 / 优先队列 · 题目索引

> 学完 `heap.md` 后，用这些题训练“堆里保存哪些当前候选”。

## 模板要点

```text
Top K：小根堆保存当前最大的 K 个，堆顶是第 K 大
多路归并：每一路放一个当前候选，弹出后放入该路后继
双堆：大根堆保存左半，小根堆保存右半，维持数量平衡
候选扩展：弹出当前最优后，把不会重复的新候选入堆
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 23 | [合并 K 个升序链表](./../../interview/23.merge-k-sorted-lists.go) | 困难 | ★★★ | 多路归并 | `go test -run TestMergeKLists` |
| 215 | [数组中的第 K 个最大元素](./../../interview/215.kth-largest-element-in-an-array.go) | 中等 | ★★★ | Top K/小根堆 | `go test -run TestFindKthLargest` |
| 295 | [数据流的中位数](./../../interview/295.find-median-from-data-stream.go) | 困难 | ★★★ | 双堆 | `go test -run TestMedianFinder` |
| 373 | [查找和最小的 K 对数字](./../../interview/373.find-k-pairs-with-smallest-sums.go) | 中等 | ★★☆ | 候选扩展 | `go test -run TestKSmallestPairs` |

## 学习路径建议

```text
第一步：215（Top K 堆大小不变量）
  ↓
第二步：23（多路归并，每路一个候选）
  ↓
第三步：373（候选扩展与去重）
  ↓
第四步：295（双堆平衡）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./interview/ -run 'Test(MergeKLists|FindKthLargest|MedianFinder|KSmallestPairs)' -count=1
```
