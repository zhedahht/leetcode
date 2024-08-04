/*
LeetCode 419: https://leetcode.com/problems/battleships-in-a-board/description/
*/

package leetcode

func countBattleships(board [][]byte) int {
	count := 0
	rows, cols := len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'X' && !(i > 0 && board[i-1][j] == 'X') &&
				!(j > 0 && board[i][j-1] == 'X') {
				count++
			}
		}
	}

	return count
}
