package leetcode

// 208.实现 Trie (前缀树)

import (
	"testing"
)

/**
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检
查。

 请你实现 Trie 类：


 Trie() 初始化前缀树对象。
 void insert(String word) 向前缀树中插入字符串 word 。
 boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。

 boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，
返回 false 。




 示例：


输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True




 提示：


 1 <= word.length, prefix.length <= 2000
 word 和 prefix 仅由小写英文字母组成
 insert、search 和 startsWith 调用次数 总计 不超过 3 * 10⁴ 次


 👍 1848 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	children [26]*Trie // 子节点数组，对应26个小写字母
	isEnd    bool      // 标记是否为单词结尾
}

func Constructor() Trie {
	return Trie{}
}

// Insert 向前缀树中插入字符串 word
func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		index := ch - 'a' // 计算字符在数组中的索引
		// 如果子节点不存在，则创建新节点
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		// 移动到子节点
		node = node.children[index]
	}
	// 标记单词结尾
	node.isEnd = true
}

// Search 搜索字符串 word 是否在前缀树中
func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		index := ch - 'a'
		// 如果子节点不存在，说明单词不在前缀树中
		if node.children[index] == nil {
			return false
		}
		// 移动到子节点
		node = node.children[index]
	}
	// 检查是否为单词结尾
	return node.isEnd
}

// StartsWith 检查是否存在以 prefix 为前缀的字符串
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		index := ch - 'a'
		// 如果子节点不存在，说明前缀不存在
		if node.children[index] == nil {
			return false
		}
		// 移动到子节点
		node = node.children[index]
	}
	// 只要能遍历完整个前缀，就说明前缀存在
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestImplementTriePrefixTree(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []string
		expected   []interface{}
	}{
		{
			name:       "示例测试",
			operations: []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
			values:     []string{"", "apple", "apple", "app", "app", "app", "app"},
			expected:   []interface{}{nil, nil, true, false, true, nil, true},
		},
		{
			name:       "空字符串测试",
			operations: []string{"Trie", "insert", "search", "startsWith"},
			values:     []string{"", "", "", ""},
			expected:   []interface{}{nil, nil, true, true},
		},
		{
			name:       "重复插入测试",
			operations: []string{"Trie", "insert", "insert", "search", "search"},
			values:     []string{"", "app", "app", "app", "apple"},
			expected:   []interface{}{nil, nil, nil, true, false},
		},
		{
			name:       "前缀与完整单词测试",
			operations: []string{"Trie", "insert", "insert", "search", "startsWith", "search", "startsWith"},
			values:     []string{"", "apple", "application", "app", "app", "apple", "appli"},
			expected:   []interface{}{nil, nil, nil, false, true, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var trie Trie
			for i, op := range tt.operations {
				var result interface{}
				switch op {
				case "Trie":
					trie = Constructor()
					result = nil
				case "insert":
					trie.Insert(tt.values[i])
					result = nil
				case "search":
					result = trie.Search(tt.values[i])
				case "startsWith":
					result = trie.StartsWith(tt.values[i])
				}

				if result != tt.expected[i] {
					t.Errorf("Operation %s with value %s: expected %v, got %v",
						op, tt.values[i], tt.expected[i], result)
				}
			}
		})
	}
}
