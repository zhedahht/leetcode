/*
LeetCode 938: https://leetcode.com/problems/range-sum-of-bst/
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
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	result := 0
	if root.Val >= L && root.Val <= R {
		result += root.Val
	}

	if root.Val >= L {
		result += rangeSumBST(root.Left, L, R)
	}

	if root.Val <= R {
		result += rangeSumBST(root.Right, L, R)
	}

	return result
}
