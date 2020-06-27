/*
LeetCode 983: https://leetcode.com/problems/minimum-cost-for-tickets/
*/

package leetcode

import "math"

func mincostTickets(days []int, costs []int) int {
	min := func(a, b int) int {
		return int(math.Min(float64(a), float64(b)))
	}

	max := func(a, b int) int {
		return int(math.Max(float64(a), float64(b)))
	}

	dp := make([]int, len(days)+1)
	for i := range days {
		dp[i+1] = dp[i] + costs[0]
		for j := max(0, i-6); j <= i; j++ {
			if days[i]-days[j] < 7 {
				dp[i+1] = min(dp[i+1], dp[j]+costs[1])
				break
			}
		}

		for j := max(0, i-29); j <= i; j++ {
			if days[i]-days[j] < 30 {
				dp[i+1] = min(dp[i+1], dp[j]+costs[2])
				break
			}
		}
	}

	return dp[len(days)]
}
