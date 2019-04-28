/*
LeetCode 62: https://leetcode.com/problems/unique-paths/
*/

package leetcode

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[n-1]
}
