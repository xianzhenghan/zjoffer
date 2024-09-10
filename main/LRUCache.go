package main

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (lru *LRUCache) Get(key int) int {
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		return elem.Value.(*pair).value
	} else {
		return -1
	}
}

func (lru *LRUCache) Put(key int, value int) {
	if elem, ok := lru.cache[key]; ok {
		elem.Value.(*pair).value = value
		lru.list.MoveToFront(elem)
	} else {
		if len(lru.cache) == lru.capacity {
			delete(lru.cache, lru.list.Back().Value.(*pair).key)
			lru.list.Remove(lru.list.Back())
		}
		p := &pair{key, value}
		elem := lru.list.PushFront(p)
		lru.cache[key] = elem
	}
}

func (lru *LRUCache) Put2(key int, value int) {
	if elem, ok := lru.cache[key]; ok {
		elem.Value.(*pair).value = value
		lru.list.MoveToFront(elem)
	} else {
		if len(lru.cache) == lru.capacity {
			delete(lru.cache, lru.list.Back().Value.(*pair).key)
			lru.list.Remove(lru.list.Back())
		}
		p := &pair{key, value}
		elem := lru.list.PushFront(p)
		lru.cache[key] = elem
	}
}
