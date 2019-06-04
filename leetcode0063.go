/*
LeetCode 63: https://leetcode.com/problems/unique-paths-ii/
*/

package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	for i := range dp {
		if obstacleGrid[0][i] == 1 {
			dp[i] = 0
		} else if i > 0 {
			dp[i] = dp[i-1]
		} else {
			dp[i] = 1
		}
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[j] += dp[j-1]
			} else {
				dp[j] = 0
			}
		}
	}

	return dp[n-1]
}
