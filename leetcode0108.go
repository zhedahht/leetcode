/*
LeetCode 108: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
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
func sortedArrayToBST(nums []int) *TreeNode {
	return helper108(nums, 0, len(nums)-1)
}

func helper108(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper108(nums, start, mid-1)
	root.Right = helper108(nums, mid+1, end)
	return root
}
