/*
LeetCode 510: https://leetcode.com/problems/inorder-successor-in-bst-ii/description/
*/

package leetcode

// The name should be Node. It's renamed to avoid conflicts.
type Node510 struct {
	Val    int
	Left   *Node510
	Right  *Node510
	Parent *Node510
}

func inorderSuccessor(node *Node510) *Node510 {
	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}

		return node
	}

	for node.Parent != nil && node.Parent.Val < node.Val {
		node = node.Parent
	}

	if node.Parent != nil {
		return node.Parent
	}

	return nil
}
