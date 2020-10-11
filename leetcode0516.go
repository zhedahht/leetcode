/*
LeetCode 516: https://leetcode.com/problems/longest-palindromic-subsequence/
*/

package leetcode

func longestPalindromeSubseq(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
		}
	}

	for le := 3; le <= len(s); le++ {
		for i := 0; i <= len(s)-le; i++ {
			j := i + le - 1
			if s[i] == s[j] {
				dp[i][j] += dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][len(s)-1]
}
