/*
LeetCode 529: https://leetcode.com/problems/minesweeper/
*/

package leetcode

func updateBoard(board [][]byte, click []int) [][]byte {
	rows, cols := len(board), len(board[0])

	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	queue := make([][]int, 0)
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	queue = append(queue, click)
	visited[click[0]][click[1]] = true
	dirs := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		count := 0
		empties := make([][]int, 0)
		for _, dir := range dirs {
			r, c := pos[0]+dir[0], pos[1]+dir[1]
			if r >= 0 && r < rows && c >= 0 && c < cols {
				if board[r][c] == 'E' && !visited[r][c] {
					empties = append(empties, []int{r, c})
				} else if board[r][c] == 'M' {
					count++
				}
			}
		}

		if count == 0 {
			board[pos[0]][pos[1]] = 'B'
			for _, empty := range empties {
				visited[empty[0]][empty[1]] = true
				queue = append(queue, empty)
			}
		} else {
			board[pos[0]][pos[1]] = '0' + byte(count)
		}
	}

	return board
}
