/*
LeetCode 130: https://leetcode.com/problems/surrounded-regions/
*/

package leetcode

func solve(board [][]byte) {
	helper := func(board [][]byte, visited [][]bool, i, j int) [][]int {
		var result, queue [][]int
		dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		queue = append(queue, []int{i, j})
		visited[i][j] = true
		rows, cols := len(board), len(board[0])
		valid := true
		for len(queue) > 0 {
			pos := queue[0]
			queue = queue[1:]
			result = append(result, pos)
			if pos[0] == 0 || pos[0] == rows-1 || pos[1] == 0 || pos[1] == cols-1 {
				valid = false
			}

			for _, dir := range dirs {
				row, col := pos[0]+dir[0], pos[1]+dir[1]
				if row < 0 || row >= rows || col < 0 || col >= cols {
					continue
				}

				if board[row][col] == 'O' && !visited[row][col] {
					queue = append(queue, []int{row, col})
					visited[row][col] = true
				}
			}
		}

		if valid {
			return result
		}

		return result[0:0]
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	for i, row := range board {
		for j, cell := range row {
			if cell == 'O' && !visited[i][j] {
				pos := helper(board, visited, i, j)
				for _, p := range pos {
					board[p[0]][p[1]] = 'X'
				}
			}
		}
	}
}
