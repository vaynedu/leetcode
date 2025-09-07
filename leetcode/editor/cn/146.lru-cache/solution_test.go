package leetcode

// 146.LRU 缓存

import (
	"fmt"
	"testing"
)

/**

 请你设计并实现一个满足
 LRU (最近最少使用) 缓存 约束的数据结构。



 实现
 LRUCache 类：





 LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-
value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。




 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



 示例：


输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4




 提示：


 1 <= capacity <= 3000
 0 <= key <= 10000
 0 <= value <= 10⁵
 最多调用 2 * 10⁵ 次 get 和 put


 👍 3584 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)

// 定义双向链表节点
type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

// LRUCache LRU缓存结构体
type LRUCache struct {
	Capacity int
	Cache    map[int]*Node // 哈希表，存储key到节点的映射
	Head     *Node         // 虚拟头节点
	Tail     *Node         // 虚拟尾节点
}

// Constructor 构造函数
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*Node),
		Head:     &Node{}, // 创建虚拟头节点
		Tail:     &Node{}, // 创建虚拟尾节点
	}

	// 初始化双向链表
	lru.Head.Next = lru.Tail
	lru.Tail.Prev = lru.Head

	return lru
}

// Get 获取key对应的值
func (this *LRUCache) Get(key int) int {
	// 如果key不存在，返回-1
	if node, exists := this.Cache[key]; exists {
		// 将节点移到链表头部（标记为最新使用）
		this.moveToHead(node)
		return node.Value
	}
	return -1
}

// Put 插入或更新key-value对
func (this *LRUCache) Put(key int, value int) {
	// 如果key已存在
	if node, exists := this.Cache[key]; exists {
		// 更新值
		node.Value = value
		// 移到链表头部
		this.moveToHead(node)
		return
	}

	// 如果key不存在，创建新节点
	newNode := &Node{
		Key:   key,
		Value: value,
	}

	// 如果缓存已满，删除最久未使用的节点（尾部节点）
	if len(this.Cache) >= this.Capacity {
		// 删除尾部节点
		tail := this.removeTail()
		// 从哈希表中删除
		delete(this.Cache, tail.Key)
	}

	// 添加新节点到头部
	this.addToHead(newNode)
	// 在哈希表中添加映射
	this.Cache[key] = newNode
}

// addToHead 将节点添加到链表头部
func (this *LRUCache) addToHead(node *Node) {
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
}

// removeNode 从链表中删除指定节点
// 在双向链表中删除节点O(1)
func (this *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// moveToHead 将节点移到链表头部
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

// removeTail 删除链表尾部节点并返回
func (this *LRUCache) removeTail() *Node {
	lastNode := this.Tail.Prev
	this.removeNode(lastNode)
	return lastNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestLruCache(t *testing.T) {
	// 测试示例
	lru := Constructor(2)

	lru.Put(1, 1) // 缓存是 {1=1}
	lru.Put(2, 2) // 缓存是 {1=1, 2=2}

	if lru.Get(1) != 1 { // 返回 1
		t.Errorf("Expected 1, got %d", lru.Get(1))
	}

	lru.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}

	if lru.Get(2) != -1 { // 返回 -1 (未找到)
		t.Errorf("Expected -1, got %d", lru.Get(2))
	}

	lru.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}

	if lru.Get(1) != -1 { // 返回 -1 (未找到)
		t.Errorf("Expected -1, got %d", lru.Get(1))
	}

	if lru.Get(3) != 3 { // 返回 3
		t.Errorf("Expected 3, got %d", lru.Get(3))
	}

	if lru.Get(4) != 4 { // 返回 4
		t.Errorf("Expected 4, got %d", lru.Get(4))
	}

	fmt.Println("LRU Cache测试通过!")
}
