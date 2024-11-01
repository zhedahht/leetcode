/*
LeetCode 827: https://leetcode.com/problems/making-a-large-island/description/
*/

package leetcode

func largestIsland(grid [][]int) int {
	graphSizes := make(map[int]int)
	graphIndex := 2
	result := 0
	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				size := getGraphSize(grid, i, j, graphIndex)
				graphSizes[graphIndex] = size
				graphIndex++
				result = max(result, size)
			}
		}
	}

	rows, cols := len(grid), len(grid[0])
	dirs := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for i, row := range grid {
		for j, val := range row {
			if val == 0 {
				size := 1
				added := make(map[int]bool)
				for _, dir := range dirs {
					pos := []int{i + dir[0], j + dir[1]}
					if pos[0] >= 0 && pos[0] < rows && pos[1] >= 0 && pos[1] < cols &&
						grid[pos[0]][pos[1]] >= 2 {
						index := grid[pos[0]][pos[1]]
						if !added[index] {
							size += graphSizes[index]
							added[index] = true
						}
					}
				}

				result = max(result, size)
			}
		}
	}

	return result
}

func getGraphSize(grid [][]int, i, j int, graphIndex int) int {
	queue := [][]int{{i, j}}
	grid[i][j] = graphIndex
	size := 0
	rows, cols := len(grid), len(grid[0])
	dirs := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		size++

		for _, dir := range dirs {
			next := []int{pos[0] + dir[0], pos[1] + dir[1]}
			if next[0] >= 0 && next[0] < rows && next[1] >= 0 && next[1] < cols &&
				grid[next[0]][next[1]] == 1 {
				queue = append(queue, next)
				grid[next[0]][next[1]] = graphIndex
			}
		}
	}

	return size
}
