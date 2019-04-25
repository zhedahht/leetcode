/*
LeetCode 98: https://leetcode.com/problems/validate-binary-search-tree/
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
func isValidBST(root *TreeNode) bool {
	node := root
	var prev *TreeNode
	stack := make([]*TreeNode, 0)
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		if prev != nil && prev.Val >= node.Val {
			return false
		}

		stack = stack[:len(stack)-1]
		prev = node
		node = node.Right
	}

	return true
}
