/*
LeetCode 70: https://leetcode.com/problems/climbing-stairs/
*/

package leetcode

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	var dp [2]int
	dp[0], dp[1] = 1, 2
	for i := 2; i < n; i++ {
		dp[i%2] = dp[i%2] + dp[(i-1)%2]
	}

	return dp[(n-1)%2]
}
