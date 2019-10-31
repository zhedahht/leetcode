/*
LeetCode 865: https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/
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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	result, _ := help865(root)
	return result
}

func help865(root *TreeNode) (*TreeNode, int) {
	left, right := 0, 0
	var leftResult, rightResult *TreeNode
	if root.Left != nil {
		leftResult, left = help865(root.Left)
	}

	if root.Right != nil {
		rightResult, right = help865(root.Right)
	}

	if left > right {
		return leftResult, left + 1
	}

	if left < right {
		return rightResult, right + 1
	}

	return root, left + 1
}
