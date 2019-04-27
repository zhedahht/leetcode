/*
LeetCode 96: https://leetcode.com/problems/unique-binary-search-trees/
*/

package leetcode

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		num := 0
		for j := 0; j <= i-1; j++ {
			num += dp[j] * dp[i-1-j]
		}

		dp[i] = num
	}

	return dp[n]
}
