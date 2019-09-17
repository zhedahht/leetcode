/*
LeetCode 113: https://leetcode.com/problems/path-sum-ii/
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
func helper113(root *TreeNode, sum int, path []int, result [][]int) [][]int {
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		clone := make([]int, len(path))
		copy(clone, path)
		result = append(result, clone)
	}

	if root.Left != nil {
		result = helper113(root.Left, sum-root.Val, path, result)
	}

	if root.Right != nil {
		result = helper113(root.Right, sum-root.Val, path, result)
	}

	path = path[:len(path)-1]
	return result
}

func pathSum(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	path := make([]int, 0)
	return helper113(root, sum, path, result)
}
