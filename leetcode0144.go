/*
LeetCode 144: https://leetcode.com/problems/binary-tree-preorder-traversal/
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
func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	node := root
	result := make([]int, 0)
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			result = append(result, node.Val)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}

	return result
}
