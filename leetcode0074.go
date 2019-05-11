/*
LeetCode 74: https://leetcode.com/problems/search-a-2d-matrix/
*/

package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1
	for left <= right {
		mid := (left + right) / 2
		row, col := mid/cols, mid%cols
		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
