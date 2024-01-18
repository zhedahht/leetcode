/*
LeetCode 669: https://leetcode.com/problems/trim-a-binary-search-tree/description/
*/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBSTCore(parent, node *TreeNode, isLeft bool, low, high int) {
	if node == nil {
		if isLeft {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else if node.Val < low {
		trimBSTCore(parent, node.Right, isLeft, low, high)
	} else if node.Val <= high {
		if isLeft {
			parent.Left = node
		} else {
			parent.Right = node
		}

		trimBSTCore(node, node.Left, true, low, high)
		trimBSTCore(node, node.Right, false, low, high)
	} else {
		trimBSTCore(parent, node.Left, isLeft, low, high)
	}
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	dummy := &TreeNode{
		Left: root,
	}

	trimBSTCore(dummy, root, true, low, high)
	return dummy.Left
}
