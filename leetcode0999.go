/*
LeetCode 999: https://leetcode.com/problems/available-captures-for-rook/
*/

package leetcode

func numRookCaptures(board [][]byte) int {
	helper := func(i, j int) int {
		dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		count := 0
		for _, dir := range dirs {
			r, c := i, j
			for true {
				r, c = r+dir[0], c+dir[1]
				if r < 0 || r >= 8 || c < 0 || c >= 8 ||
					board[r][c] == 'B' {
					break
				}
				if board[r][c] == 'p' {
					count++
					break
				}
			}
		}
		return count
	}

	for i, row := range board {
		for j, v := range row {
			if v == 'R' {
				return helper(i, j)
			}
		}
	}

	return 0
}
