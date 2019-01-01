/*
LeetCode 404: https://leetcode.com/problems/sum-of-left-leaves/
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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return sumOfLeafChildren(root)
}

func sumOfLeft(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	return sumOfLeafChildren(root)
}

func sumOfRight(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}

	return sumOfLeafChildren(root)
}

func sumOfLeafChildren(root *TreeNode) int {
	sum := 0
	if root.Left != nil {
		sum += sumOfLeft(root.Left)
	}

	if root.Right != nil {
		sum += sumOfRight(root.Right)
	}

	return sum
}
