/*
LeetCode 317: https://leetcode.com/problems/shortest-distance-from-all-buildings/
*/

package leetcode

import "math"

// The function name should be shortestDistance.
// Rename the function name to resolve the name conflicts with other solutions.
func shortestDistance317(grid [][]int) int {
	pos := make([][]int, 0)
	for row, _ := range grid {
		for col, val := range grid[row] {
			if val == 1 {
				pos = append(pos, []int{row, col})
			}
		}
	}

	minDists := make([][]int, len(grid))
	visited := make([][]bool, len(grid))
	reached := make([][]bool, len(grid))
	for i, _ := range grid {
		minDists[i] = make([]int, len(grid[0]))
		visited[i] = make([]bool, len(grid[0]))
		reached[i] = make([]bool, len(grid[0]))
		for j, _ := range reached[i] {
			reached[i][j] = true
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, p := range pos {
		dist := 0
		queue1 := make([][]int, 0)
		queue2 := make([][]int, 0)
		for i, _ := range visited {
			for j, _ := range visited[i] {
				visited[i][j] = false
			}
		}

		queue1 = append(queue1, p)
		visited[p[0]][p[1]] = true

		for len(queue1) > 0 {
			cur := queue1[0]
			queue1 = queue1[1:]

			for _, dir := range dirs {
				row := cur[0] + dir[0]
				col := cur[1] + dir[1]
				if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) && !visited[row][col] && grid[row][col] == 0 {
					visited[row][col] = true
					minDists[row][col] += dist + 1
					queue2 = append(queue2, []int{row, col})
				}
			}

			if len(queue1) == 0 {
				dist++
				queue1 = queue2
				queue2 = make([][]int, 0)
			}
		}

		for i, _ := range reached {
			for j, _ := range reached[i] {
				reached[i][j] = reached[i][j] && visited[i][j]
			}
		}
	}

	minDist := 0x7FFFFFFF
	for row, _ := range minDists {
		for col, val := range minDists[row] {
			if grid[row][col] == 0 && reached[row][col] {
				minDist = int(math.Min(float64(minDist), float64(val)))
			}
		}
	}

	if minDist == 0x7FFFFFFF {
		return -1
	}

	return minDist
}
