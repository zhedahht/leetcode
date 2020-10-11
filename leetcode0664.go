/*
LeetCode 664: https://leetcode.com/problems/strange-printer/
*/

package leetcode

func strangePrinter(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}

	for i := range s {
		dp[i][i] = 1
	}

	for le := 2; le <= len(s); le++ {
		for i := 0; i <= len(s)-le; i++ {
			j := i + le - 1
			dp[i][j] = le
			for k := i + 1; k <= j; k++ {
				turns := dp[i][k-1] + dp[k][j]
				if s[k-1] == s[j] {
					turns--
				}

				if dp[i][j] > turns {
					dp[i][j] = turns
				}
			}
		}
	}

	return dp[0][len(s)-1]
}
