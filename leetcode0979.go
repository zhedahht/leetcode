/*
LeetCode 979: https://leetcode.com/problems/distribute-coins-in-binary-tree/
*/

package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	distribute, _, _ := helper979(root)
	return distribute
}

func helper979(root *TreeNode) (int, int, int) {
	distributeLeft, countLeft, totalLeft := 0, 0, 0
	if root.Left != nil {
		distributeLeft, countLeft, totalLeft = helper979(root.Left)
	}

	distributeRight, countRight, totalRight := 0, 0, 0
	if root.Right != nil {
		distributeRight, countRight, totalRight = helper979(root.Right)
	}

	count := countLeft + countRight + 1
	total := totalLeft + totalRight + root.Val
	return int(math.Abs(float64(total-count))) + distributeLeft + distributeRight, count, total
}
