package interview

// 146. LRU 缓存
// HashMap + 双向链表：HashMap 做到 O(1) 查找，链表维护顺序
// 最近使用的放在链表头部，最久未使用的在尾部
// 时间：O(1) | 空间：O(capacity)

type LRUCache struct {
	capacity int
	head     *Node // 虚拟头（最近使用）
	tail     *Node // 虚拟尾（最久未使用）
	cache    map[int]*Node
}

type Node struct {
	key, val int
	prev, next *Node
}

// NewLRUCache 别名，兼容 interview_test.go
func NewLRUCache(capacity int) LRUCache {
	return Constructor146(capacity)
}

func Constructor146(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		cache:    make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		// 移到头部（最近使用）
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		// 已存在：更新值，移到头部
		node.val = value
		this.moveToHead(node)
	} else {
		// 新建节点
		node := &Node{key: key, val: value}
		this.cache[key] = node
		this.addToHead(node)
		// 超过容量，删除尾部节点
		if len(this.cache) > this.capacity {
			tailNode := this.tail.prev
			this.removeNode(tailNode)
			delete(this.cache, tailNode.key)
		}
	}
}

// 将节点移到链表头部（最近使用）
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

// 在头部添加节点
func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 删除节点（从任意位置）
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
