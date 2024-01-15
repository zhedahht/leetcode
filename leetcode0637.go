/*
LeetCode 637: https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
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
func averageOfLevels(root *TreeNode) []float64 {
	queue1, queue2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	queue1 = append(queue1, root)
	count := 0
	sum := 0
	result := make([]float64, 0)

	for len(queue1) > 0 {
		node := queue1[0]
		count++
		sum += node.Val

		if node.Left != nil {
			queue2 = append(queue2, node.Left)
		}

		if node.Right != nil {
			queue2 = append(queue2, node.Right)
		}

		queue1 = queue1[1:]
		if len(queue1) == 0 {
			result = append(result, float64(sum)/float64(count))
			sum = 0
			count = 0
			queue1 = queue2
			queue2 = make([]*TreeNode, 0)
		}
	}

	return result
}
