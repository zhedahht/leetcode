/*
LeetCode 863: https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
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
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if K == 0 {
		return []int{target.Val}
	}

	nodeToNexts := make(map[int][]int)
	buildGraph863(root, nodeToNexts)
	dist := 0
	var queue1, queue2 []int
	visited := make(map[int]bool)
	queue1 = append(queue1, target.Val)
	visited[target.Val] = true
	for len(queue1) > 0 {
		node := queue1[0]
		queue1 = queue1[1:]

		for _, next := range nodeToNexts[node] {
			if !visited[next] {
				queue2 = append(queue2, next)
				visited[next] = true
			}
		}

		if len(queue1) == 0 {
			dist++
			if dist == K {
				return queue2
			}

			queue1 = queue2
			queue2 = make([]int, 0)
		}
	}

	return make([]int, 0)
}

func addKeyVal863(k1, k2 int, keyToValues map[int][]int) {
	keyToValues[k1] = append(keyToValues[k1], k2)
	keyToValues[k2] = append(keyToValues[k2], k1)
}

func buildGraph863(root *TreeNode, nodeToNexts map[int][]int) {
	if root.Left != nil {
		addKeyVal863(root.Val, root.Left.Val, nodeToNexts)
		buildGraph863(root.Left, nodeToNexts)
	}

	if root.Right != nil {
		addKeyVal863(root.Val, root.Right.Val, nodeToNexts)
		buildGraph863(root.Right, nodeToNexts)
	}
}
