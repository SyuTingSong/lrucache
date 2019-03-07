package lrucache

import (
	"container/list"
	"sync"
)

type pair struct {
	key interface{}
	val interface{}
}

type LRUCache struct {
	Size  int
	cache map[interface{}]*list.Element
	lru   *list.List
	mutex *sync.Mutex
}

func New(size int) *LRUCache {
	return &LRUCache{
		Size:  size,
		cache: make(map[interface{}]*list.Element),
		lru:   list.New(),
		mutex: &sync.Mutex{},
	}
}

func (c *LRUCache) Put(key interface{}, val interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if e, ok := c.cache[key]; ok {
		e.Value = pair{key, val}
		c.lru.MoveToFront(e)
	} else {
		e := c.lru.PushFront(pair{key, val})
		c.cache[key] = e
		if c.lru.Len() > c.Size {
			toDel := c.lru.Back()
			delete(c.cache, toDel.Value.(pair).key)
			c.lru.Remove(toDel)
		}
	}
}

func (c *LRUCache) Get(key interface{}) (result interface{}, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if e, ok := c.cache[key]; ok {
		c.lru.MoveToFront(e)
		return e.Value.(pair).val, true
	}
	return nil, false
}

func (c *LRUCache) Del(key interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if e, ok := c.cache[key]; ok {
		delete(c.cache, key)
		c.lru.Remove(e)
	}
}

func (c *LRUCache) Len() int {
	return c.lru.Len()
}
