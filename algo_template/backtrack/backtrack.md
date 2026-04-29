# 回溯算法（Backtracking）速查手册

## 识别信号

- 要求列出所有组合、排列、子集、路径、分割方案。
- 题目有“可以选择/不选择”“顺序是否重要”“不能重复使用”等决策约束。
- 暴力枚举看似可行，但需要剪枝或去重。

## 核心不变量

- `path` 始终表示当前递归路径上的合法前缀。
- 递归层表示当前决策位置。
- 做选择后进入下一层，返回前必须撤销选择，恢复现场。
- 去重只跳过同一层的重复选择，不跳过不同路径中的合法重复值。

## 决策流程

```text
1. 定义 path 和结果集。
2. 明确递归参数：start、idx、remain、used、left/right count。
3. 写终止条件：何时收集答案，何时剪枝返回。
4. 遍历当前层选择列表。
5. 做选择 → 递归 → 撤销选择。
6. 如果有重复元素，先排序，再判断同层去重条件。
```

## 为什么不会漏答案

递归树覆盖了每个合法决策序列；剪枝只剪掉已经不可能成为合法答案的分支；撤销选择保证兄弟分支从相同状态开始。

## 一、核心思想

回溯 = 枚举 + 撤销选择

```
选择 → 递归 → 撤销
```

通过**选/不选**或**选这个还是那个**的决策树遍历，找出所有满足条件的解。

---

## 二、回溯解题checklist

```
1. 明确问题类型
   ├── 子集？      → 每个元素选/不选
   ├── 组合？      → 从n个选k个
   ├── 排列？      → 所有元素顺序
   ├── 分割？      → 按条件切分
   ├── 棋盘？      → 二维网格+路径
   └── 括号？      → 左/右括号计数

2. 确定数据结构
   ├── path []int       → 当前路径
   ├── used []bool      → 元素是否用过
   ├── cols, diag1, diag2 → N皇后攻击位置
   └── start int        → 下一层起始位置

3. 确定终止条件
   ├── len(path) == k        → 选够k个
   ├── idx == len(nums)      → 遍历完
   └── len(path) == n        → 放满棋盘

4. 剪枝策略
   ├── 排序去重（跳过相同值）
   ├── 前后相同跳过（同层）
   ├── 越界跳过
   └── 攻击位置跳过
```

---

## 三、分类模板

### 类型1: 子集（Subset）

**特征**：每个元素选或不选

```
nums = [1,2,3]
选[] → [1] → [1,2] → [1,2,3]
          → [1,3]
     → [2] → [2,3]
     → [3]

共 2^n 个子集
```

```go
var dfs func(idx int)
dfs = func(idx int) {
    res = append(res, path)  // 收集所有节点
    
    for i := idx; i < len(nums); i++ {
        path = append(path, nums[i])
        dfs(i + 1)
        path = path[:len(path)-1]  // 撤销
    }
}
```

**经典题**：

| 题号 | 名称 | 核心 |
|-----|------|------|
| 78 | 子集 | 每个元素选/不选 |
| 90 | 子集 II | 排序+剪枝去重 |

---

### 类型2: 组合（Combination）

**特征**：从 n 个选 k 个，不讲究顺序

```
n=4, k=2
[1,2], [1,3], [1,4]
[2,3], [2,4]
[3,4]
```

```go
var dfs func(start int)
dfs = func(start int) {
    if len(path) == k {
        res = append(res, copy(path))
        return
    }
    
    // 剪枝：剩余不足
    for i := start; i <= n-(k-len(path))+1; i++ {
        path = append(path, i)
        dfs(i + 1)
        path = path[:len(path)-1]
    }
}
```

**经典题**：

| 题号 | 名称 | 核心 |
|-----|------|------|
| 77 | 组合 | C(n,k) |
| 39 | 组合总和 | 可重复选，无重复 |
| 40 | 组合总和 II | 不可重复，有重复 |
| 216 | 组合总和 III | k个数和为n |

---

### 类型3: 排列（Permutation）

**特征**：所有元素都要排，讲究顺序

```
[1,2,3]
[1,2,3], [1,3,2]
[2,1,3], [2,3,1]
[3,1,2], [3,2,1]

共 n! 个排列
```

```go
used := make([]bool, len(nums))
var dfs func()
dfs = func() {
    if len(path) == len(nums) {
        res = append(res, copy(path))
        return
    }
    
    for i := 0; i < len(nums); i++ {
        if used[i] { continue }
        
        path = append(path, nums[i])
        used[i] = true
        dfs()
        used[i] = false
        path = path[:len(path)-1]
    }
}
```

**经典题**：

| 题号 | 名称 | 核心 |
|-----|------|------|
| 46 | 全排列 | 无重复 |
| 47 | 全排列 II | 有重复，排序+剪枝 |
| 784 | 字母大小写全排列 | 数字不变，字母大小写 |

---

### 类型4: 括号生成

**特征**：左括号随时可加，右括号必须等左括号够了

```
n=2
()()
(())
n=3
((()))
(()())
(())()
()(())
()()()
```

```go
var dfs func(left, right int)
dfs = func(left, right int) {
    if left == n && right == n {
        res = append(res, string(path))
        return
    }
    
    // 可以加左括号
    if left < n {
        path = append(path, '(')
        dfs(left+1, right)
        path = path[:len(path)-1]
    }
    
    // 可以加右括号（right < left）
    if right < left {
        path = append(path, ')')
        dfs(left, right+1)
        path = path[:len(path)-1]
    }
}
```

**经典题**：

| 题号 | 名称 | 核心 |
|-----|------|------|
| 22 | 括号生成 | n对括号 |

---

### 类型5: 棋盘（N-Queens）

**特征**：二维网格放置，按行递归

```
N=4 棋盘
. Q . .
. . . Q
Q . . .
. . Q .

攻击位置：
- 同列 cols[col]
- 主对角线 row-col 相同
- 副对角线 row+col 相同
```

```go
var dfs func(row int, cols, diag1, diag2 map[int]bool)
dfs = func(row int, cols, diag1, diag2 map[int]bool) {
    if row == n {
        res = append(res, boardToString(path))
        return
    }
    
    for col := 0; col < n; col++ {
        d1 := row - col + n  // 主对角线
        d2 := row + col      // 副对角线
        
        if cols[col] || diag1[d1] || diag2[d2] {
            continue
        }
        
        path[row][col] = 'Q'
        cols[col], diag1[d1], diag2[d2] = true, true, true
        
        dfs(row+1, cols, diag1, diag2)
        
        path[row][col] = '.'
        cols[col], diag1[d1], diag2[d2] = false, false, false
    }
}
```

**经典题**：

| 题号 | 名称 | 核心 |
|-----|------|------|
| 51 | N皇后 | 返回所有解 |
| 52 | N皇后 II | 返回解的数量 |
| 37 | 解数独 | 3×3宫+行列 |
| 79 | 单词搜索 | 上下左右搜索 |

---

### 类型6: 分割（Palindrome/Split）

**特征**：把字符串切成若干段，每段满足条件

```
s = "aab"
[a,a,b]
[aa,b]
```

```go
var dfs func(idx int)
dfs = func(idx int) {
    if idx == len(s) {
        res = append(res, copy(path))
        return
    }
    
    for i := idx; i < len(s); i++ {
        if isPalindrome(idx, i) {
            path = append(path, s[idx:i+1])
            dfs(i + 1)
            path = path[:len(path)-1]
        }
    }
}
```

**经典题**：

| 题号 | 名称 | 核心 |
|-----|------|------|
| 131 | 分割回文串 | 每段回文 |
| 93 | 复原IP地址 | 每段0-255 |

---

### 类型7: 路径/搜索

**特征**：二维网格中搜索路径

```go
dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}

var dfs func(i, j, idx int) bool
dfs = func(i, j, idx int) bool {
    if idx == len(word) {
        return true
    }
    
    if i<0 || i>=m || j<0 || j>=n || visited[i][j] || board[i][j]!=word[idx] {
        return false
    }
    
    visited[i][j] = true
    for _, d := range dirs {
        if dfs(i+d[0], j+d[1], idx+1) {
            visited[i][j] = false
            return true
        }
    }
    visited[i][j] = false
    return false
}
```

**经典题**：

| 题号 | 名称 | 核心 |
|-----|------|------|
| 79 | 单词搜索 | 四方向 |
| 130 | 被围绕的区域 | 边界O感染 |

---

## 四、一图总结

```
┌────────────────────────────────────────────────────────────┐
│                      回溯算法分类                            │
├──────────────┬───────────────┬─────────────────────────────┤
│   类型        │   特征        │   关键代码                     │
├──────────────┼───────────────┼─────────────────────────────┤
│ 子集          │ 选/不选       │ for i:=start; i<n; i++      │
│ 组合          │ 选k个         │ len(path)==k 终止            │
│ 排列          │ 讲究顺序      │ used[] 标记已选              │
│ 括号          │ 左<右         │ right<left 才能放右         │
│ N皇后         │ 行列斜攻击    │ cols,diag1,diag2 剪枝       │
│ 分割          │ 切分回文      │ isPalindrome 检查            │
│ 路径搜索      │ 二维网格      │ dirs 四方向                  │
└──────────────┴───────────────┴─────────────────────────────┘
```

---

## 五、易错点汇总

| 题目 | 易错点 |
|------|--------|
| 子集II/排列II | 不排序直接去重会漏，要先排序 |
| 组合总和 | 每个元素可重复用 → dfs(i)，不可重复 → dfs(i+1) |
| N皇后 | row-col 和 row+col 可能为负，要加n偏移 |
| 复原IP | 前导0不合法（"01"不行，"0"行）|
| 单词搜索 | visited 要在递归返回前撤销 |

---

## 六、通用代码框架

```go
func backtrack(res *[][]int, path []int, choices []int) {
    // 1. 终止条件
    if len(path) == k {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *res = append(*res, tmp)
        return
    }
    
    // 2. 遍历选择列表
    for i := start; i < len(choices); i++ {
        // 3. 做选择
        path = append(path, choices[i])
        
        // 4. 递归
        backtrack(res, path, i+1)
        
        // 5. 撤销选择
        path = path[:len(path)-1]
    }
}
```

---

## 七、复杂度分析

| 问题类型 | 时间复杂度 | 空间复杂度 |
|---------|-----------|-----------|
| 子集 | O(2^n) | O(n) 递归栈 |
| 组合 | O(C(n,k)) | O(k) |
| 排列 | O(n!) | O(n) |
| N皇后 | O(N!) | O(N) |
