---
name: algorithm-visualization
description: Use when creating, updating, reviewing, or debugging LeetCode algorithm visualization HTML pages under doc/algorithm-visualizations, especially when matching the established 53.maximum-subarray.html style, adding interactive demos, or preventing layout/script regressions.
---

# Algorithm Visualization

## Core Rule

Build every `doc/algorithm-visualizations/*.html` page from the project standard:

- Reference style: `doc/algorithm-visualizations/53.maximum-subarray.html`
- Written spec: `doc/algorithm-visualizations/TEMPLATE.md`
- New-page skeleton: `doc/algorithm-visualizations/SKELETON.html`

Do not invent a new layout unless the user explicitly asks.

## Required Page Order

1. Hero: `eyebrow`, `h1`, `subtitle`
2. 核心洞察: `.insight-grid` with three `.insight-card`
3. 手推全过程: `.step` cards with `.step-num`, `.step-range`, `.kv`, `.kv.focal`
4. Go 代码: `.code-wrap`
5. 易错点: `.pitfall-grid` with `.pit-wrong` and `.pit-right`
6. 互动演示: `.anim-section`
7. Summary: exactly three summary cards
8. Footer: current HTML path

## Interactive Demo Contract

Use these names for onclick handlers:

```js
function animReset() {}
function animNext() {}
function animPrev() {}

window.animReset = animReset;
window.animNext = animNext;
window.animPrev = animPrev;

animReset();
```

Rules:

- `animReset()` must render a non-empty initial visual state.
- If functions live inside a closure, still expose them on `window`.
- Do not leave old prefixed handlers like `a3Next` unless the page already requires them and the user asks to preserve them.
- Do not leave `SELF-CHECK`, `sc-result`, temporary debug overlays, or accidental `window.animNext = animNext` text in page copy.

## Visual Constraints

- Reuse the paper palette and typography from `SKELETON.html`.
- Use `.anim-btn`, `.anim-btn-prev`, `.anim-btn-reset`, `.anim-btn-next` for controls.
- Use `.pit-wrong pre, .pit-right pre { white-space: pre-wrap; overflow-wrap: anywhere; }`.
- Keep cards at 8px radius or less.
- Avoid hidden final content: no `display:none` on `.anim-section`, info bars, primary visuals, or timelines.

## Algorithm Content Constraints

- Static trace, Go code, and interactive demo must use consistent state transitions.
- Highlight turning points: repeated character, boundary movement, answer update, pruning, reset, pop/push, or data-structure mutation.
- Explain the invariant in the insight cards and again in the summary.
- Include one concrete wrong case in 易错点.

## Verification

Run these checks after edits:

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

Also inspect for:

- `SELF-CHECK`, `sc-result`, `display:none`
- missing `window.animReset/window.animNext/window.animPrev`
- empty timeline or main visual after reset
- text overflowing pitfall cards
