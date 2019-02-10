/*
LeetCode 298: https://leetcode.com/problems/binary-tree-longest-consecutive-sequence/
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
func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0
	helper298(root, nil, &max)
	return max
}

func helper298(root *TreeNode, parent *TreeNode, max *int) {
	if parent != nil && root.Val == parent.Val+1 {
		*max++
	} else {
		*max = 1
	}

	maxLeft := *max
	if root.Left != nil {
		helper298(root.Left, root, &maxLeft)
	}

	maxRight := *max
	if root.Right != nil {
		helper298(root.Right, root, &maxRight)
	}

	*max = int(math.Max(float64(*max), float64(maxLeft)))
	*max = int(math.Max(float64(*max), float64(maxRight)))
}
