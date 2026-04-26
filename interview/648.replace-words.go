package interview

/*
648. 单词替换

在句子中，将每个单词替换成词根（如果在字典中有词根的话）
词典：dictionary 中任意词根是某个单词的前缀

输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the baby"
输出："the cat was rat by the baby"

思路：Trie
1. 将所有词根插入 Trie
2. 遍历 sentence 单词，在 Trie 中查找最短词根

时间：O(m * n) m=单词数，n=平均长度 | 空间：O(m * n)
*/

// TrieNode648 前缀树节点
type TrieNode648 struct {
	children map[rune]*TrieNode648
	isEnd    bool
}

func NewTrieNode648() *TrieNode648 {
	return &TrieNode648{children: make(map[rune]*TrieNode648)}
}

type Trie648 struct {
	root *TrieNode648
}

func NewTrie648() Trie648 {
	return Trie648{root: NewTrieNode648()}
}

// Insert 插入词根
func (t *Trie648) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode648()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// FindRoot 查找单词的最短词根，返回原始单词（如果不存在）
func (t *Trie648) FindRoot(word string) string {
	node := t.root
	var rootBytes []byte
	for _, ch := range word {
		if _, ok := node.children[rune(ch)]; !ok {
			// 无此路径 → 返回原单词
			return word
		}
		rootBytes = append(rootBytes, byte(ch))
		node = node.children[rune(ch)]
		if node.isEnd {
			// 找到词根，立即返回
			return string(rootBytes)
		}
	}
	return word
}

// ReplaceWords 将 sentence 中的单词替换成词根
func ReplaceWords(dictionary []string, sentence string) string {
	trie := NewTrie648()
	for _, w := range dictionary {
		trie.Insert(w)
	}

	// 切分 sentence 单词
	start := 0
	var result []byte
	for i := 0; i <= len(sentence); i++ {
		if i == len(sentence) || sentence[i] == ' ' {
			word := sentence[start:i]
			result = append(result, trie.FindRoot(word)...)
			if i < len(sentence) {
				result = append(result, ' ')
			}
			start = i + 1
		}
	}
	return string(result)
}
