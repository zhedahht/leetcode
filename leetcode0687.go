/*
LeetCode 687: https://leetcode.com/problems/longest-univalue-path/description/
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
func longestUnivaluePath(root *TreeNode) int {
	_, _, maxUnderRoot := helper687(root)
	return maxUnderRoot - 1
}

func helper687(root *TreeNode) (int, int, int) {
	if root == nil {
		return 1, 1, 1
	}

	fromLeft, fromRight, maxUnderRoot := 1, 1, 0
	if root.Left != nil {
		leftFromLeft, leftFromRight, maxLeft := helper687(root.Left)
		maxUnderRoot = max(maxUnderRoot, maxLeft)
		if root.Val == root.Left.Val {
			fromLeft = max(leftFromLeft, leftFromRight) + 1
		}
	}

	if root.Right != nil {
		rightFromLeft, rightFromRight, maxRight := helper687(root.Right)
		maxUnderRoot = max(maxUnderRoot, maxRight)
		if root.Val == root.Right.Val {
			fromRight = max(rightFromLeft, rightFromRight) + 1
		}
	}

	maxUnderRoot = max(maxUnderRoot, fromLeft+fromRight-1)
	return fromLeft, fromRight, maxUnderRoot
}
