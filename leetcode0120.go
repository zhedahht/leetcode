/*
LeetCode 120: https://leetcode.com/problems/triangle/
*/

package leetcode

func minimumTotal(triangle [][]int) int {
	const max = 0x7fffffff
	min := func(m, n int) int {
		if m < n {
			return m
		}

		return n
	}

	if len(triangle) == 0 {
		return 0
	}

	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			k := max
			if j < i {
				k = min(k, dp[j])
			}

			if j >= 1 {
				k = min(k, dp[j-1])
			}

			dp[j] = k + triangle[i][j]
		}
	}

	k := max
	for i := 0; i < len(triangle); i++ {
		k = min(k, dp[i])
	}

	return k
}
