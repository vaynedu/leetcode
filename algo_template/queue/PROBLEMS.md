# 队列 · 题目索引

> 学完 `queue.go` 模板后，用这些题实战检验。

## 识别信号

- 按层处理、最短步数、从起点向外一圈圈扩散。
- 需要先进先出处理状态。
- 需要维护一个窗口内的候选顺序。

## 核心不变量

- BFS 队列中保存的是已发现但尚未处理的节点。
- 按层 BFS 时，当前层的 `size` 固定，处理完才进入下一层。
- 单调队列中，队首始终是当前窗口最优候选。

## 决策流程

```text
1. 初始状态入队并标记 visited。
2. 弹出队首，处理当前状态。
3. 枚举合法邻居，未访问则入队。
4. 如果按层计步，处理完当前 size 后 step++。
5. 单调队列则在入队前从队尾删除劣候选。
```

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
