/*
LeetCode 145: https://leetcode.com/problems/binary-tree-postorder-traversal/
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
func postorderTraversal(root *TreeNode) []int {
	stack, result := make([]*TreeNode, 0), make([]int, 0)
	var prev *TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		if node.Right != nil && node.Right != prev {
			node = node.Right
		} else {
			result = append(result, node.Val)
			stack = stack[:len(stack)-1]
			prev, node = node, nil
		}
	}

	return result
}
