# 算法面试整理技能

## 触发条件
当用户说"整理算法"、"训练算法"、"出题"时触发，加载这个技能。

---

## 技能定位
面试算法整理与训练双轨并行：
- `interview/` 目录：高频面试题，每题 3 文件（.go / .md / _test.go）
- `algo_template/{type}/` 目录：同类算法通用模板，3 文件（template.go / template.md / template_test.go）

## 目录结构
```
~/go/src/leetcode/
├── interview/           # 高频面试题
│   ├── 33.search-in-rotated-sorted-array.go
│   ├── 33.search-in-rotated-sorted-array.md      ← 含完整引导过程
│   ├── 33.search-in-rotated-sorted-array_test.go
│   └── ...
└── algo_template/      # 算法模板
    ├── binary-search/
    ├── sliding-window/  ← 3文件: sliding_window.go, sliding_window.md, sliding_window_test.go
    ├── backtrack/
    ├── dp/
    ├── heap/
    └── trie/
```

## 文件命名规范
- `.go`：`{number}.{title-lowercase}.go`，package interview，函数名用 LeetCode 原名
- `.md`：`{number}.{title-lowercase}.md`，含引导过程（像教练笔记一步步拆解）
- `_test.go`：`{number}.{title-lowercase}_test.go`，至少 5 个测试用例

## Git Commit 规范
```
feat: {题号}.{中文标题} - 简短描述
feat: 34.在排序数组中查找元素的第一个和最后一个位置 - 两次二分，左闭右开区间
feat: 二分查找速查手册更新 - 补充决策树/旋转数组四姐妹
```
- 每道题的每次修改可独立 commit
- message 里写解题思路/细节

## MD 引导过程格式（核心）
每道题的 .md 应包含：
1. 题目基本信息（难度/标签/示例）
2. **引导过程**：像教练一样一步步拆解思考路径（7步法）
3. 完整可运行代码
4. 至少 2 个例题手动走查
5. 关键点总结表格
6. 复杂度分析

## 已覆盖类型
| 类型 | interview 题号 | 模板 |
|------|--------------|------|
| 滑动窗口 | 3,11,76,239,438,567,643,904,1004,1456 | ✅ |
| 二分查找 | 33,34,35,69,74,81,153,154,162,278,367,704 | ✅ |
| 回溯/DFS | 17,22,39,40,46,47,78,90 | ✅ 完整 |
| 动态规划 | 53,62,63,64,70,96,120 | ✅ 完整 |
| 堆/优先队列 | 23,215,146,295,373 | ❌ 模板未建 |
| Trie | 208,211,212,648 | ❌ 待建 |
| 栈 | 32 | ✅ |
## Subagent 并行注意事项
- subagent 多次运行可能产生重复文件，先 `ls` 确认实际数量
- 大量文件时先摸清现状再并行
- 每次 git commit 前用 `git status` 确认

---

## 二分查找边界详解（常见混淆点）

**混乱根源**：把三件事混在一起
- 区间定义 `[left, right]` vs `[left, right)`
- 循环条件 `left <= right` vs `left < right`
- 收缩方式 `right = mid - 1` vs `right = mid`

**核心对比**

```
闭区间 [l,r]：      while(left <= right)    right = mid - 1, left = mid + 1
左闭右开 [l,r)：    while(left < right)     right = mid,    left = mid + 1

左边界（第一个 >= target）：nums[mid] < target → left = mid + 1
右边界（第一个 > target）：nums[mid] <= target → left = mid + 1
```

**为什么 right = mid 而不是 mid - 1？**
- 左闭右开 `right = mid`：mid 可能是答案，保留
- 闭区间 `right = mid - 1`：mid 已确定不是答案，跳过

**什么时候用哪种？**
- 闭区间 `[l,r]` → 找存在的 target
- 左闭右开 `[l,r)` → 找第一个 >= 或第一个 > 的位置

---

## 两层目录结构

```
interview/              ← 单题训练，每题独立3文件
  34.find-first-and-last-position.go
  34.find-first-and-last-position.md
  34.find-first-and-last-position_test.go

algo_template/xxx/      ← 同类题通用模板/口诀/测试
  xxx.go                 统一模板实现
  xxx.md                 速查手册+记忆口诀
  xxx_test.go            模板自身测试
```

**训练闭环：** interview 高频题 → 归纳到 algo_template 模板 → 训练模式

---

## 输出格式

每道算法整理成 3 个文件：

| 文件 | 内容 |
|------|------|
| `XXX.go` | Go 代码实现，package interview |
| `XXX.md` | 思路/关键点/易错点/记忆方法 |
| `XXX_test.go` | Go 测试用例，至少5个 |

---

## Git Commit 规范

每次修改都可以单独 commit，同一题可分多次：

```
feat: {题号}.{中文标题}
feat: {题号}.{中文标题} - 补充说明（如实现/测试/模板）
```

示例：
```
feat: 34.在排序数组中查找元素的第一个和最后一个位置 - 两次二分，左闭右开区间，左边界lowerBound右边界lowerBound(target+1)-1
feat: 81.搜索旋转排序数组 II - 含重复，当nums[left]==nums[mid]==nums[right]时left++跳过，时间最坏O(n)
```

---

## 工作流程

```
1. 确认用户要求的题目列表
2. 查 interview/ 目录：已有题跳过，缺失题创建
3. 同步更新 algo_template/ 对应类型的模板
4. 每道题创建 3 文件：.go → .md → _test.go
5. go test 验证
6. 分 commit（每题或每类单独 commit）
```

---

## ⚠️ 项目特有经验（Go 文件规范）

**类型定义必须放在 .go 文件里**：不要只在 .md 里定义 `ListNode`，Go 编译器需要类型在同包 `.go` 文件中。类似 `TreeNode`、`Interval` 等共享类型也需要在对应的 `.go` 文件中正确定义。

**测试函数名冲突**：如果 `interview_test.go` 已声明 `TestXXX`，新写的测试文件要改名（如 `TestXXXInterface`）。

**构造函数别名**：如果已有 `ConstructorNNN` 但 `interview_test.go` 用 `NewXXX`，加一个别名函数：
```go
func NewXXX(capacity int) XXX { return ConstructorNNN(capacity) }
```

**Go import 顺序**：`import` 必须紧跟在 `package` 声明之后，不能放在类型定义之后。

---

## ⚠️ 重要经验

**delegate_task 对多文件创建不可靠**：subagent 在创建大量文件时容易中断、超时或重复创建文件（多次写入同一文件导致文件内容异常）。这次整理 binary-search 8道题时 subagent 创建了大量额外文件。

**正确做法**：
- 小批量（1-4道题）：直接用 write_file 工具逐文件创建，稳定可靠
- 大量题库：分批小批量，每批不超过4道，避免 subagent 超时
- 不要一次性分配过多任务给 subagent
- 文件创建后立即 go test 验证

**commit 规范（用户明确要求）**：
```
feat: {题号}.{中文标题}
feat: {题号}.{中文标题} - 补充说明（如实现/测试/模板）
```
同一道题的 .go/.md/_test.go 可以分多次 commit，每次记录一个解题思路或修改点。

---

## 算法文档结构

### 题目信息
```
- 题号：LeetCode XXX
- 标题：题目名称
- 难度：简单/中等/困难
```

### 一句话思路
快速回忆用，一行说明核心方法。

### 核心思路
详细步骤，先说暴力解再说优化。

### 关键点/考点
- 时间/空间复杂度
- 常见追问点

### 易错点
边界条件、陷阱、容易出错的地方。

### 记忆口诀
四句押韵，面试前快速背诵。

### Go 代码实现
完整可运行的代码，包含必要注释。

### 测试用例
覆盖正常、边界、极端情况。

---

## 工作流程

```
1. 用户说"整理算法"，附带题目信息
2. 确认题意（输入/输出/边界）
3. 生成 .md 思路文档
4. 生成 .go 实现代码
5. 生成 _test.go 测试用例
6. 保存到当前目录的 interview/ 子目录
7. 运行测试验证
```

---

## 输出目录
统一保存到 `interview/` 目录下。

---

## 验证
代码完成后必须运行 `go test` 验证测试用例通过。

---

## 已整理的算法类型

| 类型 | 模板状态 | interview 题数 | 备注 |
|------|---------|--------------|------|
| 滑动窗口 | ✅ 完整 | 10题 | 含单调deque/变长窗口/固定窗口 |
| 二分查找 | ✅ 完整 | 12题 | 含旋转数组四姐妹/矩阵映射 |
| 回溯/DFS | ✅ 完整 | 8题 | 17/22/39/40/46/47/78/90 |
| 动态规划 | ✅ 完整 | 7题 | 53/62/63/64/70/96/120 |
| 堆/优先队列 | ⚠️ 部分 | 4题(23,215,146,295?) | 待补：295,373，模板未建 |
| Trie | ❌ 空 | 0题 | 待建：208/211/212/648 |
| 栈 | ✅ 完整 | 1题(32) | 待补：更多 |
| 双指针 | ✅ 完整 | 2题(11,42) | 含装水容器/接雨水可视化 |

训练闭环：interview 高频题 → 归纳到 algo_template 模板 → 引导式训练
训练闭环：interview 高频题 → 归纳到 algo_template 模板 → 引导式训练

---

## 引导式训练方法

**原则：不直接给答案，通过提问引导思考。**

每道题训练流程：
1. 给出题目和简单示例
2. 提出引导问题（层层递进）
3. 给出关键思考题（触及核心难点）
4. 让用户先动手尝试
5. 根据用户反馈决定是继续引导还是给出答案

**引导问题分类：**

| 问题类型 | 典型问法 | 目的 |
|---------|---------|------|
| 复杂度提示 | "看到 O(log n) 要求，你想到什么？" | 确定算法方向 |
| 特征识别 | "这道题的特征是什么？" | 归类到已知模板 |
| 步骤分解 | "第一次二分找什么？" | 拆解复杂问题 |
| 边界确认 | "left/right 初始化成什么？" | 避免边界错误 |
| 反例验证 | "如果数组是 [5,5,5]，你的逻辑对吗？" | 暴露漏洞 |

**关键思考题技巧：**
- 问"怎么找右侧边界"而不是直接说"对 target+1 做 lowerBound"
- 问"为什么是 <= 而不是 <"来触发边界条件讨论
- 问"这个条件和旋转数组的特征有什么关系"来建立联系

**适用题型：** 所有需要"想清楚再写"的题，尤其二分/回溯/DP/滑动窗口

```go
func slidingWindow(s string, t string) int {
    need := make(map[rune]int)
    window := make(map[rune]int)

    // 初始化 need
    for _, c := range t {
        need[c]++
    }

    L, R := 0, 0
    valid := 0  // 已满足条件的字符种类数
    maxLen := 0

    for R < len(s) {
        c := rune(s[R])
        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] { // ★ 必须是 == 不是 >=
                valid++
            }
        }
        R++

        // 收缩窗口
        for valid == len(need) {
            if R-L < maxLen {
                maxLen = R - L
            }
            d := rune(s[L])
            if _, ok := need[d]; ok {
                if window[d] == need[d] { // ★ 必须是 == 不是 >=
                    valid--
                }
                window[d]--
            }
            L++
        }
    }
    return maxLen
}
```

### ★ 关键 Bug：valid++/valid-- 用 == 不是 >=

**原因：** valid 统计的是"有多少种字符已经满足条件"，每种字符只贡献 +1 或 -1。

**错误用法 (>=)：**
```go
if window[c] >= need[c] {  // 错！A 从 1→2 又 +1，重复累加
    valid++
}
```

**正确用法 (==)：**
```go
if window[c] == need[c] {  // 对！只有首次达到 need 时 +1
    valid++
}
```

**适用题目：** 76. 最小覆盖子串、438. 找到所有字母异位词、567. 字符串的排列
