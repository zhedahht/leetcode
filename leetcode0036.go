/*
LeetCode 36: https://leetcode.com/problems/valid-sudoku/
*/

package leetcode

func isValidSudoku(board [][]byte) bool {
	insertIfNotExist := func(set map[byte]bool, b byte) bool {
		if _, exist := set[b]; exist {
			return false
		}

		set[b] = true
		return true
	}

	for i := 0; i < 9; i++ {
		set := map[byte]bool{}
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' && !insertIfNotExist(set, board[i][j]) {
				return false
			}
		}
	}

	for j := 0; j < 9; j++ {
		set := map[byte]bool{}
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' && !insertIfNotExist(set, board[i][j]) {
				return false
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			set := map[byte]bool{}
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					row := i*3 + m
					col := j*3 + n
					if board[row][col] != '.' && !insertIfNotExist(set, board[row][col]) {
						return false
					}
				}
			}
		}
	}

	return true
}
