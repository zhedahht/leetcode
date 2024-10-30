/*
LeetCode 270: https://leetcode.com/problems/closest-binary-search-tree-value/description/
*/

package leetcode

import "math"

func closestValue(root *TreeNode, target float64) int {
	var result *TreeNode
	node := root
	for node != nil {
		if result == nil {
			result = node
		} else if math.Abs(float64(node.Val)-target) < math.Abs(float64(result.Val)-target) ||
			(math.Abs(float64(node.Val)-target) == math.Abs(float64(result.Val)-target) && node.Val < result.Val) {
			result = node
		}

		if target < float64(node.Val) {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return result.Val
}
