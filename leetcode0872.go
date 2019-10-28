/*
LeetCode 872: https://leetcode.com/problems/leaf-similar-trees/
*/

package leetcode

import "reflect"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	vals1 := helper872(root1, make([]int, 0))
	vals2 := helper872(root2, make([]int, 0))
	return reflect.DeepEqual(vals1, vals2)
}

func helper872(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}

	if root.Left == nil && root.Right == nil {
		result = append(result, root.Val)
	}

	if root.Left != nil {
		result = helper872(root.Left, result)
	}

	if root.Right != nil {
		result = helper872(root.Right, result)
	}

	return result
}
