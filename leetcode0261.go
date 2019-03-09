/*
LeetCode 261: https://leetcode.com/problems/graph-valid-tree/
*/

package leetcode

func validTree(n int, edges [][]int) bool {
	if n != len(edges)+1 {
		return false
	}

	parents := make([]int, n)
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}

	for _, edge := range edges {
		parent1 := getParent(parents, edge[0])
		parent2 := getParent(parents, edge[1])
		if parent1 == parent2 {
			return false
		}

		union(parents, parent1, parent2)
	}

	return true
}

func getParent(parents []int, node int) int {
	parent := parents[node]
	if parent != node {
		parent = getParent(parents, parent)
		parents[node] = parent
	}

	return parent
}

func union(parents []int, node1, node2 int) {
	parents[node1] = node2
}
