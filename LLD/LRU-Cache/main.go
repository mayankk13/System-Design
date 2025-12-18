package main

import (
	"fmt"
	"sync"
)

/*
LRU Cache Design:
- HashMap for O(1) access
- Doubly Linked List to maintain usage order
- Most Recently Used (MRU) -> near head
- Least Recently Used (LRU) -> near tail
*/

// Node represents a doubly linked list node
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCache structure
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // Dummy head
	tail     *Node // Dummy tail
	mutex    sync.Mutex
}

// Constructor initializes the LRU Cache
func Constructor(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// Get returns the value if key exists, else -1
func (lru *LRUCache) Get(key int) int {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if node, ok := lru.cache[key]; ok {
		// Move accessed node to front (MRU)
		lru.remove(node)
		lru.addToFront(node)
		return node.value
	}
	return -1
}

// Put inserts or updates the value
func (lru *LRUCache) Put(key int, value int) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	// If key exists, update and move to front
	if node, ok := lru.cache[key]; ok {
		node.value = value
		lru.remove(node)
		lru.addToFront(node)
		return
	}

	// If capacity exceeded, remove LRU node
	if len(lru.cache) == lru.capacity {
		lruNode := lru.tail.prev
		lru.remove(lruNode)
		delete(lru.cache, lruNode.key)
	}

	// Insert new node
	newNode := &Node{
		key:   key,
		value: value,
	}
	lru.cache[key] = newNode
	lru.addToFront(newNode)
}

// remove removes a node from the linked list
func (lru *LRUCache) remove(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

// addToFront adds node right after head (MRU)
func (lru *LRUCache) addToFront(node *Node) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// ----------------- Example Usage -----------------

func main() {
	lru := Constructor(2)

	lru.Put(1, 10)
	lru.Put(2, 20)
	fmt.Println(lru.Get(1)) // 10

	lru.Put(3, 30)          // Evicts key 2
	fmt.Println(lru.Get(2)) // -1

	lru.Put(4, 40)          // Evicts key 1
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 30
	fmt.Println(lru.Get(4)) // 40
}
