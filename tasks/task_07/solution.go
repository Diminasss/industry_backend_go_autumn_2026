package main

import (
	"container/list"
	"sync"
)

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	mu       sync.Mutex
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	if c.capacity <= 0 {
		return value, false
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	element, ok := c.items[key]
	if !ok {
		return value, false
	}

	item := element.Value.(entry[K, V])
	c.ll.MoveToFront(element)

	return item.value, true
}
func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if element, ok := c.items[key]; ok {
		item := element.Value.(entry[K, V])
		item.value = value
		element.Value = item
		return
	}

	if len(c.items) == c.capacity {
		oldest := c.ll.Back()
		if oldest == nil {
			return
		}

		item := oldest.Value.(entry[K, V])
		delete(c.items, item.key)
		c.ll.Remove(oldest)
	}
	c.items[key] = c.ll.PushFront(entry[K, V]{key: key, value: value})
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
