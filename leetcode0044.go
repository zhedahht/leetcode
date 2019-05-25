/*
LeetCode 44: https://leetcode.com/problems/wildcard-matching/
*/

package leetcode

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	for j := 0; j < len(p); j++ {
		if p[j] == '*' {
			dp[0][j+1] = dp[0][j]
		}
	}

	for i := 0; i < len(s); i++ {
		dp[i+1][0] = false
		for j := 0; j < len(p); j++ {
			if p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j]
			} else if s[i] == p[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = false
			}
		}
	}

	return dp[len(s)][len(p)]
}
