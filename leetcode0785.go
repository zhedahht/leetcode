/*
LeetCode 785: https://leetcode.com/problems/is-graph-bipartite/description/
*/

package leetcode

func isBipartite(graph [][]int) bool {
	values := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if values[i] == 0 {
			if !isBipartiteCore(graph, values, i) {
				return false
			}
		}
	}

	return true
}

func isBipartiteCore(graph [][]int, values []int, i int) bool {
	// values[i]: 0 -> not visited, 1 -> black, 2 -> white
	queue := [][2]int{{i, 1}}
	for len(queue) > 0 {
		node, color := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, next := range graph[node] {
			if values[next] == 0 {
				values[next] = 3 - color
				queue = append(queue, [2]int{next, 3 - color})
			} else if values[next] != 3-color {
				return false
			}
		}
	}

	return true
}
