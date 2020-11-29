/*
LeetCode 117: https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
*/

package leetcode

/**
 * Definition for a Node.
 * type Node TreeNodeWithNext {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// The function name should be connect. Renamed it to avoid conflicts.
func connect117(root *TreeNodeWithNext) *TreeNodeWithNext {
	node := root
	for node != nil {
		sibling, next, isLeft := getNextChild117(node, true)
		child := next
		for sibling != nil && child != nil {
			var nextSibling, nextChild *TreeNodeWithNext
			if isLeft {
				nextSibling, nextChild, isLeft = getNextChild117(sibling, false)
			} else {
				nextSibling, nextChild, isLeft = getNextChild117(sibling.Next, true)
			}

			child.Next = nextChild

			sibling, child = nextSibling, nextChild
		}

		node = next
	}

	return root
}

func getNextChild117(parent *TreeNodeWithNext, tryLeft bool) (*TreeNodeWithNext, *TreeNodeWithNext, bool) {
	for parent != nil {
		if tryLeft && parent.Left != nil {
			return parent, parent.Left, true
		}

		if parent.Right != nil {
			return parent, parent.Right, false
		}

		parent = parent.Next
		tryLeft = true
	}

	return nil, nil, false
}
