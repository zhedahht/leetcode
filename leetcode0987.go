/*
LeetCode 987: https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/description/
*/

package leetcode

import "sort"

func verticalTraversal(root *TreeNode) [][]int {
	nodes := make(map[int][]int)
	nodesInLevel := make(map[int][]int)
	nodeQueue1, nodeQueue2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	posQueue1, posQueue2 := make([]int, 0), make([]int, 0)
	leftMost, rightMost := 0, 0
	if root != nil {
		nodeQueue1 = append(nodeQueue1, root)
		posQueue1 = append(posQueue1, 0)
	}

	for len(nodeQueue1) > 0 {
		node := nodeQueue1[0]
		pos := posQueue1[0]
		nodeQueue1 = nodeQueue1[1:]
		posQueue1 = posQueue1[1:]

		leftMost, rightMost = min(leftMost, pos), max(rightMost, pos)
		if _, exists := nodesInLevel[pos]; !exists {
			nodesInLevel[pos] = make([]int, 0)
		}

		nodesInLevel[pos] = append(nodesInLevel[pos], node.Val)
		if node.Left != nil {
			nodeQueue2 = append(nodeQueue2, node.Left)
			posQueue2 = append(posQueue2, pos-1)
		}

		if node.Right != nil {
			nodeQueue2 = append(nodeQueue2, node.Right)
			posQueue2 = append(posQueue2, pos+1)
		}

		if len(nodeQueue1) == 0 {
			nodeQueue1 = nodeQueue2
			nodeQueue2 = make([]*TreeNode, 0)
			posQueue1 = posQueue2
			posQueue2 = make([]int, 0)

			for p, ns := range nodesInLevel {
				sort.Slice(ns, func(i, j int) bool {
					return ns[i] < ns[j]
				})

				if _, exists := nodes[p]; !exists {
					nodes[p] = make([]int, 0)
				}

				nodes[p] = append(nodes[p], ns...)
			}

			nodesInLevel = make(map[int][]int)
		}
	}

	result := make([][]int, 0)
	for i := leftMost; i <= rightMost; i++ {
		result = append(result, nodes[i])
	}

	return result
}
