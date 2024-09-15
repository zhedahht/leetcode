/*
LeetCode 460: https://leetcode.com/problems/lfu-cache/description/
*/

package leetcode

type keyValueNode struct {
	key   int
	value int
	host  *countGroupNode
	prev  *keyValueNode
	next  *keyValueNode
}

type countGroupNode struct {
	count int
	head  *keyValueNode
	tail  *keyValueNode
	prev  *countGroupNode
	next  *countGroupNode
}

func createCountGroupNode(count int) *countGroupNode {
	head, tail := &keyValueNode{}, &keyValueNode{}
	head.next = tail
	tail.prev = head
	return &countGroupNode{
		count: count,
		head:  head,
		tail:  tail,
	}
}

type LFUCache struct {
	capacity int
	data     map[int]*keyValueNode
	head     *countGroupNode
	tail     *countGroupNode
}

// The name should be Constructor. It's renamed to avoid conflicts.
func Constructor0460(capacity int) LFUCache {
	head, tail := createCountGroupNode(0), createCountGroupNode(0)
	head.next = tail
	tail.prev = head
	return LFUCache{
		capacity: capacity,
		data:     make(map[int]*keyValueNode),
		head:     head,
		tail:     tail,
	}
}

func (this *LFUCache) Get(key int) int {
	node, exists := this.data[key]
	if !exists {
		return -1
	}

	val := node.value
	count := node.host.count
	nextHost := node.host.next
	this.removeNodeFromHost(node)
	this.addNodeToHost(node, nextHost, count+1)
	return val
}

func (this *LFUCache) Put(key int, value int) {
	node, exists := this.data[key]
	if !exists {
		if len(this.data) == this.capacity {
			host := this.head.next
			node := host.head.next
			key := node.key
			this.removeNodeFromHost(node)
			delete(this.data, key)
		}

		node = &keyValueNode{
			key:   key,
			value: value,
		}

		this.addNodeToHost(node, this.head.next, 1)
		this.data[key] = node
	} else {
		node.value = value
		count := node.host.count
		nextHost := node.host.next
		this.removeNodeFromHost(node)
		this.addNodeToHost(node, nextHost, count+1)
	}
}

func (this *LFUCache) removeNodeFromHost(node *keyValueNode) {
	host := node.host

	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev

	if host.head.next == host.tail {
		prevHost, nextHost := host.prev, host.next
		prevHost.next = nextHost
		nextHost.prev = prevHost
	}
}

func (this *LFUCache) addNodeToHost(node *keyValueNode, nextHost *countGroupNode, count int) {
	if nextHost == this.tail || nextHost.count != count {
		newHost := createCountGroupNode(count)
		prev := nextHost.prev
		prev.next = newHost
		newHost.prev = prev
		newHost.next = nextHost
		nextHost.prev = newHost

		nextHost = newHost
	}

	node.host = nextHost

	prevNode, nextNode := nextHost.tail.prev, nextHost.tail
	node.prev = prevNode
	prevNode.next = node
	node.next = nextNode
	nextNode.prev = node
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
