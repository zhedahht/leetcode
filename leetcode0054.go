/*
LeetCode 54: https://leetcode.com/problems/spiral-matrix/
*/

package leetcode

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}

	rows, cols := len(matrix), len(matrix[0])
	for i := 0; 2*i < rows && 2*i < cols; i++ {
		row, col := i, i
		for ; col < cols-i; col++ {
			result = append(result, matrix[row][col])
		}

		row, col = row+1, col-1
		for ; row < rows-i; row++ {
			result = append(result, matrix[row][col])
		}

		row, col = row-1, col-1
		for ; row > i && col >= i; col-- {
			result = append(result, matrix[row][col])
		}

		row, col = row-1, col+1
		for ; 2*i+1 < cols && row > i; row-- {
			result = append(result, matrix[row][col])
		}
	}

	return result
}
