# Trie · 题目索引

> 学完 `trie.go` 模板后，用这些题实战检验。

## 识别信号

- 多个字符串共享前缀。
- 需要快速判断单词是否存在、前缀是否存在、最短词根、通配符匹配。
- 单个字符串查找可以用 hash，但大量前缀查询会重复扫描。

## 核心不变量

- 从根到某个节点的路径，唯一表示一个前缀。
- `isEnd` 只表示“这里是否是完整单词结尾”，不等于“这里能不能继续往下走”。
- 查询过程中只要某个字符不存在，当前前缀或单词就不存在。

## 决策流程

```text
1. 插入：逐字符创建节点，最后标记 isEnd。
2. 搜索完整单词：逐字符走到末尾，再检查 isEnd。
3. 搜索前缀：逐字符走到末尾即可，不要求 isEnd。
4. 通配符：遇到 '.' 时 DFS 枚举所有 children。
5. 网格 + Trie：DFS 过程中同步走 Trie，走不下去就剪枝。
```

## 模板要点

```
TrieNode:
  children map[rune]*TrieNode   // 子节点
  isEnd    bool                // 是否单词结尾

三个操作：
  Insert(word)   — 逐字符插入，结尾标记isEnd
  Search(word)   — 走到结尾看isEnd
  StartsWith(prefix) — 走到前缀末尾（不关心isEnd）

常见变体：
  计数：额外存count
  前缀查询：存最短词根/所有词
  通配符：'.' 匹配任意字符（DFS搜索children）
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 208 | [实现 Trie](./../../interview/208.implement-trie.go) | 中等 | ★★★ | 基本操作 | `go test -run TestTrie` |
| 211 | [添加与搜索单词](./../../interview/211.design-add-and-search-word.go) | 中等 | ★★★ | Trie+'.'通配符 | `go test -run TestWordDictionary` |
| 212 | [单词搜索 II](./../../interview/212.word-search-ii.go) | 困难 | ★★★ | Trie+DFS网格 | `go test -run TestFindWords` |
| 648 | [单词替换](./../../interview/648.replace-words.go) | 中等 | ★★★ | 最短词根 | `go test -run TestReplaceWords` |

## 学习路径建议

```
第一步：208（完全掌握Trie三个基本操作）
  ↓
第二步：211（加DFS实现'.'通配符）
  ↓
第三步：648（最短词根：插入+查询）
  ↓
第四步：212（最难的Trie+网格DFS，优化：剪枝/截止点）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./algo_template/trie/ -v
```
