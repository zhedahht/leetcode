/*
LeetCode 655: https://leetcode.com/problems/print-binary-tree/
*/

package leetcode

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	levels := bfs(root)
	return pad(levels)
}

func bfs(root *TreeNode) [][]string {
	var levels [][]string
	level := make([]string, 0)
	queue1 := make([]*TreeNode, 0)
	queue2 := make([]*TreeNode, 0)

	queue1 = append(queue1, root)
	meetValue := false
	for len(queue1) > 0 {
		node := queue1[0]
		queue1 = queue1[1:]

		if node == nil {
			level = append(level, "")
		} else {
			val := strconv.Itoa(node.Val)
			level = append(level, val)
			meetValue = true
		}

		if node != nil {
			queue2 = append(queue2, node.Left)
			queue2 = append(queue2, node.Right)
		} else {
			queue2 = append(queue2, nil)
			queue2 = append(queue2, nil)
		}

		if len(queue1) == 0 && meetValue {
			levels = append(levels, level)
			level = make([]string, 0)

			queue1 = queue2
			queue2 = make([]*TreeNode, 0)

			meetValue = false
		}
	}

	return levels
}

func pad(levels [][]string) [][]string {
	var result [][]string
	depth := len(levels)
	for i, level := range levels {
		count := (1 << uint(depth-i-1)) - 1
		fmt.Printf("count: %d\n", count)

		row := make([]string, 0)
		for j := 0; j < len(level); j++ {
			for k := 0; k < count; k++ {
				row = append(row, "")
			}

			if j != 0 {
				row = append(row, "")
			}

			row = append(row, level[j])

			for k := 0; k < count; k++ {
				row = append(row, "")
			}
		}

		result = append(result, row)
	}

	return result
}
