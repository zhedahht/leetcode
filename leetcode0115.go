/*
LeetCode 115: https://leetcode.com/problems/distinct-subsequences/
*/

package leetcode

func numDistinct(s string, t string) int {
	dp := make([]int, len(s)+1)
	for i := range dp {
		dp[i] = 1
	}

	for i := range t {
		prev := dp[0]
		dp[0] = 0
		for j := range s {
			cur := dp[j]
			if t[i] == s[j] {
				cur += prev
			}

			prev, dp[j+1] = dp[j+1], cur
		}
	}

	return dp[len(s)]
}
