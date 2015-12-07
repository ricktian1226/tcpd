package xycache

//常用的双缓存基类定义

import (
	"sync/atomic"
)

type CacheBase struct {
	index int32
}

func (c *CacheBase) Major() int32 {
	return atomic.LoadInt32(&c.index)
}

func (c *CacheBase) Secondary() int32 {
	return (atomic.LoadInt32(&c.index) + 1) % 2
}

func (c *CacheBase) Switch() {
	atomic.StoreInt32(&c.index, (atomic.LoadInt32(&c.index)+1)%2)
}
