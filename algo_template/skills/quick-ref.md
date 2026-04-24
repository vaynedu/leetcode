# 算法面试速查卡

> 每道题：1句话思路 + 核心模板 + 记忆口诀 + 追问点

---

## 1. 合并K个升序链表 (23)

**一句话**：分治归并，两两合并直到只剩一个

```go
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 { return nil }
    return mergeSort(lists, 0, len(lists)-1)
}
func mergeSort(lists []*ListNode, l, r int) *ListNode {
    if l == r { return lists[l] }
    mid := l + (r-l)/2  // 防溢出
    return mergeTwoLists(mergeSort(lists,l,mid), mergeSort(lists,mid+1,r))
}
```

**追问**：
- 时间 O(N log K)，空间 O(log K) 递归栈
- 为什么 log K？→ 每次合并减少一半

**口诀**：分治递归两两合，mid防溢要记牢

---

## 2. LRU 缓存 (146)

**一句话**：HashMap O(1) 查 + 双向链表 O(1) 增删

```go
type LRUCache struct {
    capacity int
    cache map[int]*LRUNode
    head, tail *LRUNode  // 哑结点
}
func (c *LRUCache) Get(key int) int {
    if node, ok := c.cache[key]; ok {
        c.moveToHead(node); return node.Val
    }
    return -1
}
func (c *LRUCache) Put(key, val int) {
    // 存在则更新移表头，不存在则新建，容量超则删tail
}
```

**追问**：
- 为什么双向链表？→ 单向需找前驱 O(n)
- 淘汰谁？→ tail 前一个

**口诀**：HashMap来查找，双向链表快增删

---

## 3. 合并区间 (56)

**一句话**：排序后相邻比较，重叠取 max

```go
func merge(intervals []Interval) []Interval {
    sort(intervals)  // 按Start排序
    result := []Interval{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        if intervals[i].Start <= result[-1].End {
            result[-1].End = max(result[-1].End, intervals[i].End)
        } else {
            result = append(result, intervals[i])
        }
    }
    return result
}
```

**追问**：
- 时间 O(N log N)，空间 O(N)
- 重叠条件？→ curr.Start ≤ last.End
- 合并方式？→ End 取 max

**口诀**：先排起始，相邻比较，重叠取end最大

---

## 4. 岛屿数量 (200)

**一句话**：DFS 感染，把相邻1全变0

```go
func numIslands(grid [][]byte) int {
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                count++
                dfs(i, j)
            }
        }
    }
    return count
}
func dfs(i, j int) {
    if 越界 || grid[i][j]=='0' { return }
    grid[i][j] = '0'  // 感染
    dfs(i+1,j); dfs(i-1,j); dfs(i,j+1); dfs(i,j-1)
}
```

**追问**：
- 时间/空间 O(MN)
- 为什么标记0？→ 避免重复计数
- 递归深度？→ 可能栈溢出

**口诀**：遍历找岛屿，遇一就计数，DFS感染四邻，标记零来去重

---

## 5. 二叉树层序遍历 (102)

**一句话**：队列 + 记录当前层长度

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil { return nil }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)  // ★ 关键：当前层节点数
        level := []int{}
        for i := 0; i < size; i++ {
            node := queue[0]; queue = queue[1:]
            level = append(level, node.Val)
            if node.Left != nil { queue = append(queue, node.Left) }
            if node.Right != nil { queue = append(queue, node.Right) }
        }
        result = append(result, level)
    }
}
```

**追问**：
- 时间 O(N)，空间 O(N) 队列
- 如何分层？→ size = len(queue)
- BFS vs DFS？→ BFS 层次，DFS 深度

**口诀**：层序遍历用队列，记录长度分层次

---

## 速记总表

| 题号 | 数据结构 | 核心技巧 | 时间 |
|-----|---------|---------|-----|
| 23 | 链表 | 分治归并 | O(N log K) |
| 146 | HashMap + 双向链表 | 表头表尾操作 | O(1) |
| 56 | 数组/排序 | 相邻比较 | O(N log N) |
| 200 | 网格/DFS | 感染标记 | O(MN) |
| 102 | 队列 | size分层 | O(N) |

---

## 面试 Checklist

```
☐ 确认题意（输入输出边界）
☐ 说思路（先暴力后优化）
☐ 分析复杂度
☐ 写代码（处理空指针）
☐ 主动提测试用例
```
