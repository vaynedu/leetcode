# 算法面试整理技能

## 触发条件
当用户说"整理算法"时触发，加载这个技能。

---

## 技能定位
当需要整理算法题时，按照标准 3 文件格式生成完整内容。

---

## 输出格式

每道算法整理成 3 个文件：

| 文件 | 内容 |
|------|------|
| `XXX.go` | Go 代码实现 |
| `XXX.md` | 思路/关键点/易错点/记忆方法 |
| `XXX_test.go` | Go 测试用例 |

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

## 滑动窗口模板（高频模板）

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
