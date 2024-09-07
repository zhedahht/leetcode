/*
LeetCode 463: https://leetcode.com/problems/island-perimeter/description/
*/

package leetcode

func islandPerimeter(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	result := 0
	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				dirs := [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
				for _, dir := range dirs {
					x := i + dir[0]
					y := j + dir[1]
					if x < 0 || x == rows || y < 0 || y == cols || grid[x][y] == 0 {
						result++
					}
				}
			}
		}
	}

	return result
}
