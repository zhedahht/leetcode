/*
LeetCode 430: https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/description/
*/

package leetcode

// The type name should be Node. It's renamed to avoid conflicts.
type Node430 struct {
	Val   int
	Prev  *Node430
	Next  *Node430
	Child *Node430
}

// The function name should be flatten. It's renamed to avoid conflicts.
func flatten430(root *Node430) *Node430 {
	first, _ := helper430(root)
	return first
}

func helper430(root *Node430) (*Node430, *Node430) {
	first, last, node := root, root, root
	for node != nil {
		last = node
		if node.Child != nil {
			childFirst, childLast := helper430(node.Child)
			node.Child = nil
			next := node.Next
			node.Next = childFirst
			childFirst.Prev = node
			childLast.Next = next
			if next != nil {
				next.Prev = childLast
			}
		}

		node = node.Next
	}

	return first, last
}
