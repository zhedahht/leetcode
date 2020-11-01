/*
LeetCode 314: https://leetcode.com/problems/binary-tree-vertical-order-traversal/
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
func verticalOrder(root *TreeNode) [][]int {
	type nodePos struct {
		node *TreeNode
		row  int
		col  int
	}

	addNode := func(node *TreeNode, row, col int, queue []*nodePos) []*nodePos {
		return append(queue, &nodePos{
			node: node,
			row:  row,
			col:  col,
		})
	}

	queue := make([]*nodePos, 0)
	if root != nil {
		queue = addNode(root, 0, 0, queue)
	}

	pos := make(map[int][]int)
	min, max := 0, 0
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		pos[n.col] = append(pos[n.col], n.node.Val)

		if n.node.Left != nil {
			queue = addNode(n.node.Left, n.row+1, n.col-1, queue)
			if n.col-1 < min {
				min = n.col - 1
			}
		}

		if n.node.Right != nil {
			queue = addNode(n.node.Right, n.row+1, n.col+1, queue)
			if n.col+1 > max {
				max = n.col + 1
			}
		}
	}

	result := make([][]int, 0)
	for i := min; i <= max; i++ {
		if pos[i] != nil {
			result = append(result, pos[i])
		}
	}

	return result
}
