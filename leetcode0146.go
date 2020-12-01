/*
LeetCode 146: https://leetcode.com/problems/lru-cache/
*/

package leetcode

// LRUCacheListNode is the node type of a double-linked list
type LRUCacheListNode struct {
	key  int
	val  int
	prev *LRUCacheListNode
	next *LRUCacheListNode
}

// LRUCache is the Least Recently Used (LRU) cache
type LRUCache struct {
	capacity int
	cache    map[int]*LRUCacheListNode
	head     *LRUCacheListNode
	tail     *LRUCacheListNode
}

// Constructor createa a Least Recently Used (LRU) cache
func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*LRUCacheListNode),
		head:     &LRUCacheListNode{},
		tail:     &LRUCacheListNode{},
	}

	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head

	return lruCache
}

// Get gets the value of a specified key
func (c *LRUCache) Get(key int) int {
	node := c.cache[key]
	if node == nil {
		return -1
	}

	val := node.val
	c.moveToTail(node)
	return val
}

// Put add or update a pair of <key, value>
func (c *LRUCache) Put(key int, value int) {
	node := c.cache[key]
	if node != nil {
		node.val = value
		c.moveToTail(node)
	} else {
		node = &LRUCacheListNode{
			key: key,
			val: value,
		}

		if c.capacity == len(c.cache) {
			first := c.head.next
			delete(c.cache, first.key)
			c.deleteNode(first)
		}

		c.cache[key] = node
		c.addToTail(node)
	}
}

func (c *LRUCache) addToTail(node *LRUCacheListNode) {
	prev := c.tail.prev
	prev.next, node.prev = node, prev
	c.tail.prev, node.next = node, c.tail
}

func (c *LRUCache) moveToTail(node *LRUCacheListNode) {
	c.deleteNode(node)
	c.addToTail(node)
}

func (c *LRUCache) deleteNode(node *LRUCacheListNode) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}
