package main

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/TreeNode/lru.go
 * @Description:
 * @datetime: 2022/6/23 15:32:46
 * software: GoLand
**/

type LRUCache struct {
	// 目前size 大小
	size int
	// 容量大小
	capacity int
	cache    map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

func initDlinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: map[int]*DLinkedNode{},
		head: initDlinkedNode(0, 0),
		tail: initDlinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.next = l.head
	return l
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok {
		node := initDlinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++

		if this.size > this.capacity {
			removeNode := this.removeTail()
			delete(this.cache, removeNode.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}

}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
