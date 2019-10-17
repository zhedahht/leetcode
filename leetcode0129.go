/*
LeetCode 129: https://leetcode.com/problems/sum-root-to-leaf-numbers/
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
func helper129(root *TreeNode, num int) int {
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return num
	}

	result := 0
	if root.Left != nil {
		result += helper129(root.Left, num)
	}

	if root.Right != nil {
		result += helper129(root.Right, num)
	}

	return result
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return helper129(root, 0)
}
