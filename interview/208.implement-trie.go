package interview

/*
208. 实现 Trie (前缀树)

Trie() 初始化前缀树对象
void insert(String word) 向前缀树插入字符串 word
boolean search(String word) 返回字符串 word 是否已经存在
boolean startsWith(String prefix) 返回是否存在前缀 prefix

时间：O(m) 每操作 | 空间：O(m * n) m=长度，n=字符集
*/

// TrieNode 前缀树节点
type TrieNode208 struct {
	children map[rune]*TrieNode208
	isEnd    bool
}

func NewTrieNode208() *TrieNode208 {
	return &TrieNode208{children: make(map[rune]*TrieNode208)}
}

type Trie208 struct {
	root *TrieNode208
}

func NewTrie208() Trie208 {
	return Trie208{root: NewTrieNode208()}
}

// Insert 插入
func (t *Trie208) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode208()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search 查找完整单词
func (t *Trie208) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

// StartsWith 查找前缀
func (t *Trie208) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

func (t *Trie208) searchPrefix(prefix string) *TrieNode208 {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return nil
		}
		node = node.children[ch]
	}
	return node
}
