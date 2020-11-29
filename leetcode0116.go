/*
LeetCode 116: https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
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

func connect(root *TreeNodeWithNext) *TreeNodeWithNext {
	node := root
	for node != nil {
		next := node.Left
		for node != nil && node.Left != nil {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
			node = node.Next
		}

		node = next
	}

	return root
}
