/*
LeetCode 103: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	addNodeIfNotNil := func(stack []*TreeNode, node *TreeNode) []*TreeNode {
		if node != nil {
			stack = append(stack, node)
		}

		return stack
	}

	result := make([][]int, 0)
	if root == nil {
		return result
	}

	var stacks [2][]*TreeNode
	stacks[0], stacks[1] = make([]*TreeNode, 0), make([]*TreeNode, 0)
	cur, next := 0, 1
	stacks[cur] = append(stacks[cur], root)
	level := make([]int, 0)
	for len(stacks[cur]) > 0 {
		node := stacks[cur][len(stacks[cur])-1]
		stacks[cur] = stacks[cur][:len(stacks[cur])-1]
		level = append(level, node.Val)

		if cur == 0 {
			stacks[next] = addNodeIfNotNil(stacks[next], node.Left)
			stacks[next] = addNodeIfNotNil(stacks[next], node.Right)
		} else {
			stacks[next] = addNodeIfNotNil(stacks[next], node.Right)
			stacks[next] = addNodeIfNotNil(stacks[next], node.Left)
		}

		if len(stacks[cur]) == 0 {
			result = append(result, level)
			level = make([]int, 0)
			cur, next = 1-cur, cur
		}
	}

	return result
}
