/*
LeetCode 99: https://leetcode.com/problems/recover-binary-search-tree/
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
func recoverTree(root *TreeNode) {
	var node1, node2, prev *TreeNode
	cur := root
	stack := make([]*TreeNode, 0)
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && prev.Val > cur.Val {
			node2 = cur
			if node1 == nil {
				node1 = prev
			} else {
				break
			}
		}

		prev = cur
		cur = cur.Right
	}

	if node1 != nil && node2 != nil {
		node1.Val, node2.Val = node2.Val, node1.Val
	}
}
