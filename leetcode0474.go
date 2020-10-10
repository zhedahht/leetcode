/*
LeetCode 474: https://leetcode.com/problems/ones-and-zeroes/
*/

package leetcode

func findMaxForm(strs []string, m int, n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	dp := make([][][]int, 2)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i, str := range strs {
		zero, one := 0, 0
		for _, b := range str {
			if b == '0' {
				zero++
			} else {
				one++
			}
		}

		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[(i+1)%2][j][k] = dp[i%2][j][k]
				if j >= zero && k >= one {
					dp[(i+1)%2][j][k] = max(dp[(i+1)%2][j][k], dp[i%2][j-zero][k-one]+1)
				}
			}
		}
	}

	return dp[len(strs)%2][m][n]
}
