package main

type Cache[K comparable, V any] struct {
	capacity int
	items    map[K]V
}

// NewCache Якобы конструктор
func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	var cache Cache[K, V]
	cache.capacity = capacity
	cache.items = make(map[K]V, capacity)
	return &cache
}

// Set и Get Якобы методы
func (c *Cache[K, V]) Set(k K, v V) bool {
	_, ok := c.items[k]

	if ok {
		c.items[k] = v
		return true
	}

	if len(c.items) >= c.capacity {
		return false
	}

	c.items[k] = v
	return true
}

func (c *Cache[K, V]) Get(k K) (v V, ok bool) {
	value, ok := c.items[k]
	return value, ok
}
