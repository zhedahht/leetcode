/*
LeetCode 200: https://leetcode.com/problems/number-of-islands/
*/

package leetcode

func numIslands(grid [][]byte) int {
	fill := func(grid [][]byte, row, col int) {
		queue := make([][]int, 0)
		queue = append(queue, []int{row, col})
		grid[row][col] = '0'
		dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for len(queue) > 0 {
			pos := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				r, c := pos[0]+dir[0], pos[1]+dir[1]
				if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == '1' {
					queue = append(queue, []int{r, c})
					grid[r][c] = '0'
				}
			}
		}
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				result++
				fill(grid, i, j)
			}
		}
	}

	return result
}
