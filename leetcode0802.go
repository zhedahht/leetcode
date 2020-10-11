/*
LeetCode 802: https://leetcode.com/problems/find-eventual-safe-states/
*/

package leetcode

func eventualSafeNodes(graph [][]int) []int {
	circle := make([]bool, len(graph))
	color := make([]int, len(graph))

	for n := 0; n < len(graph); n++ {
		helper802(graph, n, circle, color)
	}

	result := make([]int, 0)
	for n, c := range circle {
		if !c {
			result = append(result, n)
		}
	}

	return result
}

func helper802(graph [][]int, n int, circle []bool, color []int) {
	if color[n] == 1 {
		circle[n] = true
		return
	}

	if color[n] == 2 {
		return
	}

	color[n] = 1
	for _, neighbor := range graph[n] {
		helper802(graph, neighbor, circle, color)
		if circle[neighbor] {
			circle[n] = true
		}
	}

	color[n] = 2
}
