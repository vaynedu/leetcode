package interview

/*
212. 单词搜索 II

给定 m x n 二维字符网格 board 和字符串列表 words
返回所有二维网格中存在的单词

思路：Trie + DFS 回溯
1. 将所有 words 插入 Trie
2. 从每个格子开始 DFS Trie 路径
3. 搜到 isEnd 的节点 → 收集单词

时间：O(m * n * 4^L) L=单词最大长度 | 空间：O(w * L)
*/

// TrieNode212 前缀树节点
type TrieNode212 struct {
	children map[rune]*TrieNode212
	isEnd    bool
	word     string // 记录完整单词（isEnd 时填充）
}

func NewTrieNode212() *TrieNode212 {
	return &TrieNode212{children: make(map[rune]*TrieNode212)}
}

type Trie212 struct {
	root *TrieNode212
}

func NewTrie212() Trie212 {
	return Trie212{root: NewTrieNode212()}
}

func (t *Trie212) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode212()
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.word = word
}

func (t *Trie212) SearchPrefix(prefix string) *TrieNode212 {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

// FindWords 在 board 上搜索所有存在于 Trie 中的单词
func FindWords(board [][]byte, words []string) []string {
	trie := NewTrie212()
	for _, w := range words {
		trie.Insert(w)
	}
	m, n := len(board), len(board[0])
	var result []string

	// DFS：从 (i,j) 出发，尝试匹配 Trie 路径
	var dfs func(i, j int, node *TrieNode212)
	dfs = func(i, j int, node *TrieNode212) {
		ch := rune(board[i][j])
		next, ok := node.children[ch]
		if !ok {
			return
		}

		// 如果是单词结尾，收集（去重：找到后标记 isEnd=false）
		if next.isEnd {
			result = append(result, next.word)
			next.isEnd = false // 防止从其他路径重复找到
		}

		// 标记已访问，四个方向 DFS
		board[i][j] = '#' // 用 '#' 标记已访问
		dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && board[ni][nj] != '#' {
				dfs(ni, nj, next)
			}
		}
		board[i][j] = byte(ch) // 恢复
	}

	// 从每个格子启动 DFS
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, trie.root)
		}
	}

	return result
}
