/*
LeetCode 114: https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left, right := root.Left, root.Right
	root.Left = nil
	node := left
	if node != nil {
		for node.Right != nil {
			node = node.Right
		}

		root.Right = left
		node.Right = right
	}
}
