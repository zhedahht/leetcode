/*
LeetCode 474: https://leetcode.com/problems/ones-and-zeroes/
*/

package leetcode

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for j := range dp {
		dp[j] = make([]int, n+1)
	}

	for _, str := range strs {
		zero, one := 0, 0
		for _, b := range str {
			if b == '0' {
				zero++
			} else {
				one++
			}
		}

		for j := m; j >= 0; j-- {
			for k := n; k >= 0; k-- {
				if j >= zero && k >= one && dp[j-zero][k-one]+1 > dp[j][k] {
					dp[j][k] = dp[j-zero][k-one] + 1
				}
			}
		}
	}

	return dp[m][n]
}
