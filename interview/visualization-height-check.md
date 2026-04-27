# Visualization Height Check

> 快速自检清单 — 防止图表柱条/水柱超出容器高度

## 核心原则
所有柱状图/条形图必须使用**统一比例尺**，禁止子元素高度超出容器高度。

## 比例尺公式
```js
const maxVal = Math.max(...dataArray);
const SCALE = 20;    // px per unit（统一比例尺）
const containerHeight = 200;  // px
const barHeight = (value / maxVal) * (containerHeight * 0.8);
// 确保 barHeight <= containerHeight
```

## 自检清单（每次提交前必检）

1. 找到所有 `height: Npx` 的容器
2. 找到容器内所有 `height: Mpx` 的柱子/bar/water-fill
3. 验证：max(子元素高度) ≤ 容器高度

## 常见错误模式

```html
<!-- 错误：容器 80px，柱子 200px -->
<div style="...height:80px">
  <div style="...height:200px">8</div>  <!-- 200 > 80 ✗ -->
</div>

<!-- 正确：容器 200px，柱子 160px（8/8*160=160px） -->
<div style="...height:200px">
  <div style="...height:160px">8</div>  <!-- 160 <= 200 ✓ -->
</div>
```

## 示例：LeetCode 11 装水容器

- 数据：[1, 8, 6, 2, 5, 4, 8, 3, 7]，最大值 8
- 容器：200px
- 比例尺：20px/unit（8 × 20 = 160px ≤ 200px ✓）
- 各柱高度：h=1→20px, h=8→160px, h=6→120px, ...

## 验证命令
```bash
grep -n 'height:[0-9]*px' file.html
grep -A5 'class="step"' file.html | grep 'height:'
```

## 相关 Skills
- `algorithm-visualization-format`：完整格式标准（暖色排版、5 section、JS 规范）
- `visualization-height-check`（hermes skills）：详细自检清单与修复方法
