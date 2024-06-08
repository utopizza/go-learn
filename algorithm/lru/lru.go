package lru

import "fmt"

type Cache struct {
	cap  int32
	list *DoubleLinkList
	keys map[string]*Node
}

type DoubleLinkList struct {
	size int32
	head *Node
	tail *Node
}

type Node struct {
	key  string
	val  interface{}
	last *Node
	next *Node
}

func NewCache(cap int32) *Cache {
	return &Cache{
		cap:  cap,
		list: &DoubleLinkList{},
		keys: make(map[string]*Node),
	}
}

func (c *Cache) Get(key string) interface{} {
	defer c.print()

	if node, exist := c.keys[key]; exist {
		c.deleteNode(node)
		c.insertHead(node)
		return node.val
	}

	return nil
}

func (c *Cache) Put(key string, val interface{}) {
	defer c.print()

	// 先看key是否已存在，若已存在则直接更新val，并提到队首
	if node, exist := c.keys[key]; exist {
		node.val = val
		c.deleteNode(node)
		c.insertHead(node)
		return
	}

	// 新key入队头
	c.insertHead(&Node{
		key: key,
		val: val,
	})

	// 如果cache已满，逐出队尾key
	if c.list.size > c.cap {
		c.deleteNode(c.list.tail)
	}
}

func (c *Cache) insertHead(toInsertNode *Node) {
	if c.list.size == 0 {
		c.list.head, c.list.tail = toInsertNode, toInsertNode
	} else {
		oldHead := c.list.head
		oldHead.last = toInsertNode
		toInsertNode.next = oldHead
		c.list.head = toInsertNode
	}
	c.keys[toInsertNode.key] = toInsertNode
	c.list.size++
}

func (c *Cache) deleteNode(toDeleteNode *Node) {
	// 处理前后节点
	last, next := toDeleteNode.last, toDeleteNode.next
	if last != nil {
		last.next = next
	}
	if next != nil {
		next.last = last
	}

	// 处理head、tail指针
	if c.list.head == toDeleteNode {
		c.list.head = nil
	}
	if c.list.tail == toDeleteNode {
		c.list.tail = toDeleteNode.last
	}

	// 从map中删除
	delete(c.keys, toDeleteNode.key)
	c.list.size--
}

func (c *Cache) print() {
	var keys []string
	for p := c.list.head; p != nil; p = p.next {
		keys = append(keys, p.key)
	}
	fmt.Println(keys)
}
