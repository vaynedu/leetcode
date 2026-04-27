# 队列 · 题目索引

> 学完 `queue.go` 模板后，用这些题实战检验。

## 模板要点

```
队列（FIFO）：用于 BFS 层序遍历、滑窗记录顺序
双端队列（Deque）：头尾均可 O(1) 增删，用于单调队列

常用场景：
  - BFS 层序遍历（队列）
  - 滑动窗口最大值（单调递减 deque）
  - 约瑟夫环（队列模拟）
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 239 | [滑动窗口最大值](./../../interview/239.sliding-window-maximum.go) | 困难 | ★★★ | 单调队列 | `go test -run TestSlidingWindowMaximum` |
| 102 | [二叉树层序遍历](./../../interview/102.binary-tree-level-order-traversal.go) | 中等 | ★★★ | BFS队列 | `go test -run TestLevelOrder` |
| 200 | [岛屿数量](./../../interview/200.number-of-islands.go) | 中等 | ★★★ | BFS网格 | `go test -run TestNumIslands` |

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./algo_template/queue/ -v
```
