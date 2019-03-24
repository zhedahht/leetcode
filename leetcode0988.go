/*
LeetCode 988: https://leetcode.com/problems/smallest-string-starting-from-leaf/
*/

package leetcode

import "strings"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	return helper988(root, make([]int, 0), "")
}

func helper988(root *TreeNode, path []int, min string) string {
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		str := make([]byte, len(path))
		for i, val := range path {
			str[len(str)-i-1] = byte(val + 97)
		}

		min = getSmaller(string(str), min)
	}

	if root.Left != nil {
		min = getSmaller(helper988(root.Left, path, min), min)
	}

	if root.Right != nil {
		min = getSmaller(helper988(root.Right, path, min), min)
	}

	return min
}

func getSmaller(str, min string) string {
	if len(min) == 0 {
		return str
	}

	if strings.Compare(min, str) <= 0 {
		return min
	}

	return str
}
