/*
LeetCode 742: https://leetcode.com/problems/closest-leaf-in-a-binary-tree/
*/

package leetcode

func findClosestLeaf(root *TreeNode, k int) int {
	graph := make(map[int]map[int]bool)
	addNode742(root, graph)
	buildGraph742(root, graph)

	queue := make([]int, 0)
	visited := make(map[int]bool)
	queue = append(queue, k)
	visited[k] = true
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if (n != root.Val && len(graph[n]) == 1) || (n == root.Val && root.Left == nil && root.Right == nil) {
			return n
		}

		for c := range graph[n] {
			if !visited[c] {
				visited[c] = true
				queue = append(queue, c)
			}
		}
	}

	return -1
}

func buildGraph742(root *TreeNode, graph map[int]map[int]bool) {
	if root == nil {
		return
	}

	if root.Left != nil {
		addNode742(root.Left, graph)
		graph[root.Val][root.Left.Val] = true
		graph[root.Left.Val][root.Val] = true
		buildGraph742(root.Left, graph)
	}

	if root.Right != nil {
		addNode742(root.Right, graph)
		graph[root.Val][root.Right.Val] = true
		graph[root.Right.Val][root.Val] = true
		buildGraph742(root.Right, graph)
	}
}

func addNode742(root *TreeNode, graph map[int]map[int]bool) {
	if _, ok := graph[root.Val]; !ok {
		graph[root.Val] = make(map[int]bool)
	}
}
