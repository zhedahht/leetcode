/*
LeetCode 432: https://leetcode.com/problems/all-oone-data-structure/description/
*/

package leetcode

type countNode struct {
	keys  map[string]bool
	count int
	prev  *countNode
	next  *countNode
}

type AllOne struct {
	keyCounts map[string]*countNode
	head      *countNode
	tail      *countNode
}

// The name should be Constructor. It's renamed to avoid conflicts.
func Constructor0432() AllOne {
	head := &countNode{}
	tail := &countNode{}
	head.next = tail
	tail.prev = head
	return AllOne{
		keyCounts: make(map[string]*countNode),
		head:      head,
		tail:      tail,
	}
}

func (this *AllOne) Inc(key string) {
	node, exists := this.keyCounts[key]
	if exists {
		newCount := node.count + 1
		next := node.next
		if next == this.tail || next.count != newCount {
			newNode := &countNode{
				keys:  make(map[string]bool),
				count: newCount,
				prev:  node,
				next:  next,
			}

			node.next = newNode
			next.prev = newNode
		}

		next = node.next
		next.keys[key] = true
		this.keyCounts[key] = next

		delete(node.keys, key)
		if len(node.keys) == 0 {
			prev := node.prev
			prev.next = next
			next.prev = prev
		}
	} else {
		if this.head.next == this.tail || this.head.next.count != 1 {
			node := &countNode{
				keys:  make(map[string]bool),
				count: 1,
				prev:  this.head,
				next:  this.head.next,
			}

			next := this.head.next
			this.head.next = node
			node.prev = this.head
			next.prev = node
			node.next = next
		}

		node = this.head.next
		node.keys[key] = true
		this.keyCounts[key] = node
	}
}

func (this *AllOne) Dec(key string) {
	node, exists := this.keyCounts[key]
	if !exists {
		return
	}

	newCount := node.count - 1
	prev := node.prev
	if newCount != 0 {
		if prev == this.head || prev.count != newCount {
			newNode := &countNode{
				keys:  make(map[string]bool),
				count: newCount,
				prev:  prev,
				next:  node,
			}

			node.prev = newNode
			prev.next = newNode
		}

		prev = node.prev
		prev.keys[key] = true
		this.keyCounts[key] = prev
	} else {
		delete(this.keyCounts, key)
	}

	delete(node.keys, key)
	if len(node.keys) == 0 {
		next := node.next
		prev.next = next
		next.prev = prev
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.head.next == this.tail {
		return ""
	}

	for k := range this.tail.prev.keys {
		return k
	}

	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.head.next == this.tail {
		return ""
	}

	for k := range this.head.next.keys {
		return k
	}

	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
