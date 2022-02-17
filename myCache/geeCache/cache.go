package geeCache

import (
	"geeCache/lru"
	"sync"
)

// author: songyanhui
// datetime: 2022/1/18 16:37:26
// software: GoLand

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 延迟初始化 该对象的创建会延迟至第一次使用该对象时，提高性能
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.GET(key); ok {
		return v.(ByteView), ok
	}

	return
}
