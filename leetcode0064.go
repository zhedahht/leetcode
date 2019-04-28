/*
LeetCode 64: https://leetcode.com/problems/minimum-path-sum/
*/

package leetcode

import (
	"math"
)

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	dp := make([]int, cols)

	sum := 0
	for i := 0; i < cols; i++ {
		dp[i] = sum + grid[0][i]
		sum = dp[i]
	}

	for i := 1; i < rows; i++ {
		dp[0] += grid[i][0]

		for j := 1; j < cols; j++ {
			min := int(math.Min(float64(dp[j]), float64(dp[j-1])))
			dp[j] = grid[i][j] + min
		}
	}

	return dp[cols-1]
}
