/*
LeetCode 289: https://leetcode.com/problems/game-of-life/description/
*/

package leetcode

func gameOfLife(board [][]int) {
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	rows, cols := len(board), len(board[0])
	for i, row := range board {
		for j := range row {
			count := 0
			for k, dir := range dirs {
				next := []int{i + dir[0], j + dir[1]}
				if next[0] < 0 || next[0] == rows || next[1] < 0 || next[1] == cols {
					continue
				}

				v := board[next[0]][next[1]]
				if k < 4 {
					if v < 9 {
						v = 0
					} else {
						v = 1
					}
				}

				if v == 1 {
					count++
				}
			}

			if board[i][j] == 1 {
				board[i][j] = count + 9
			} else {
				board[i][j] = count
			}
		}
	}

	for i, row := range board {
		for j, count := range row {
			val := 0
			if count >= 9 {
				val = 1
				count -= 9
			}

			if val == 1 {
				if count == 2 || count == 3 {
					board[i][j] = 1
				} else {
					board[i][j] = 0
				}
			} else {
				if count == 3 {
					board[i][j] = 1
				} else {
					board[i][j] = 0
				}
			}
		}
	}
}
