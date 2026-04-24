# 算法可视化 HTML 规范

## 目的

为难以理解的算法生成图文并茂的 HTML 图解，放在 `doc/algorithm-visualizations/` 目录下。

## 设计风格

**配色方案（纸笔风格）**

```css
--paper: #faf7f2;           /* 米白底色 */
--paper-2: #f2ede4;        /* 卡片背景 */
--ink: #1c1917;             /* 深色文字/边框 */
--muted: #57534e;          /* 次要文字 */
--soft: #78716c;            /* 淡化文字 */
--rule: rgba(28,25,23,0.12); /* 分割线 */
--accent: #b5523a;          /* rust 强调色 */
--accent-tint: rgba(181,82,58,0.08);
--accent-strong: rgba(181,82,58,0.18);
```

**字体**

```html
<link href="https://fonts.googleapis.com/css2?family=Instrument+Serif:ital@0;1&family=Geist:wght@400;500;600&family=Geist+Mono:wght@400;500;600&display=swap" rel="stylesheet">
```

| 用途 | 字体 |
|------|------|
| 标题 | Instrument Serif |
| 正文 | Geist |
| 代码/数字 | Geist Mono |

## 页面结构

每个 HTML 分为 3-4 个 section：

```
Section 1: 概念引入（用图示解释基本概念）
Section 2: 核心洞察（为什么能work，关键insight可视化）
Section 3: 完整Trace（一步步的手推过程，高亮当前焦点）
Section 4+: 补充（如有多种情况，用case-panel对比）
```

底部放 summary cards（三张）和 footer。

## 通用 CSS 组件

### 数组格子

```html
<div class="array-row">
  <div class="cell sm">4</div>
  <div class="cell sm mid">6</div>
  <div class="cell sm eliminated">7</div>
</div>
```

| class | 效果 |
|-------|------|
| `.cell` | 56×56px，默认白底黑框 |
| `.cell.sm` | 40×40px，用于小格 |
| `.cell.mid` | 脉冲动画，强调当前mid |
| `.cell.hit` | 找到答案，accent背景 |
| `.cell.eliminated` | 划线淡化，表示被排除 |
| `.cell.sorted-half` | accent-tint背景，表示有序半 |

### 指针行

```html
<div class="pointer-row">
  <div class="ptr-slot L"><span class="arr">↑</span><span class="label">L</span></div>
  <div class="ptr-slot M"><span class="arr">↑</span><span class="label">M</span></div>
  <div class="ptr-slot R"><span class="arr">↑</span><span class="label">R</span></div>
</div>
```

### 断崖标记

```html
<div class="pivot sm"></div>
```

中间插入一条 2px 竖线，表示旋转断点，配合 `--accent` 色。

### Step 卡片

```html
<div class="step">
  <div class="step-head">
    <span class="step-num">Step 1</span>
    <span class="step-range">L=0 &nbsp; M=3 &nbsp; R=6</span>
  </div>
  <!-- array-row + pointer-row -->
  <div class="decision">
    <div class="kv">条件判断...</div>
    <div class="kv focal">焦点结论...</div>
  </div>
</div>
```

### Case 对比面板

```html
<div class="cases">
  <div class="case-panel">
    <p class="case-title">情况 ①</p>
    <p class="case-cond">条件</p>
    <!-- 图示 -->
    <p class="conclusion">结论</p>
  </div>
  <div class="case-panel">...</div>
</div>
```

### Summary Cards

```html
<div class="summary">
  <div class="card">
    <p class="card-eyebrow">Step 1</p>
    <h3><span class="dot accent"></span>标题</h3>
    <p>说明文字</p>
  </div>
  <div class="card">...</div>
  <div class="card">...</div>
</div>
```

## 文件命名

```
doc/algorithm-visualizations/{题号}.{英文标题}.html
```

例如：
```
33.rotated-array-search.html
11.container-with-most-water.html
```

## 创建流程

1. 理解算法核心难点在哪里
2. 设计 3-4 个可视化 section
3. 从参考 HTML 复制 CSS 部分，保留配色方案
4. 写入 `doc/algorithm-visualizations/{题号}.{标题}.html`
5. 本地打开验证效果
6. git commit: `docs: 算法可视化 - {题号} {标题}`

## 动画效果（可选）

```css
/* 脉冲：mid 格子 */
@keyframes pulse {
  0%, 100% { box-shadow: 0 0 0 0 rgba(181,82,58,0.4); }
  50% { box-shadow: 0 0 0 8px rgba(181,82,58,0); }
}
.cell.mid { animation: pulse 1.8s ease-in-out infinite; }

/* 找到答案：庆祝动画 */
@keyframes celebrate {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.08); }
}
.cell.hit { animation: celebrate 1.8s ease-in-out infinite; }

/* 入场动画 */
section { animation: fadeIn 0.6s ease-out backwards; }
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to { opacity: 1; transform: none; }
}
```

## 参考文件

- `doc/algorithm-visualizations/33.rotated-array-search.html` — 旋转排序数组完整示例
