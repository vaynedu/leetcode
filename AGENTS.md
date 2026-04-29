<claude-mem-context>
# Memory Context

# [leetcode] recent context, 2026-04-29 2:02pm GMT+8

Legend: 🎯session 🔴bugfix 🟣feature 🔄refactor ✅change 🔵discovery ⚖️decision 🚨security_alert 🔐security_note
Format: ID TIME TYPE TITLE
Fetch details: get_observations([IDs]) | Search: mem-search skill

Stats: 42 obs (10,511t read) | 2,298,772t work | 100% savings

### Apr 28, 2026
101 4:44p 🔵 Codex CLI 配置：使用 paperhub 自定义 provider，模型为 gpt-5.5
102 " 🔵 ~/.codex 目录结构：含 history、日志、SQLite 数据库及 plugins 缓存
103 4:53p 🔵 LeetCode 3 算法可视化 HTML 文件结构
106 4:58p 🔵 HTML 静态说明区存在索引值笔误
108 5:06p 🔴 修复 LeetCode 3 可视化 HTML 多处文字错误
116 5:08p ⚖️ 算法可视化统一模板方向确定
117 5:32p 🔵 算法可视化目录结构与 53.maximum-subarray.html 模板结构探明
118 " 🔵 53.maximum-subarray.html 内嵌 JS 自检 IIFE 模式
119 " ⚖️ 53.maximum-subarray.html 发现两个 bug，并决定升级 TEMPLATE.md
120 5:52p 🔵 53.maximum-subarray.html JS 自检失败根因定位
121 5:53p 🔴 53.maximum-subarray.html — JS self-check IIFEs removed and window assignments consolidated
122 " 🔵 53.maximum-subarray.html patches from prior session did not persist
123 5:57p 🔴 53.maximum-subarray.html — all patches applied successfully, validation passes
124 " 🟣 TEMPLATE.md rewritten and SKELETON.html created for algorithm visualization standard
125 6:59p 🔴 53.maximum-subarray.html — final validation confirmed clean across all checks
126 7:01p ✅ algorithm-visualizations work complete — 4 files changed, uncommitted
### Apr 29, 2026
127 10:29a ⚖️ Session pivots to align 3.longest-substring-without-repeating-characters.html with 53 template standard
128 " 🔵 3.html Diverges from Template Standard in 8 Areas
129 " ⚖️ Repository Knowledge Structure Redesign with 4-Question Framework
130 10:53a 🟣 3.html Fully Rewritten to Match Template Standard
131 " 🟣 algorithm-visualization Skill Created and Installed
132 " ✅ README.md Updated with Accurate Stats and skills/ Directory
133 " 🔵 Algorithm Template Docs Inventory: Patterns With and Without .md Files
134 10:56a ⚖️ Git Commit 规范：一题一 commit，格式含题号+核心速记
135 11:04a 🟣 6 个算法模式文档统一补充识别信号、核心不变量、决策流程、为什么不会漏答案四大板块
136 " ✅ 新建 two-pointers、stack、heap、graph 四个模式目录
137 " 🟣 新增双指针、栈/单调栈、堆、图四个模式的完整文档
138 " ✅ README.md 更新：算法模板方向数从 10 增至 14，目录树新增 4 个模式
139 " 🔵 greedy.md、queue/PROBLEMS.md、array/array.md 仍为旧表格格式，缺少识别信号/核心不变量
140 11:06a ⚖️ 新任务：逐题补写四问结构笔记并单独 commit
141 11:07a 🟣 题 3 和 53 的 .md 升级为四问结构
142 11:20a 🟣 8 道第一优先级高频题全部升级为四问结构笔记
143 " 🟣 8 高频题 .md 升级为四问结构
144 " ⚖️ 每道题单独一个 commit 的提交规范
145 11:42a 🟣 高频题四问复盘文档逐题提交进行中
146 " 🟣 commit: 239.滑动窗口最大值 四问复盘
147 11:58a 🔵 仓库当前状态全景：8 道高频题已完成，大量基础设施变更待提交
149 " 🟣 算法模式题目索引全覆盖
150 " 🟣 高频题复习索引 interview/高频题清单.md
151 " 🔵 仓库待提交变更清单（截至 c950613）
148 12:01p 🔵 仓库结构审计：四问复盘覆盖率 8/102，模式文档存在结构性缺口
152 12:30p 🟣 第二批 6 题四问复盘结构补齐（438/567/904/81/153/154）

Access 2299k tokens of past work via get_observations([IDs]) or mem-search skill.
</claude-mem-context>

<INSTRUCTIONS>
## Commit 规范

- 修改一个具体算法题时，默认一个题目一个 commit，便于回溯和速记。
- commit message 格式优先使用：`feat: {题号}.{题目名称} - {本次修改核心}`。
- 核心内容要一眼看出改了什么：算法思路、边界修复、可视化调整、测试补充等。
- 示例：`feat: 3.无重复最长子串 - 统一可视化模板并修复窗口 trace`
- 示例：`feat: 53.最大子数组和 - 移除自检浮层并修复易错点布局`
- 跨题型的模板、README、skill 等基础设施修改可单独 commit，message 写清范围，例如：`docs: 完善算法模式总览`。
</INSTRUCTIONS>
