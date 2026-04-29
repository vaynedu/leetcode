# 算法可视化 HTML 模板规范

## 基准

后续 `doc/algorithm-visualizations/*.html` 统一参考 `53.maximum-subarray.html` 的信息架构和视觉语言，但不要复制其中历史调试残留。新页面优先从 `doc/algorithm-visualizations/SKELETON.html` 开始。

目标是做“能讲题、能手推、能互动”的单文件 HTML：

- 页面打开即可阅读，不依赖构建工具。
- 静态讲解和互动动画状态一致。
- 视觉风格统一，方便连续浏览多道题。
- 调试代码不得留在最终页面。

## 文件命名

```text
doc/algorithm-visualizations/{题号}.{英文标题}.html
```

示例：

```text
doc/algorithm-visualizations/53.maximum-subarray.html
doc/algorithm-visualizations/3.longest-substring-without-repeating-characters.html
```

## 页面结构

每个页面按以下顺序组织。

### 1. Hero

```html
<p class="eyebrow">LeetCode 53 · Greedy + DP</p>
<h1>最大子数组和</h1>
<p class="subtitle">一句话说明核心直觉</p>
```

要求：

- `eyebrow` 写题号和算法类别。
- `h1` 写中文题名。
- `subtitle` 用一句话说明核心算法，不写长段背景。

### 2. 核心洞察

用 3 张 `.insight-card` 说明：

- 观察：题目里最关键的结构。
- 选择：为什么这个算法能 work。
- 复杂度或关键变量：帮助快速记忆。

### 3. 手推全过程

用 `.step` 卡片做完整 trace。每步至少包含：

```html
<div class="step">
  <div class="step-head">
    <span class="step-num">i=3</span>
    <span class="step-range">关键变量状态</span>
  </div>
  <div class="kv">普通推导</div>
  <div class="kv focal">关键结论</div>
</div>
```

要求：

- trace 不必覆盖所有机械重复步骤，但必须覆盖转折点、边界点和答案出现点。
- `.step-range` 中的变量必须和正文推导、互动演示一致。
- 不要用纯文字代替状态变化，能画数组/窗口/栈/树时优先画出来。

### 4. Go 代码

使用 `.code-wrap`。代码应是可直接搬到 LeetCode 的版本。

要求：

- 少量注释解释关键分支即可。
- 不在代码块里塞长篇说明。
- 行过长时自然换行，避免横向溢出。

### 5. 易错点

使用左右对比：

```html
<div class="pitfall-grid">
  <div class="pit-wrong">错误写法</div>
  <div class="pit-right">正确写法</div>
</div>
```

要求：

- `.pit-wrong pre, .pit-right pre` 必须设置 `white-space: pre-wrap; overflow-wrap: anywhere;`。
- 右侧正确代码不要写超长单行注释，长解释放到代码块下方的 `<p>`。
- 错误案例必须具体到输入和错误结果。

### 6. 互动演示

互动演示使用 `.anim-section`，控件顺序固定：

```html
<button onclick="animPrev()">上一步</button>
<button onclick="animReset()">重置</button>
<button onclick="animNext()">下一步</button>
```

必须提供：

- info bar：当前关键变量。
- phase text：当前步骤解释。
- main visual：数组、窗口、栈、树、图或其他主视觉。
- controls：上一步、重置、下一步。

脚本要求：

```js
function animReset() {}
function animNext() {}
function animPrev() {}

window.animReset = animReset;
window.animNext = animNext;
window.animPrev = animPrev;

animReset();
```

注意：

- 如果函数写在 IIFE 内，也必须显式暴露到 `window`。
- 不要把 `window.animNext = animNext` 等暴露语句插入页面文案。
- `animReset()` 必须渲染初始视觉状态，不能留下空白主区域或空 timeline。
- 不要在最终页面保留 `SELF-CHECK` 浮层、自检 IIFE、临时 console 调试代码。

### 7. Summary

固定 3 张 summary card：

- 核心算法
- 正确性直觉
- 易错关键点

### 8. Footer

footer 显示当前 HTML 路径，便于浏览器里确认文件。

## 统一视觉

### 配色

```css
:root {
  --paper: #faf7f2;
  --paper-2: #f2ede4;
  --ink: #1c1917;
  --muted: #57534e;
  --soft: #78716c;
  --rule: rgba(28,25,23,0.12);
  --accent: #b5523a;
  --accent-tint: rgba(181,82,58,0.08);
  --accent-strong: rgba(181,82,58,0.18);
}
```

### 字体

```html
<link href="https://fonts.googleapis.com/css2?family=Instrument+Serif:ital@0;1&family=Geist:wght@400;500;600&family=Geist+Mono:wght@400;500;600&display=swap" rel="stylesheet">
```

| 用途 | 字体 |
|------|------|
| 标题 | Instrument Serif |
| 正文 | Geist |
| 代码/数字 | Geist Mono |

### 布局

- `.container` 最大宽度 960px。
- `section` 和 `.anim-section` 使用 8px 圆角。
- 卡片圆角不超过 8px。
- 移动端不允许文字溢出，`pre` 必须可换行或可横向滚动。

## 常用组件

### Insight Cards

```html
<div class="insight-grid">
  <div class="insight-card">
    <p class="card-eyebrow">观察</p>
    <h3><span class="dot"></span>标题</h3>
    <p>说明文字</p>
  </div>
</div>
```

### Step Cards

```html
<div class="step">
  <div class="step-head">
    <span class="step-num">Step 1</span>
    <span class="step-range">变量状态</span>
  </div>
  <div class="kv">推导</div>
  <div class="kv focal">结论</div>
</div>
```

### 数组格子

```html
<div class="nums-display">
  <div class="num-cell n-default">4</div>
  <div class="num-cell n-cur">-1</div>
  <div class="num-cell n-best">2</div>
</div>
```

状态建议：

- `.n-default`：普通元素。
- `.n-cur`：当前处理元素。
- `.n-window`：当前窗口或候选区间。
- `.n-best`：当前最优答案区间。

### 代码块

```html
<pre class="code-wrap"><span class="fn">func</span> <span class="hl">solve</span>() {
    <span class="cmt">// key idea</span>
}</pre>
```

## 创建流程

1. 明确算法核心难点：状态转移、边界、剪枝、窗口收缩、数据结构不变量等。
2. 从 `SKELETON.html` 复制为目标文件。
3. 填写 Hero、核心洞察、手推全过程、Go 代码、易错点、互动演示、Summary。
4. 确保静态 trace 与 JS 动画使用同一组样例和状态。
5. 运行验证命令。
6. 浏览器打开页面，至少点击：下一步到末尾、上一步、重置。

## 验证清单

创建或修改页面后，至少检查：

```bash
node -e 'const fs=require("fs"); const s=fs.readFileSync("doc/algorithm-visualizations/FILE.html","utf8"); const m=s.match(/<script>([\s\S]*?)<\/script>/); if (m) new Function(m[1]); console.log("js syntax ok")'
```

```bash
python3 -c 'from pathlib import Path
s=Path("doc/algorithm-visualizations/FILE.html").read_text()
inscript=False; depth=0
for line in s.splitlines():
    if "<script" in line: inscript=True
    if not inscript:
        depth += line.count("<div") - line.count("</div>")
    if "</script>" in line: inscript=False
assert depth == 0, depth
print("div depth ok")'
```

人工检查：

- 页面没有 `SELF-CHECK`、`FAIL`、临时调试浮层。
- `onclick` 使用的函数都在 `window` 上。
- `animReset()` 后主视觉不空白。
- `display:none` 没有隐藏互动区。
- 易错点代码框不越界。
- 静态推导、代码、互动演示三者答案一致。

## 参考文件

- `doc/algorithm-visualizations/53.maximum-subarray.html`：当前推荐风格基准。
- `doc/algorithm-visualizations/SKELETON.html`：新题起始模板。
