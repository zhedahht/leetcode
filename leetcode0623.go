/*
LeetCode 623: https://leetcode.com/problems/add-one-row-to-tree/description/
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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	dummy := &TreeNode{
		Left: root,
	}

	level := 0
	queue1, queue2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	queue1 = append(queue1, dummy)

	for len(queue1) > 0 {
		node := queue1[0]
		queue1 = queue1[1:]

		left, right := node.Left, node.Right
		if left != nil {
			queue2 = append(queue2, node.Left)
		}

		if right != nil {
			queue2 = append(queue2, node.Right)
		}

		if level == depth-1 {
			newLeft := &TreeNode{
				Val:  val,
				Left: left,
			}

			newRight := &TreeNode{
				Val:   val,
				Right: right,
			}

			node.Left = newLeft
			node.Right = newRight
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]*TreeNode, 0)
			level++
		}
	}

	return dummy.Left
}
