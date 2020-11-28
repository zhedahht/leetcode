/*
LeetCode 101: https://leetcode.com/problems/symmetric-tree/
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

func isSame101(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val &&
		isSame101(node1.Left, node2.Right) &&
		isSame101(node1.Right, node2.Left)
}

func isSymmetric(root *TreeNode) bool {
	return isSame101(root, root)
}
