/*
LeetCode 156: https://leetcode.com/problems/binary-tree-upside-down/
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
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	_, newRoot := helper156(root)
	return newRoot
}

func helper156(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil || root.Left == nil {
		return root, root
	}

	left, newRoot := helper156(root.Left)
	left.Right = root
	left.Left = root.Right
	root.Left = nil
	root.Right = nil

	return root, newRoot
}
