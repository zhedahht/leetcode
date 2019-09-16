/*
LeetCode 112: https://leetcode.com/problems/path-sum/
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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}

	if root.Left != nil && hasPathSum(root.Left, sum-root.Val) {
		return true
	}

	return root.Right != nil && hasPathSum(root.Right, sum-root.Val)

}
