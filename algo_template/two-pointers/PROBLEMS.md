# 双指针 · 题目索引

> 学完 `two_pointers.md` 后，用这些题实战检验“每次移动指针排除了什么”。

## 模板要点

```text
对撞指针：left/right 从两端收敛，移动能排除一批候选的一侧
排序双指针：先排序，固定一个数后用 left/right 找目标和
快慢指针：快慢速度差或固定距离表达链表结构
同步指针：两个指针切换起点，用总路程抵消长度差
dummy 链表：用虚拟头统一处理删除/拼接头节点
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 11 | [盛最多水的容器](./../../interview/11.container-with-most-water.go) | 中等 | ★★★ | 对撞指针/短板 | `go test -run TestMaxArea` |
| 15 | [三数之和](./../../interview/15.3sum.go) | 中等 | ★★★ | 排序+对撞 | `go test -run TestThreeSum` |
| 42 | [接雨水](./../../interview/42.trapping-rain-water.go) | 困难 | ★★★ | 双指针/单调栈 | `go test -run TestTrap` |
| 19 | [删除倒数第 N 个节点](./../../interview/19.remove-nth-node-from-end-of-list.go) | 中等 | ★★★ | 快慢指针/dummy | `go test -run TestRemoveNthFromEnd` |
| 141 | [环形链表](./../../interview/141.linked-list-cycle.go) | 简单 | ★★★ | 快慢指针 | `go test -run TestHasCycle` |
| 142 | [环形链表 II](./../../interview/142.linked-list-cycle-ii.go) | 中等 | ★★★ | 快慢指针/入环点 | `go test -run TestDetectCycle` |
| 160 | [相交链表](./../../interview/160.intersection-of-two-linked-lists.go) | 简单 | ★★☆ | 同步指针 | `go test -run TestGetIntersectionNode` |

## 学习路径建议

```text
第一步：11（理解移动短板为什么不漏答案）
  ↓
第二步：15（排序后用双指针控制重复和方向）
  ↓
第三步：19 / 141 / 142 / 160（链表快慢与同步指针）
  ↓
第四步：42（双指针与单调栈两种证明方式）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./interview/ -run 'Test(MaxArea|ThreeSum|Trap|RemoveNthFromEnd|HasCycle|DetectCycle|GetIntersectionNode)' -count=1
```
