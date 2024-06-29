package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

var mu = sync.Mutex{}

func (lru *lruCache) Set(key Key, value interface{}) bool {
	mu.Lock()
	defer mu.Unlock()

	v, ok := lru.items[key]

	if ok {
		v.Value = value
		lru.queue.MoveToFront(v)
	} else {
		listItem := lru.queue.PushFront(value)
		listItem.Key = key
		lru.items[key] = listItem
	}

	if lru.queue.Len() > lru.capacity {
		last := lru.queue.Back()
		lru.queue.Remove(last)
		delete(lru.items, lru.queue.Back().Key)
	}

	return ok
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	mu.Lock()
	defer mu.Unlock()

	v, ok := lru.items[key]

	if ok {
		lru.queue.MoveToFront(v)
		return v.Value, true
	}

	return nil, false
}

func (lru *lruCache) Clear() {
	mu.Lock()
	defer mu.Unlock()

	lru.queue = NewList()
	lru.items = make(map[Key]*ListItem)
}
