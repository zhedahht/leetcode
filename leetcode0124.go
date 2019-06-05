/*
LeetCode 124: https://leetcode.com/problems/binary-tree-maximum-path-sum/
*/

package leetcode

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_, _, result := helper124(root)
	return result
}

func helper124(root *TreeNode) (int, int, int) {
	var leftLeft, leftRight, leftResult int
	if root.Left != nil {
		leftLeft, leftRight, leftResult = helper124(root.Left)
	}

	var rightLeft, rightRight, rightResult int
	if root.Right != nil {
		rightLeft, rightRight, rightResult = helper124(root.Right)
	}

	rootLeft, rootRight := root.Val, root.Val
	leftMax := int(math.Max(float64(leftLeft), float64(leftRight)))
	if leftMax > 0 {
		rootLeft += leftMax
	}

	rightMax := int(math.Max(float64(rightLeft), float64(rightRight)))
	if rightMax > 0 {
		rootRight += rightMax
	}

	result := rootLeft + rootRight - root.Val
	if root.Left != nil {
		result = int(math.Max(float64(result), float64(leftResult)))
	}

	if root.Right != nil {
		result = int(math.Max(float64(result), float64(rightResult)))
	}

	return rootLeft, rootRight, result
}
