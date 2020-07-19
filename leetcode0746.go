/*
LeetCode 746: https://leetcode.com/problems/min-cost-climbing-stairs/
*/

package leetcode

import "math"

func minCostClimbingStairs(cost []int) int {
	min := func(a, b int) int {
		return int(math.Min(float64(a), float64(b)))
	}

	dp := make([]int, 2)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i%2] = cost[i] + min(dp[0], dp[1])
	}

	return min(dp[0], dp[1])
}
