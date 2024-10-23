/*
LeetCode 426: https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/description/
*/

package leetcode

func treeToDoublyList(root *TreeNode) *TreeNode {
	node := root
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	var head *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev == nil {
			head = node
		} else {
			prev.Right = node
			node.Left = prev
		}

		prev = node
		node = node.Right
	}

	if prev != nil {
		head.Left = prev
		prev.Right = head
	}

	return head
}
