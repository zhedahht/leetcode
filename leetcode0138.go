/*
LeetCode 138: https://leetcode.com/problems/copy-list-with-random-pointer/
*/

package leetcode

// Node138 is the node type for LeetCode 138.
// It should be Node. Renamed to avoid conflicts.
type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

func copyRandomList(head *Node138) *Node138 {
	node := head
	for node != nil {
		clone := &Node138{
			Val:  node.Val,
			Next: node.Next,
		}

		node.Next = clone
		node = clone.Next
	}

	node = head
	for node != nil {
		clone := node.Next
		random := node.Random
		if random != nil {
			clone.Random = random.Next
		}

		node = clone.Next
	}

	newNode := &Node138{}
	dummy, node := newNode, head
	for node != nil {
		clone := node.Next
		newNode.Next = clone
		newNode = newNode.Next

		node.Next = clone.Next
		node = node.Next
	}

	return dummy.Next
}
