/*
LeetCode 97: https://leetcode.com/problems/interleaving-string/
*/

package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true
	for i := range s1 {
		if s1[i] == s3[i] {
			dp[i+1][0] = true
		} else {
			break
		}
	}

	for i := range s2 {
		if s2[i] == s3[i] {
			dp[0][i+1] = true
		} else {
			break
		}
	}

	for i := range s1 {
		for j := range s2 {
			if s1[i] == s3[i+j+1] {
				dp[i+1][j+1] = dp[i][j+1]
			}

			if s2[j] == s3[i+j+1] {
				dp[i+1][j+1] = dp[i+1][j+1] || dp[i+1][j]
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
