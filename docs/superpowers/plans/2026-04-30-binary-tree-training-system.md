# 二叉树训练系统 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 建立二叉树专题第一阶段训练系统：先盘点高价值题，再收口 617 可视化样板，最后沉淀二叉树训练索引。

**Architecture:** 第一阶段只动文档和 617 可视化页面，不改算法实现。新增一份高价值题盘点作为事实清单；用 617 作为可视化样板；更新二叉树训练索引，让后续学习和复习按能力路径推进。

**Tech Stack:** Markdown 文档、Go 测试、单文件 HTML/CSS/JavaScript 可视化、现有 `algo_template/binary-tree` 和 `doc/algorithm-visualizations` 结构。

---

## 文件结构

- Create: `algo_template/binary-tree/HIGH_VALUE_INVENTORY.md`
  - 负责记录常见、高频、有训练价值的二叉树题盘点。
  - 按能力分类，而不是按题号罗列。
  - 标记 Go、测试、笔记、四问结构、可视化、优先级、下一步动作。

- Modify: `doc/algorithm-visualizations/617.merge-two-binary-trees.html`
  - 收口当前未提交修改。
  - 对齐 `53.maximum-subarray.html` 的页面结构和视觉标准。
  - 保留 617 的核心训练重点：同步递归、单空返回非空、原地修改。

- Modify: `algo_template/binary-tree/PROBLEMS.md`
  - 从普通题目列表升级为二叉树训练索引。
  - 链接 `HIGH_VALUE_INVENTORY.md`。
  - 给出学习路径、能力分层、下一题迁移建议。

- Read-only reference: `docs/superpowers/specs/2026-04-30-binary-tree-training-system-design.md`
  - 第一阶段范围说明。

- Read-only reference: `doc/algorithm-visualizations/53.maximum-subarray.html`
  - 可视化风格基准。

- Read-only reference: `doc/algorithm-visualizations/TEMPLATE.md`
  - 可视化页面结构规范。

---

### Task 1: 二叉树高价值题盘点

**Files:**
- Create: `algo_template/binary-tree/HIGH_VALUE_INVENTORY.md`

- [ ] **Step 1: 确认当前二叉树文件状态**

Run:

```bash
rg --files interview doc/algorithm-visualizations algo_template/binary-tree | rg '(binary-tree|path-sum|lowest-common-ancestor|subtree|invert|merge-two-binary-trees|validate-binary-search-tree|kth-smallest|serialize-and-deserialize|construct-binary-tree|diameter)'
```

Expected:

```text
输出包含 94/102/104/105/112/113/114/124/144/145/1644/173/226/230/236/257/297/437/543/572/617/98 等二叉树相关文件。
```

- [ ] **Step 2: 新增高价值题盘点文档**

Create `algo_template/binary-tree/HIGH_VALUE_INVENTORY.md` with this content:

```markdown
# 二叉树高价值题盘点

> 范围：只整理常见、高频、有训练价值的二叉树题。目标不是覆盖所有树题，而是服务面试表达、模式理解、长期练习和必要可视化。

## 盘点口径

- 高频：面试中常见，或是某类二叉树题的代表。
- 高训练价值：能训练可迁移能力，例如遍历选择、递归返回值、路径状态、BST 不变量、结构修改、LCA 判断。
- 优先级：P0 为第一阶段必须打样或复习；P1 为二叉树专题核心题；P2 为后续补强题。

## 总览

| 题号 | 题目 | 分类 | Go | Test | 笔记 | 四问 | 可视化 | 优先级 | 下一步 |
|---|---|---|---|---|---|---|---|---|---|
| 102 | 二叉树层序遍历 | BFS 层序 | 有 | 有 | 有 | 待核对 | 无 | P0 | 作为 BFS 层序样板补训练表达 |
| 104 | 二叉树最大深度 | 深度/返回值递归 | 有 | 有 | 有 | 待核对 | 无 | P1 | 补充“返回高度”的递归不变量 |
| 112 | 路径总和 | 根到叶路径判断 | 有 | 有 | 有 | 已有 | 有 | P0 | 保持为路径 boolean 样板 |
| 113 | 路径总和 II | 路径收集/回溯 | 有 | 有 | 有 | 已有 | 有 | P1 | 强化路径撤销和切片拷贝易错点 |
| 114 | 二叉树展开为链表 | 结构修改/后序 | 有 | 有 | 有 | 已有 | 无 | P0 | 作为结构修改题重点复盘 |
| 124 | 二叉树最大路径和 | 返回值递归/全局答案 | 有 | 有 | 有 | 待核对 | 无 | P1 | 后续作为 Hard 递归样板 |
| 144 | 二叉树前序遍历 | DFS 遍历 | 有 | 有 | 有 | 已有 | 无 | P1 | 保持为前序基础样板 |
| 145 | 二叉树后序遍历 | DFS 遍历 | 有 | 有 | 有 | 已有 | 无 | P1 | 保持为后序基础样板 |
| 1644 | 最近公共祖先 II | LCA 变体 | 有 | 有 | 有 | 待核对 | 无 | P2 | 放入 LCA 进阶训练 |
| 173 | BST 迭代器 | BST 中序/栈 | 有 | 有 | 有 | 待核对 | 无 | P2 | 后续补 BST 迭代训练 |
| 226 | 翻转二叉树 | 结构修改/递归 | 有 | 有 | 有 | 已有 | 无 | P1 | 可作为结构修改入门题 |
| 230 | BST 第 K 小 | BST 中序 | 有 | 有 | 有 | 待核对 | 无 | P1 | 补 BST 有序性表达 |
| 236 | 最近公共祖先 | LCA/后序返回 | 有 | 有 | 有 | 待核对 | 有 | P0 | 作为 LCA 核心题复盘 |
| 257 | 二叉树所有路径 | 根到叶路径收集 | 有 | 有 | 有 | 待核对 | 有 | P1 | 与 112/113 对比训练 |
| 297 | 序列化与反序列化二叉树 | 编码/解码 | 有 | 有 | 有 | 待核对 | 有 | P1 | 后续作为状态机/队列图解复盘 |
| 437 | 路径总和 III | 前缀和/路径计数 | 有 | 有 | 有 | 待核对 | 无 | P1 | 作为路径类进阶题 |
| 543 | 二叉树直径 | 返回值递归/全局答案 | 有 | 有 | 有 | 待核对 | 无 | P1 | 与 124 对比训练 |
| 572 | 另一棵树的子树 | 子树判断/递归匹配 | 有 | 有 | 有 | 待核对 | 无 | P2 | 后续补递归匹配表达 |
| 617 | 合并二叉树 | 同步递归/结构合并 | 有 | 有 | 有 | 已有 | 有 | P0 | 收口为第一阶段可视化样板 |
| 94 | 二叉树中序遍历 | DFS 遍历/BST 基础 | 有 | 有 | 有 | 待核对 | 无 | P1 | 与 BST 题一起复盘 |
| 96 | 不同的二叉搜索树 | BST 计数/DP | 有 | 有 | 有 | 待核对 | 无 | P2 | 放入 BST/DP 交叉专题 |
| 98 | 验证二叉搜索树 | BST 不变量 | 有 | 有 | 有 | 待核对 | 无 | P0 | 作为 BST 不变量样板 |
| 105 | 从前序与中序构造二叉树 | 构造/递归边界 | 有 | 有 | 有 | 待核对 | 无 | P1 | 后续补构造类训练 |

## 第一阶段 P0

| 题号 | 题目 | 训练目标 | 第一阶段动作 |
|---|---|---|---|
| 617 | 合并二叉树 | 两棵树同步递归，单空保留非空子树 | 收口可视化样板 |
| 102 | 二叉树层序遍历 | BFS 按层推进，队列 size 固定当前层 | 后续补 BFS 表达 |
| 112 | 路径总和 | 根到叶 boolean 判断，目标值递减 | 保持路径判断样板 |
| 114 | 二叉树展开为链表 | 后序改结构，避免丢右子树 | 后续复盘结构修改 |
| 236 | 最近公共祖先 | 后序返回命中节点，左右命中则当前为答案 | 后续复盘 LCA 样板 |
| 98 | 验证二叉搜索树 | 每个节点必须满足上下界，而不只是左右孩子 | 后续补 BST 不变量 |

## 训练分组

### 1. 遍历基本功

- 144 前序遍历：根节点先处理，适合自顶向下传状态。
- 94 中序遍历：BST 中天然有序。
- 145 后序遍历：左右子树先处理，适合合并子树信息。
- 102 层序遍历：队列维护下一批节点，按层时先固定当前层大小。

### 2. 路径问题

- 112 路径总和：只需要判断是否存在路径。
- 113 路径总和 II：需要收集路径，注意回溯撤销和拷贝。
- 257 二叉树所有路径：根到叶路径表达。
- 437 路径总和 III：路径不一定从根开始，需要前缀和。

### 3. 返回值递归

- 104 最大深度：返回当前子树高度。
- 543 二叉树直径：返回高度，同时更新经过当前节点的最大直径。
- 124 最大路径和：返回单边最大贡献，同时更新全局最大路径。

### 4. 结构修改

- 226 翻转二叉树：左右子树交换。
- 114 展开为链表：把左右子树改造成右链，注意保存原右子树。
- 617 合并二叉树：同步递归两棵树，单空直接返回非空子树。

### 5. BST 推理

- 98 验证 BST：维护上下界。
- 230 第 K 小：中序遍历天然升序。
- 173 BST 迭代器：用栈模拟受控中序遍历。
- 96 不同 BST：结构计数，属于二叉树与 DP 的交叉。

### 6. 构造、编码、匹配

- 105 构造二叉树：用前序确定根，用中序划分左右子树。
- 297 序列化与反序列化：显式记录 nil，保证结构可还原。
- 572 另一棵树的子树：外层枚举起点，内层判断两棵树是否相同。

### 7. LCA

- 236 最近公共祖先：后序判断左右子树是否分别命中。
- 1644 最近公共祖先 II：需要确认两个目标节点是否都存在。
```

- [ ] **Step 3: 校对盘点文档没有占位**

Run:

```bash
rg -n 'TODO|TBD|待填写|xxx|placeholder' algo_template/binary-tree/HIGH_VALUE_INVENTORY.md
```

Expected:

```text
No matches.
```

- [ ] **Step 4: Commit 盘点文档**

Run:

```bash
git add algo_template/binary-tree/HIGH_VALUE_INVENTORY.md
git commit -m "docs: 盘点二叉树高价值题"
```

Expected:

```text
Commit succeeds with only HIGH_VALUE_INVENTORY.md staged for this task.
```

---

### Task 2: 收口 617 可视化样板

**Files:**
- Modify: `doc/algorithm-visualizations/617.merge-two-binary-trees.html`

- [ ] **Step 1: 阅读当前 617 差异**

Run:

```bash
git diff -- doc/algorithm-visualizations/617.merge-two-binary-trees.html
```

Expected:

```text
Diff shows the existing 617 visualization rewrite. Keep the useful new sections and fix any broken HTML, JS, layout, or wording.
```

- [ ] **Step 2: 对齐页面结构**

Ensure `doc/algorithm-visualizations/617.merge-two-binary-trees.html` contains these sections in order:

```text
1. Hero: LeetCode 617 · Tree DFS + Merge
2. 核心洞察
3. 手推全过程
4. Go 代码
5. 易错点
6. 互动演示
7. Summary
8. Footer
```

Required content points:

```text
- 明确“同步递归两棵树”
- 明确“两个节点都存在时值相加”
- 明确“单边为空时直接返回另一边整棵子树”
- 明确“当前 Go 实现原地修改 root1”
- 易错点必须包含“单空返回 nil 会丢子树”
```

- [ ] **Step 3: 校验 JavaScript 语法**

Run:

```bash
node -e 'const fs=require("fs"); const s=fs.readFileSync("doc/algorithm-visualizations/617.merge-two-binary-trees.html","utf8"); const m=s.match(/<script>([\s\S]*?)<\/script>/); if (m) new Function(m[1]); console.log("js syntax ok")'
```

Expected:

```text
js syntax ok
```

- [ ] **Step 4: 校验 div 结构**

Run:

```bash
python3 -c 'from pathlib import Path
s=Path("doc/algorithm-visualizations/617.merge-two-binary-trees.html").read_text()
inscript=False
depth=0
for line in s.splitlines():
    if "<script" in line:
        inscript=True
    if not inscript:
        depth += line.count("<div") - line.count("</div>")
    if "</script>" in line:
        inscript=False
assert depth == 0, depth
print("div depth ok")'
```

Expected:

```text
div depth ok
```

- [ ] **Step 5: 跑 617 Go 测试**

Run:

```bash
go test ./interview -run TestMergeTrees -count=1 -v
```

Expected:

```text
PASS
ok  	leetcode/interview
```

- [ ] **Step 6: Commit 617 可视化**

Run:

```bash
git add doc/algorithm-visualizations/617.merge-two-binary-trees.html
git commit -m "feat: 617.合并二叉树 - 统一可视化模板"
```

Expected:

```text
Commit succeeds with only 617 visualization staged for this task.
```

---

### Task 3: 沉淀二叉树训练索引

**Files:**
- Modify: `algo_template/binary-tree/PROBLEMS.md`

- [ ] **Step 1: 重写训练索引开头**

Modify the top of `algo_template/binary-tree/PROBLEMS.md` so it starts with:

```markdown
# 二叉树 · 训练索引

> 目标：用常见、高频、有训练价值的题建立二叉树内功。先判断题型，再选择遍历方式、递归返回值和状态维护方式。

相关文档：

- [二叉树高价值题盘点](./HIGH_VALUE_INVENTORY.md)
- [遍历模板](./traversal.go)

## 如何判断二叉树题

- 输入是 `*TreeNode`，或者问题天然可以拆成左子树和右子树。
- 题目问遍历结果、路径、深度、祖先、子树属性、结构修改、BST 性质。
- 父节点答案可以由左右子树答案合并。

## 核心问题

写递归前先回答四个问题：

1. 当前节点要做什么？
2. 左右子树分别要返回什么？
3. 空节点返回什么？
4. 当前节点如何合并左右子树结果？
```

- [ ] **Step 2: 保留并更新能力分组**

Ensure the index contains these sections:

```markdown
## 训练路径

### 第一轮：遍历基本功

| 题号 | 题目 | 训练点 | 建议 |
|---|---|---|---|
| 144 | 二叉树前序遍历 | 自顶向下访问 | 理解前序位置 |
| 94 | 二叉树中序遍历 | BST 有序性基础 | 连接 BST 题 |
| 145 | 二叉树后序遍历 | 子树先处理 | 连接返回值递归 |
| 102 | 二叉树层序遍历 | 队列按层推进 | 固定当前层 size |

### 第二轮：路径问题

| 题号 | 题目 | 训练点 | 建议 |
|---|---|---|---|
| 112 | 路径总和 | boolean 路径判断 | 目标值递减 |
| 113 | 路径总和 II | 路径收集 | 注意回溯撤销和拷贝 |
| 257 | 二叉树所有路径 | 根到叶路径表达 | 对比 112/113 |
| 437 | 路径总和 III | 前缀和计数 | 路径不一定从根开始 |

### 第三轮：结构修改

| 题号 | 题目 | 训练点 | 建议 |
|---|---|---|---|
| 226 | 翻转二叉树 | 交换左右子树 | 结构修改入门 |
| 617 | 合并二叉树 | 同步递归两棵树 | 第一阶段可视化样板 |
| 114 | 二叉树展开为链表 | 后序改结构 | 防止丢失原右子树 |

### 第四轮：BST 与中序

| 题号 | 题目 | 训练点 | 建议 |
|---|---|---|---|
| 98 | 验证二叉搜索树 | 上下界不变量 | 不只比较左右孩子 |
| 230 | BST 第 K 小 | 中序升序 | 计数提前停止 |
| 173 | BST 迭代器 | 栈模拟中序 | 控制遍历节奏 |
| 96 | 不同的二叉搜索树 | 结构计数 | 放入 DP 交叉复习 |

### 第五轮：后序返回值与全局答案

| 题号 | 题目 | 训练点 | 建议 |
|---|---|---|---|
| 104 | 二叉树最大深度 | 返回高度 | 返回值递归入门 |
| 543 | 二叉树直径 | 返回高度，更新直径 | 区分返回值和全局答案 |
| 124 | 二叉树最大路径和 | 返回单边贡献 | Hard 递归样板 |

### 第六轮：构造、编码、LCA

| 题号 | 题目 | 训练点 | 建议 |
|---|---|---|---|
| 105 | 构造二叉树 | 前序定根，中序切分 | 注意区间边界 |
| 297 | 序列化与反序列化 | 显式记录 nil | 保证结构可还原 |
| 236 | 最近公共祖先 | 后序返回命中节点 | LCA 核心题 |
| 1644 | 最近公共祖先 II | 节点存在性判断 | LCA 变体 |
```

- [ ] **Step 3: 增加复习方法**

Append this section near the end of `PROBLEMS.md`:

```markdown
## 每题复习模板

复习一题时不要只看代码，按这个顺序过一遍：

1. 识别信号：为什么这是二叉树题，属于哪个分组？
2. 遍历选择：前序、中序、后序、层序，还是返回值递归？
3. 递归定义：`dfs(root)` 返回什么？
4. 空节点：base case 返回什么？
5. 合并逻辑：当前节点如何用左右子树结果得到答案？
6. 易错点：会不会漏节点、重复算、丢子树、污染路径？
7. 面试表达：能否 1 分钟讲清楚思路和复杂度？

## 第一阶段重点

第一阶段不追求全量升级，先完成：

- `617.merge-two-binary-trees`：可视化样板。
- `HIGH_VALUE_INVENTORY.md`：高价值题盘点。
- 本训练索引：形成二叉树长期练习入口。
```

- [ ] **Step 4: 校对链接和占位**

Run:

```bash
test -f algo_template/binary-tree/HIGH_VALUE_INVENTORY.md
test -f algo_template/binary-tree/traversal.go
rg -n 'TODO|TBD|待填写|xxx|placeholder' algo_template/binary-tree/PROBLEMS.md
```

Expected:

```text
The two test commands exit 0.
The rg command prints no matches.
```

- [ ] **Step 5: Commit 训练索引**

Run:

```bash
git add algo_template/binary-tree/PROBLEMS.md
git commit -m "docs: 完善二叉树训练索引"
```

Expected:

```text
Commit succeeds with only PROBLEMS.md staged for this task.
```

---

### Task 4: 第一阶段收尾验证

**Files:**
- Read: `algo_template/binary-tree/HIGH_VALUE_INVENTORY.md`
- Read: `algo_template/binary-tree/PROBLEMS.md`
- Read: `doc/algorithm-visualizations/617.merge-two-binary-trees.html`

- [ ] **Step 1: 检查 Git 状态**

Run:

```bash
git status --short
```

Expected:

```text
Only unrelated pre-existing changes may remain, such as AGENTS.md. No unstaged changes from Tasks 1-3 should remain.
```

- [ ] **Step 2: 跑二叉树模板测试**

Run:

```bash
go test ./algo_template/binary-tree -count=1 -v
```

Expected:

```text
PASS
ok  	leetcode/algo_template/binary-tree
```

- [ ] **Step 3: 跑二叉树相关面试题聚焦测试**

Run:

```bash
go test ./interview -run 'Test(MergeTrees|LevelOrder|HasPathSum|PathSumII|Flatten|LowestCommonAncestor|IsValidBST|MaxDepth|DiameterOfBinaryTree|MaxPathSum)' -count=1 -v
```

Expected:

```text
PASS
ok  	leetcode/interview
```

- [ ] **Step 4: 汇总结果**

Final response should include:

```text
- 新增的高价值题盘点文件路径
- 收口的 617 可视化文件路径
- 更新的二叉树训练索引路径
- 已运行的验证命令
- 仍然存在但未处理的预先变更，例如 AGENTS.md
```
