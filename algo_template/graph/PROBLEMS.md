# 图 / BFS / DFS / 拓扑排序 · 题目索引

> 学完 `graph.md` 后，用这些题训练“节点状态、边关系、visited 时机”。

## 模板要点

```text
网格 DFS/BFS：四方向扩展，visited 或原地标记避免重复
BFS 最短路：入队即标记，第一次到达就是最短步数
拓扑排序：入度为 0 的节点表示当前依赖已满足
图克隆：map 记录原节点到新节点，避免重复构造
染色：相邻节点必须颜色相反，冲突则失败
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 200 | [岛屿数量](./../../interview/200.number-of-islands.go) | 中等 | ★★★ | 网格 DFS/BFS | 待补测试 |
| 207 | [课程表](./../../interview/207.course-schedule.go) | 中等 | ★★★ | 拓扑排序/判环 | `go test -run TestCanFinish` |
| 210 | [课程表 II](./../../interview/210.course-schedule-ii.go) | 中等 | ★★☆ | 拓扑序输出 | `go test -run TestFindOrder` |
| 127 | [单词接龙](./../../interview/127.word-ladder.go) | 困难 | ★★★ | BFS 最短路 | `go test -run TestLadderLength` |
| 133 | [克隆图](./../../interview/133.clone-graph.go) | 中等 | ★★☆ | DFS/BFS + map | `go test -run TestCloneGraph` |
| 785 | [判断二分图](./../../interview/785.is-graph-bipartite.go) | 中等 | ★★☆ | 染色 | `go test -run TestIsBipartite` |

## 学习路径建议

```text
第一步：200（网格连通块）
  ↓
第二步：207 / 210（拓扑排序）
  ↓
第三步：127（BFS 最短步数）
  ↓
第四步：133 / 785（图结构复制与染色）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./interview/ -run 'Test(CanFinish|FindOrder|LadderLength|CloneGraph|IsBipartite)' -count=1
```
