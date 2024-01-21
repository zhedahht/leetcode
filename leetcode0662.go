/*
LeetCode 662: https://leetcode.com/problems/maximum-width-of-binary-tree/description/
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
func widthOfBinaryTree(root *TreeNode) int {
	type nodeWithPos struct {
		node *TreeNode
		pos  int
	}

	maxWidth := 0
	left := 1
	queue1, queue2 := make([]nodeWithPos, 0), make([]nodeWithPos, 0)
	if root != nil {
		queue1 = append(queue1, nodeWithPos{
			node: root,
			pos:  1,
		})
	}

	for len(queue1) > 0 {
		current := queue1[0]
		maxWidth = max(maxWidth, current.pos-left+1)

		if current.node.Left != nil {
			queue2 = append(queue2, nodeWithPos{
				node: current.node.Left,
				pos:  2*current.pos - 1,
			})
		}

		if current.node.Right != nil {
			queue2 = append(queue2, nodeWithPos{
				node: current.node.Right,
				pos:  2 * current.pos,
			})
		}

		queue1 = queue1[1:]
		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]nodeWithPos, 0)
			if len(queue1) > 0 {
				left = queue1[0].pos
			}
		}
	}

	return maxWidth
}
