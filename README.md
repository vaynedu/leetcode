# LeetCode 算法题解库

面试算法速记 · Go 实现 · 可视化图解

---

## 目录结构

```
leetcode/
├── interview/                    # 面试高频题（核心）
│   ├── *.go                     # 题解代码
│   ├── *_test.go               # 测试用例
│   └── *.md                     # 题解笔记
├── leetcode/editor/cn/          # 原始 LeetCode 题目存档
├── algo_template/               # 算法模板（学习核心）
│   ├── sliding-window/          # 滑动窗口模板 + PROBLEMS.md
│   ├── binary-search/           # 二分查找模板 + PROBLEMS.md
│   ├── two-pointers/            # 双指针模式总纲
│   ├── backtrack/               # 回溯模板 + PROBLEMS.md
│   ├── dp/                      # 动态规划模板 + PROBLEMS.md
│   ├── trie/                    # 前缀树模板 + PROBLEMS.md
│   ├── binary-tree/             # 二叉树模板 + PROBLEMS.md
│   ├── queue/                   # 队列/BFS 模板 + PROBLEMS.md
│   ├── stack/                   # 栈/单调栈模式总纲
│   ├── heap/                    # 堆/优先队列模式总纲
│   ├── graph/                   # 图/BFS/拓扑排序模式总纲
│   ├── tree/                    # 树遍历模板 + PROBLEMS.md
│   ├── array/                   # 数组技巧
│   ├── greedy/                 # 贪心
│   └── skills/                  # 算法面试整理旧技能与辅助脚本
├── doc/
│   └── algorithm-visualizations/ # 可视化图解 HTML
│       ├── index.html           # 可视化看板（入口）
│       ├── TEMPLATE.md          # 图解页面规范
│       ├── SKELETON.html        # 新图解起始骨架
│       ├── 53.maximum-subarray.html # 当前风格基准
│       └── ...                 # 60+ 个图解
├── skills/                      # 本仓库可版本化的 Codex 技能
│   └── algorithm-visualization/ # 算法图示生成/维护技能
└── model/                       # 数据结构模型（辅助）
```

---

## 仓库统计

| 类别 | 数量 |
|------|------|
| 面试高频题 | 100 道 |
| 可视化图解 | 61 个 |
| 算法模板 | 14 个方向 |

---

## 学习路径

### 第一步：选一个模板开始

每个模板目录下有 `PROBLEMS.md`，列出该模板对应的所有面试题，按难度排序。
如果要按面试优先级训练，先看 `interview/高频题清单.md`。

```
滑动窗口 → 8 题   → doc: 11/76/239/3/438/567/643/904
二分查找  → 8 题   → doc: 33/34/35/69/74/81/153/154
回溯      → 8 题   → doc: 17/22/39/40/46/47/78/90
动态规划  → 7 题   → doc: 53/62/63/64/70/96/120
前缀树    → 4 题   → doc: 208/211/212/648
二叉树    → 11 题  → doc: 112/113/257/236/617/...
堆/队列   → 4 题   → doc: 23/215/295/373
```

### 第二步：学完模板，马上实战

```bash
# 1. 看模板源码，理解核心逻辑
cat algo_template/sliding-window/sliding_window.go

# 2. 打开 PROBLEMS.md，找到对应题目
cat algo_template/sliding-window/PROBLEMS.md

# 3. 写题，对答案
vim interview/76.minimum-window-substring.go
go test ./interview/ -run TestMinimumWindow -v

# 4. 思路卡住？看可视化图解
open doc/algorithm-visualizations/76.minimum-window-substring.html
```

### 第三步：用看板总览全局

```bash
open doc/algorithm-visualizations/index.html
```

功能：
- **分类过滤**：二叉树 / 双指针 / 滑动窗口 / 二分 / DP / 堆 / 栈 / 其他
- **搜索**：按题号、标题、标签过滤
- **难度 + 频率标签**：🔥🔥🔥 = 面试常考，🔥 = 偶尔考
- **点击卡片**：直接在新窗口打开可视化图解

---

## 使用方法

### 运行测试

```bash
# 所有面试题测试
go test ./interview/ -count=1

# 单题测试
go test ./interview/ -run TestSlidingWindowMaximum -v

# 单模板测试
go test ./algo_template/sliding-window/ -v
```

### 添加新题

1. 在 `interview/` 创建三个文件：
   ```bash
   999.new-problem.go          # 题解
   999.new-problem_test.go     # 测试
   999.new-problem.md          # 题解笔记
   ```
2. 在看板 `doc/algorithm-visualizations/index.html` 的 `cards` 数组加一条
3. 在对应模板的 `PROBLEMS.md` 加一行

### 添加可视化图解

1. 从 `doc/algorithm-visualizations/SKELETON.html` 复制新文件：
   ```bash
   cp doc/algorithm-visualizations/SKELETON.html doc/algorithm-visualizations/XXX.problem-name.html
   ```
2. 按 `doc/algorithm-visualizations/TEMPLATE.md` 填充页面：
   - Hero
   - 核心洞察
   - 手推全过程
   - Go 代码
   - 易错点
   - 互动演示
   - Summary
   - Footer
3. 风格以 `doc/algorithm-visualizations/53.maximum-subarray.html` 为基准。
4. 在看板 `doc/algorithm-visualizations/index.html` 的 `cards` 数组注册。
5. 修改后验证 JS 和 HTML 结构：
   ```bash
   node -e 'const fs=require("fs"); const s=fs.readFileSync("doc/algorithm-visualizations/XXX.problem-name.html","utf8"); const m=s.match(/<script>([\s\S]*?)<\/script>/); if (m) new Function(m[1]); console.log("js syntax ok")'
   ```
   ```bash
   python3 -c 'from pathlib import Path
s=Path("doc/algorithm-visualizations/XXX.problem-name.html").read_text()
inscript=False; depth=0
for line in s.splitlines():
    if "<script" in line: inscript=True
    if not inscript:
        depth += line.count("<div") - line.count("</div>")
    if "</script>" in line: inscript=False
assert depth == 0, depth
print("div depth ok")'
   ```

### Codex 技能

本仓库把可复用的 agent 工作流放在 `skills/`：

```
skills/
└── algorithm-visualization/
    └── SKILL.md
```

`algorithm-visualization` 用于创建、调整、调试 `doc/algorithm-visualizations/*.html`，会强制沿用 `53.maximum-subarray.html` 的页面风格和 `TEMPLATE.md` / `SKELETON.html` 的结构约定。

本机 Codex 自动发现用的安装副本位于：

```
~/.codex/skills/algorithm-visualization/SKILL.md
```

### Commit 规范

修改具体算法题时，按“一题一 commit”组织，方便回溯和面试前速记。

推荐格式：

```bash
feat: {题号}.{题目名称} - {本次修改核心}
```

示例：

```bash
feat: 3.无重复最长子串 - 统一可视化模板并修复窗口 trace
feat: 53.最大子数组和 - 移除自检浮层并修复易错点布局
```

如果是跨题型的基础设施修改，比如 README、模板、技能、看板，可单独提交：

```bash
docs: 完善算法模式总览
docs: 更新算法可视化模板规范
```

---

## 题目分类速查

### 双指针
| # | 题目 | 难度 | 频率 |
|---|------|------|------|
| 11 | 装水容器 | 中等 | 🔥🔥🔥 |
| 42 | 接雨水 | 困难 | 🔥🔥🔥 |

### 滑动窗口
| # | 题目 | 难度 | 频率 |
|---|------|------|------|
| 3  | 无重复最长子串 | 中等 | 🔥🔥🔥 |
| 76 | 最小覆盖子串 | 困难 | 🔥🔥🔥 |
| 239 | 滑动窗口最大值 | 困难 | 🔥🔥🔥 |
| 438 | 字母异位词 | 中等 | 🔥🔥🔥 |
| 567 | 字符串排列 | 中等 | 🔥🔥🔥 |

### 二分查找
| # | 题目 | 难度 | 频率 |
|---|------|------|------|
| 33 | 旋转数组搜索 | 中等 | 🔥🔥🔥 |
| 34 | 二分查找边界 | 困难 | 🔥🔥🔥 |
| 81 | 旋转数组 II | 中等 | 🔥🔥 |
| 153 | 旋转数组最小值 | 中等 | 🔥🔥🔥 |

### 回溯
| # | 题目 | 难度 | 频率 |
|---|------|------|------|
| 17 | 电话号码字母组合 | 中等 | 🔥🔥 |
| 22 | 括号生成 | 中等 | 🔥🔥🔥 |
| 39 | 组合总和 | 中等 | 🔥🔥🔥 |
| 46 | 全排列 | 中等 | 🔥🔥🔥 |
| 78 | 子集 | 中等 | 🔥🔥🔥 |

### 动态规划
| # | 题目 | 难度 | 频率 |
|---|------|------|------|
| 53 | 最大子数组和 | 简单 | 🔥🔥🔥 |
| 62 | 不同路径 | 中等 | 🔥🔥🔥 |
| 70 | 爬楼梯 | 简单 | 🔥🔥🔥 |
| 96 | 不同的二叉搜索树 | 中等 | 🔥🔥🔥 |

### 堆 / 优先队列
| # | 题目 | 难度 | 频率 |
|---|------|------|------|
| 23 | 合并 K 个链表 | 困难 | 🔥🔥🔥 |
| 215 | 数组第 K 大 | 中等 | 🔥🔥🔥 |
| 295 | 数据流中位数 | 困难 | 🔥🔥🔥 |

### 二叉树
| # | 题目 | 难度 | 频率 |
|---|------|------|------|
| 112 | 路径总和 | 简单 | 🔥🔥🔥 |
| 236 | 最近公共祖先 | 中等 | 🔥🔥🔥 |
| 437 | 路径总和 III | 中等 | 🔥🔥🔥 |

### 其他高频
| # | 题目 | 难度 | 频率 |
|---|------|------|------|
| 146 | LRU 缓存 | 困难 | 🔥🔥🔥 |
| 32 | 最长有效括号 | 困难 | 🔥🔥🔥 |

---

## 设计理念

1. **先模板，后刷题**：学会一个算法模型，再去找对应的题实战，而不是零散刷题
2. **看板驱动**：用可视化看板总览全局，按难度/频率过滤，知道先刷什么
3. **图解辅助记忆**：每个核心算法都有纸笔风格的图解，帮助 3 秒回忆思路
4. **测试即验证**：每道题都有 `_test.go`，`go test -v` 即可验证答案
