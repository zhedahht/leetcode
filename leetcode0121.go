/*
LeetCode 121: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
*/

package leetcode

// NOTE: the function name should be maxProfit. Rename it to avoid conflicts.
func maxProfit121(prices []int) int {
	const maxInt = 0x7fffffff
	const minInt = 0

	min := func(m, n int) int {
		if m < n {
			return m
		}

		return n
	}

	max := func(m, n int) int {
		if m > n {
			return m
		}

		return n
	}

	result := minInt
	minP := maxInt
	for _, p := range prices {
		minP = min(minP, p)
		if p > minP {
			result = max(result, p-minP)
		}
	}

	return result
}
