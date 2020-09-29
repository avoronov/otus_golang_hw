package hw04_lru_cache //nolint:golint,stylecheck

import "sync" // Key is used for caching - it identify value in cache.
type Key string

// Cache is the interface for cache.
type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mtx      sync.RWMutex
	Capacity int
	Queue    List
	Items    map[Key]*ListItem
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if item, ok := c.Items[key]; ok {
		(item.Value.(*cacheItem)).Value = value
		c.Queue.MoveToFront(item)

		return true
	}

	if c.Queue.Len() == c.Capacity {
		last := c.Queue.Back()
		c.Queue.Remove(last)
		delete(c.Items, (last.Value.(*cacheItem)).Key)
	}

	item := c.Queue.PushFront(&cacheItem{Key: key, Value: value})
	c.Items[key] = item

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	if item, ok := c.Items[key]; ok {
		c.Queue.MoveToFront(item)
		return (item.Value.(*cacheItem)).Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.Queue = NewList()
	c.Items = map[Key]*ListItem{}
}

// NewCache build and return a new instance of lruCache struct.
func NewCache(capacity int) Cache {
	return &lruCache{
		Capacity: capacity,
		Queue:    NewList(),
		Items:    make(map[Key]*ListItem),
	}
}
