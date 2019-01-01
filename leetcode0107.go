/*
LeetCode 107: https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	queue1 := make([]*TreeNode, 0)
	queue2 := make([]*TreeNode, 0)
	result := make([][]int, 0)

	queue1 = append(queue1, root)
	level := make([]int, 0)
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

	for i := 0; i < len(result)/2; i++ {
		j := len(result) - 1 - i
		result[i], result[j] = result[j], result[i]
	}

	return result
}
