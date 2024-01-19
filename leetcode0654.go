/*
LeetCode 654: https://leetcode.com/problems/maximum-binary-tree/description/
*/

package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := math.MinInt
	maxIndex := -1
	for i, num := range nums {
		if num > max {
			max = num
			maxIndex = i
		}
	}

	return &TreeNode{
		Val:   nums[maxIndex],
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}
}
