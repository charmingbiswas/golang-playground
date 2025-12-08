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

func (this *LRUCache) Get(key int) int {
	this.mu.Lock()
	defer this.mu.Unlock()
	// to check if the key exists in the mapper
	if element, ok := this.keyToDataMapper[key]; ok {
		// move this element to front of the linked list, since it has become most recently used
		this.list.MoveToFront(element)
		// extract the value stored in the node/element
		return element.Value.(Pair).val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	// check if this key exists in the mapper
	if element, ok := this.keyToDataMapper[key]; ok {
		// update the current stored value in the node/element
		element.Value.(*Pair).val = value

		// move this node to front of linked list
		this.list.MoveToFront(element)
	} else {
		// if this key does not exist in the mapper
		newElement := this.list.PushFront(Pair{key: key, val: value})

		// update the mapper
		this.keyToDataMapper[key] = newElement
		this.capacity--
	}

	if this.capacity < 0 {
		// get the last element from the linked list
		lastElement := this.list.Back()

		// remove from mapper
		delete(this.keyToDataMapper, lastElement.Value.(Pair).key)

		// remove from linked list
		this.list.Remove(lastElement)
		this.capacity++
	}
}
