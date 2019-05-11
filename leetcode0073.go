/*
LeetCode 73: https://leetcode.com/problems/set-matrix-zeroes/
*/

package leetcode

func setZeroes(matrix [][]int) {
	var firstRow, firstCol bool
	for _, num := range matrix[0] {
		if num == 0 {
			firstRow = true
			break
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstCol = true
			break
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j], matrix[i][0] = 0, 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRow {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if firstCol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
