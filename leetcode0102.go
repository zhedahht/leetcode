/*
LeetCode 102: https://leetcode.com/problems/binary-tree-level-order-traversal/
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
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue1, queue2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	level := make([]int, 0)
	queue1 = append(queue1, root)
	for len(queue1) > 0 {
		node := queue1[0]
		queue1 = queue1[1:]
		level = append(level, node.Val)

		if node.Left != nil {
			queue2 = append(queue2, node.Left)
		}

		if node.Right != nil {
			queue2 = append(queue2, node.Right)
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]*TreeNode, 0)
			result = append(result, level)
			level = make([]int, 0)
		}
	}

	return result
}
