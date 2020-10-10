/*
LeetCode 322: https://leetcode.com/problems/coin-change/
*/

package leetcode

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if i >= c && dp[i-c] >= 0 {
				if dp[i] == -1 {
					dp[i] = dp[i-c] + 1
				} else if dp[i] > dp[i-c]+1 {
					dp[i] = dp[i-c] + 1
				}
			}
		}
	}

	return dp[amount]
}
