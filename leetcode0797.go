/*
LeetCode 797: https://leetcode.com/problems/all-paths-from-source-to-target/description/
*/

package leetcode

func allPathsSourceTarget(graph [][]int) [][]int {
	return helper797(graph, 0, len(graph)-1, []int{})
}

func helper797(graph [][]int, src, dst int, path []int) [][]int {
	paths := make([][]int, 0)
	path = append(path, src)
	if src == dst {
		clone := make([]int, len(path))
		copy(clone, path)
		paths = append(paths, clone)
	} else {
		for _, next := range graph[src] {
			more := helper797(graph, next, dst, path)
			paths = append(paths, more...)
		}
	}

	return paths
}
