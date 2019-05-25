/*
LeetCode 87: https://leetcode.com/problems/scramble-string/
*/

package leetcode

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	le := len(s1)
	dp := make([][][]bool, le+1)
	for i := 0; i <= le; i++ {
		dp[i] = make([][]bool, le)
		for j := 0; j < le; j++ {
			dp[i][j] = make([]bool, le)
		}
	}

	dp[0][0][0] = true
	for i := 1; i <= le; i++ {
		for j := 0; i+j <= le; j++ {
			for k := 0; i+k <= le; k++ {
				if s1[j:i+j] == s2[k:i+k] {
					dp[i][j][k] = true
				} else {
					for n := 1; n < i && !dp[i][j][k]; n++ {
						dp[i][j][k] = dp[i][j][k] || (dp[n][j][i+k-n] && dp[i-n][j+n][k]) || (dp[n][j][k] && dp[i-n][j+n][k+n])
					}
				}
			}
		}
	}

	return dp[le][0][0]
}
