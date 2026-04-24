package interview

// ================================================================
// 146. LRU 缓存
// ================================================================

// LRUNode 双向链表节点
type LRUNode struct {
	Key  int
	Val  int
	Prev *LRUNode
	Next *LRUNode
}

// LRUCache O(1) get/put - HashMap + 双向链表
type LRUCache struct {
	capacity int
	cache    map[int]*LRUNode // key -> Node
	head     *LRUNode         // 表头（最近使用）
	tail     *LRUNode         // 表尾（最久未使用）
}

// NewLRUCache 构造函数
func NewLRUCache(capacity int) *LRUCache {
	h := &LRUNode{}
	t := &LRUNode{}
	h.Next = t
	t.Prev = h
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*LRUNode),
		head:     h,
		tail:     t,
	}
}

// Get 获取值
func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToHead(node)
		return node.Val
	}
	return -1
}

// Put 存入值
func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.cache[key]; ok {
		node.Val = value
		c.moveToHead(node)
	} else {
		node := &LRUNode{Key: key, Val: value}
		c.cache[key] = node
		c.addToHead(node)
		if len(c.cache) > c.capacity {
			removed := c.removeTail()
			delete(c.cache, removed.Key)
		}
	}
}

func (c *LRUCache) addToHead(node *LRUNode) {
	node.Prev = c.head
	node.Next = c.head.Next
	c.head.Next.Prev = node
	c.head.Next = node
}

func (c *LRUCache) removeNode(node *LRUNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (c *LRUCache) moveToHead(node *LRUNode) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeTail() *LRUNode {
	node := c.tail.Prev
	c.removeNode(node)
	return node
}
