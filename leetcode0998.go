/*
LeetCode 998: https://leetcode.com/problems/maximum-binary-tree-ii/
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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val < val {
		node := TreeNode{
			Val: val,
		}

		node.Left = root
		return &node
	}

	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
