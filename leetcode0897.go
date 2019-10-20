/*
LeetCode 897: https://leetcode.com/problems/increasing-order-search-tree/
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
func increasingBST(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	var prev, newRoot *TreeNode
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.Right = node
		} else {
			newRoot = node
		}

		prev = node
		node = node.Right
		prev.Left, prev.Right = nil, nil
	}

	return newRoot
}
