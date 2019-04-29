/*
LeetCode 79: https://leetcode.com/problems/word-search/
*/

package leetcode

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	var found bool
	for i := 0; i < len(board) && !found; i++ {
		for j := 0; j < len(board[0]) && !found; j++ {
			found = helper79(board, word, visited, i, j, 0)
		}
	}

	return found
}

func helper79(board [][]byte, word string, visited [][]bool, row, col, index int) bool {
	if index == len(word) {
		return true
	}

	rows, cols := len(board), len(board[0])
	if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[index] || visited[row][col] {
		return false
	}

	visited[row][col] = true
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var found bool
	for i := 0; !found && i < len(dirs); i++ {
		nextRow, nextCol := dirs[i][0]+row, dirs[i][1]+col
		found = helper79(board, word, visited, nextRow, nextCol, index+1)
	}

	visited[row][col] = false
	return found
}
