package main

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	keyToDataMapper map[int]*list.Element
	list            *list.List
	capacity        int
	mu              sync.RWMutex
}

type Pair struct {
	key int
	val int
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		keyToDataMapper: make(map[int]*list.Element),
		list:            list.New(),
		capacity:        capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	// to check if the key exists in the mapper
	if element, ok := lru.keyToDataMapper[key]; ok {
		// move this element to front of the linked list, since it has become most recently used
		lru.list.MoveToFront(element)
		// extract the value stored in the node/element
		return element.Value.(Pair).val
	}

	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	// check if this key exists in the mapper
	if element, ok := lru.keyToDataMapper[key]; ok {
		// update the current stored value in the node/element
		element.Value.(*Pair).val = value

		// move this node to front of linked list
		lru.list.MoveToFront(element)
	} else {
		// if this key does not exist in the mapper
		newElement := lru.list.PushFront(Pair{key: key, val: value})

		// update the mapper
		lru.keyToDataMapper[key] = newElement
		lru.capacity--
	}

	if lru.capacity < 0 {
		// get the last element from the linked list
		lastElement := lru.list.Back()

		// remove from mapper
		delete(lru.keyToDataMapper, lastElement.Value.(Pair).key)

		// remove from linked list
		lru.list.Remove(lastElement)
		lru.capacity++
	}
}
