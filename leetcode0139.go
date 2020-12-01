/*
LeetCode 139: https://leetcode.com/problems/word-break/
*/

package leetcode

// NOTE: The name should be wordBreak. Rename it to avoid conflicts
func wordBreak139(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := range s {
		for _, word := range wordDict {
			if i+1 >= len(word) && s[i+1-len(word):i+1] == word {
				dp[i+1] = dp[i+1] || dp[i+1-len(word)]
				if dp[i+1] {
					break
				}
			}
		}
	}

	return dp[len(s)]
}
