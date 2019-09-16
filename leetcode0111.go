/*
LeetCode 111: https://leetcode.com/problems/minimum-depth-of-binary-tree/
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
func helper111(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := helper111(root.Left), helper111(root.Right)
	if left == 0 {
		return right + 1
	}

	if right == 0 {
		return left + 1
	}

	if left < right {
		return left + 1
	}

	return right + 1
}

func minDepth(root *TreeNode) int {
	return helper111(root)
}
