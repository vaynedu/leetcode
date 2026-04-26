package trie

// Trie 节点
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool // 标记单词结尾
}

// 新建节点
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// Trie 结构体
type Trie struct {
	root *TrieNode
}

// 初始化
func NewTrie() Trie {
	return Trie{root: NewTrieNode()}
}

// Insert 插入单词 O(m) m=word长度
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search 查找完整单词 O(m)
func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

// StartsWith 查找前缀 O(m)
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

// 内部：查找前缀，返回最后一个节点
func (t *Trie) searchPrefix(prefix string) *TrieNode {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return nil
		}
		node = node.children[ch]
	}
	return node
}
