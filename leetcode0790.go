/*
LeetCode 790: https://leetcode.com/problems/domino-and-tromino-tiling/description/
*/

package leetcode

func numTilings(n int) int {
	dp := []int{1, 1, 2}
	for i := 3; i <= n; i++ {
		val := dp[i-1] + dp[i-2]
		for j := i - 3; j >= 0; j = j - 2 {
			val += dp[j] * 2
		}

		for j := i - 4; j >= 0; j = j - 2 {
			val += dp[j] * 2
		}

		val = val % 1000000007
		dp = append(dp, val)
	}

	return dp[n]
}
