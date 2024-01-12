/*
LeetCode 606:https://leetcode.com/problems/construct-string-from-binary-tree/description/
*/

package leetcode

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := strconv.Itoa(root.Val)
	if root.Left != nil && root.Right != nil {
		result = result + "(" + tree2str(root.Left) + ")" + "(" + tree2str(root.Right) + ")"
	} else if root.Left != nil {
		result = result + "(" + tree2str(root.Left) + ")"
	} else if root.Right != nil {
		result = result + "()" + "(" + tree2str(root.Right) + ")"
	}

	return result
}
