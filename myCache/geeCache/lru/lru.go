package lru

import (
	"container/list"
	"fmt"
)

// author: songyanhui
// datetime: 2022/1/5 18:30:33
// software: GoLand

type Cache struct {
	maxBytes int64
	nBytes   int64
	ll       *list.List
	cache    map[string]*list.Element
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) GET(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		// 移动ele元素到 ll列表的前面，若ele不存在ll里面不做变更，ele不能为空
		c.ll.MoveToFront(ele)
		fmt.Println(ele.Value)
		kv := ele.Value.(*entry)
		fmt.Println(kv)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	// 取ll队列中的最后一个元素，若ll为空返回nil
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		// 更新当前所用的内存
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		// 回调函数，若不为 nil则调用
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add adds a value to the cache.
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	// 超出最大容量，淘汰数据
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}


