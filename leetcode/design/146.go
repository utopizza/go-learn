package design

type LRUCache struct {
	cap  int
	size int
	head *Node
	tail *Node
	hash map[int]*Node
}

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:  capacity,
		size: 0,
		hash: make(map[int]*Node),
		head: nil,
		tail: nil,
	}
}

func (c *LRUCache) Get(key int) int {
	if node, exist := c.hash[key]; exist {
		c.moveToHead(node)
		return node.Val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	// key已存在，更新value，移动到队头
	if node, exist := c.hash[key]; exist {
		node.Val = value
		c.moveToHead(node)
		return
	}
	// key未存在，插入队头
	node := &Node{Key: key, Val: value}
	c.insertHead(node)
	// 如果容量超过限制，删除tail节点
	if c.size > c.cap {
		c.deleteTail()
	}
}

func (c *LRUCache) insertHead(node *Node) {
	if c.head == nil {
		c.head, c.tail = node, node
	} else {
		node.Pre = nil
		node.Next = c.head
		c.head.Pre = node
		c.head = node
	}
	c.hash[node.Key] = node
	c.size++
}

func (c *LRUCache) deleteTail() {
	if c.tail == nil {
		return
	}
	delete(c.hash, c.tail.Key)
	pre := c.tail.Pre
	if pre != nil {
		pre.Next = nil
	}
	c.tail.Pre = nil
	c.tail = pre
	c.size--
}

func (c *LRUCache) moveToHead(node *Node) {
	if c.size <= 1 {
		return
	}
	if node == c.head {
		return
	}
	if node == c.tail {
		c.tail = node.Pre
		c.tail.Next = nil
		node.Pre = nil
		node.Next = c.head
		c.head.Pre = node
		c.head = node
		return
	}
	pre, next := node.Pre, node.Next
	pre.Next = next
	next.Pre = pre
	node.Next = c.head
	node.Pre = nil
	c.head.Pre = node
	c.head = node
}
