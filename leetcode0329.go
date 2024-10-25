/*
LeetCode 329: https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description
*/

package leetcode

func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range matrix {
		dp[i] = make([]int, len(matrix[0]))
	}

	result := 0
	for i, row := range matrix {
		for j := range row {
			l := helper329(matrix, dp, i, j)
			result = max(result, l)
		}
	}

	return result
}

func helper329(matrix, dp [][]int, i, j int) int {
	if dp[i][j] > 0 {
		return dp[i][j]
	}

	rows, cols := len(matrix), len(matrix[0])
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	result := 1
	for _, dir := range dirs {
		next := []int{i + dir[0], j + dir[1]}
		if next[0] >= 0 && next[0] < rows && next[1] >= 0 && next[1] < cols &&
			matrix[i][j] < matrix[next[0]][next[1]] {
			l := helper329(matrix, dp, next[0], next[1])
			result = max(result, 1+l)
		}
	}

	dp[i][j] = result
	return result
}
