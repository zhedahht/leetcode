/*
LeetCode 37: https://leetcode.com/problems/sudoku-solver/
*/

package leetcode

func solveSudoku(board [][]byte) {
	rows, cols, blocks := make([]map[byte]bool, 9), make([]map[byte]bool, 9), make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i], cols[i], blocks[i] = make(map[byte]bool), make(map[byte]bool), make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				blockIndex, n := (i/3)*3+j/3, board[i][j]
				rows[i][n], cols[j][n], blocks[blockIndex][n] = true, true, true
			}
		}
	}

	helper37(board, 0, 0, rows, cols, blocks)
}

func helper37(board [][]byte, i, j int, rows, cols, blocks []map[byte]bool) bool {
	if i == 9 {
		return true
	}

	nextI, nextJ := i, j+1
	if j == 8 {
		nextI, nextJ = i+1, 0
	}

	if board[i][j] != '.' {
		return helper37(board, nextI, nextJ, rows, cols, blocks)
	}

	done := false
	blockIndex := (i/3)*3 + j/3
	for n := byte('1'); n <= byte('9') && !done; n++ {
		_, existRows := rows[i][n]
		_, existCols := cols[j][n]
		_, existBlocks := blocks[blockIndex][n]
		if !existRows && !existCols && !existBlocks {
			rows[i][n], cols[j][n], blocks[blockIndex][n] = true, true, true
			board[i][j] = n
			done = helper37(board, nextI, nextJ, rows, cols, blocks)

			if done {
				return true
			}

			board[i][j] = '.'
			delete(rows[i], n)
			delete(cols[j], n)
			delete(blocks[blockIndex], n)
		}
	}

	return false
}
