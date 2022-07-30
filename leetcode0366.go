/*
LeetCode 366: https://leetcode.com/problems/find-leaves-of-binary-tree/
*/

package leetcode

func findLeaves(root *TreeNode) [][]int {
	graph, ins := buildGraph(root)

	queue1 := make([]*TreeNode, 0)
	queue2 := make([]*TreeNode, 0)
	for node, in := range ins {
		if in == 0 {
			queue1 = append(queue1, node)
		}
	}

	result := make([][]int, 0)
	level := make([]int, 0)
	for len(queue1) > 0 {
		node := queue1[0]
		queue1 = queue1[1:]
		level = append(level, node.Val)

		next := graph[node]
		ins[next]--
		if ins[next] == 0 {
			queue2 = append(queue2, next)
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]*TreeNode, 0)
			result = append(result, level)
			level = make([]int, 0)
		}
	}

	return result
}

func buildGraph(root *TreeNode) (map[*TreeNode]*TreeNode, map[*TreeNode]int) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	graph := make(map[*TreeNode]*TreeNode)
	ins := make(map[*TreeNode]int)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if _, exists := ins[node]; !exists {
			ins[node] = 0
		}

		if node.Left != nil {
			graph[node.Left] = node
			ins[node]++
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			graph[node.Right] = node
			ins[node]++
			queue = append(queue, node.Right)
		}
	}
	return graph, ins
}
