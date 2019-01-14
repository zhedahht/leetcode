/*
LeetCode 695: https://leetcode.com/problems/max-area-of-island/
*/

package leetcode

func maxAreaOfIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	visited := make([]bool, rows*cols)
	max := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 && !visited[i*cols+j] {
				area := getIslandArea(grid, visited, i, j)
				if area > max {
					max = area
				}
			}
		}
	}

	return max
}

func getIslandArea(grid [][]int, visited []bool, i int, j int) int {
	rows, cols := len(grid), len(grid[0])
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := make([][]int, 0)
	queue = append(queue, []int{i, j})
	visited[i*cols+j] = true
	area := 0
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		area++

		for _, dir := range dirs {
			row := pos[0] + dir[0]
			col := pos[1] + dir[1]
			if row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] == 1 && !visited[row*cols+col] {
				queue = append(queue, []int{row, col})
				visited[row*cols+col] = true
			}
		}
	}

	return area
}
