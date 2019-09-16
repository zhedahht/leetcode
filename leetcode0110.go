/*
LeetCode 110: https://leetcode.com/problems/balanced-binary-tree/
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
func helper110(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	left, leftDepth := helper110(root.Left)
	right, rightDepth := helper110(root.Right)
	if !left || !right || leftDepth-rightDepth > 1 || leftDepth-rightDepth < -1 {
		return false, 0
	}

	if leftDepth > rightDepth {
		return true, leftDepth + 1
	}

	return true, rightDepth + 1
}

func isBalanced(root *TreeNode) bool {
	balanced, _ := helper110(root)
	return balanced
}
