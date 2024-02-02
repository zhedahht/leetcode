/*
LeetCode 701: https://leetcode.com/problems/insert-into-a-binary-search-tree/description/
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	dummy := &TreeNode{
		Val:  math.MaxInt,
		Left: root,
	}

	insertIntoBSTHelper(dummy, val)
	return dummy.Left
}

func insertIntoBSTHelper(root *TreeNode, val int) {
	if root.Val > val {
		left := root.Left
		if left == nil {
			root.Left = &TreeNode{
				Val: val,
			}
		} else {
			insertIntoBSTHelper(left, val)
		}
	} else {
		right := root.Right
		if right == nil {
			root.Right = &TreeNode{
				Val: val,
			}
		} else {
			insertIntoBSTHelper(right, val)
		}
	}
}
