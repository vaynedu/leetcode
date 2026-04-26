package interview

/*
211. 添加与搜索单词 - 数据结构设计

WordDictionary() 初始化数据结构
void addWord(word) 添加一个单词
boolean search(word) 搜索字面单词或正则表达式
  - 正则只包含 '.'，'.' 可以匹配任意单个字符

时间：addWord O(m) | search O(m) 或 O(26^m) 含'.'
空间：O(m * n)
*/

// TrieNode211 前缀树节点（支持 '.' 通配符）
type TrieNode211 struct {
	children map[rune]*TrieNode211
	isEnd    bool
}

func NewTrieNode211() *TrieNode211 {
	return &TrieNode211{children: make(map[rune]*TrieNode211)}
}

type WordDictionary struct {
	root *TrieNode211
}

func NewWordDictionary() WordDictionary {
	return WordDictionary{root: NewTrieNode211()}
}

// AddWord 添加单词
func (w *WordDictionary) AddWord(word string) {
	node := w.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode211()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search 搜索（支持 '.' 通配符）
func (w *WordDictionary) Search(word string) bool {
	return w.searchHelper(word, w.root)
}

// 递归搜索：word 是剩余要匹配的字符串
func (w *WordDictionary) searchHelper(word string, node *TrieNode211) bool {
	for i, ch := range word {
		if ch == '.' {
			// 遍历所有子节点，匹配剩余部分
			for _, child := range node.children {
				if w.searchHelper(word[i+1:], child) {
					return true
				}
			}
			return false // '.' 但没有子节点
		}
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}
