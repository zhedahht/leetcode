/*
LeetCode 132: https://leetcode.com/problems/palindrome-partitioning-ii/
*/

package leetcode

func minCut(s string) int {
	isPalindrome := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		isPalindrome[i] = make([]bool, len(s))
	}

	for i := len(s); i >= 1; i-- {
		for j := 0; j < i; j++ {
			row, col, length := j, len(s)-i+j, len(s)-i+1
			if length == 1 || ((length == 2 || isPalindrome[row+1][col-1]) && s[row] == s[col]) {
				isPalindrome[row][col] = true
			}
		}
	}

	dp := make([]int, len(s)+1)
	for i := 1; i < len(s); i++ {
		if isPalindrome[0][i] {
			dp[i+1] = 0
		} else {
			dp[i+1] = len(s)
			for j := 0; j < i; j++ {
				if isPalindrome[j+1][i] && dp[i+1] > dp[j+1]+1 {
					dp[i+1] = dp[j+1] + 1
				}
			}
		}
	}

	return dp[len(s)]
}
