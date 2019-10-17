/*
LeetCode 122: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
*/

package leetcode

// NOTE: the function name should be maxProfit. Rename it to avoid conflicts.
func maxProfit122(prices []int) int {
	result := 0
	if len(prices) == 0 {
		return result
	}

	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prev := prices[i-1]; prices[i] <= prev {
			result += prev - min
			min = prices[i]
		} else if i == len(prices)-1 {
			result += prices[i] - min
		}
	}

	return result
}
