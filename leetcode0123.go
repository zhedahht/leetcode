/*
LeetCode 123: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
*/

package leetcode

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dp1, dp2 := make([]int, len(prices)), make([]int, len(prices))
	min, maxLeft := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > min && prices[i]-min > maxLeft {
			dp1[i] = prices[i] - min
			maxLeft = dp1[i]
		} else {
			dp1[i] = maxLeft
		}

		if prices[i] < min {
			min = prices[i]
		}
	}

	max, maxRight := prices[len(prices)-1], 0
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] < max && max-prices[i] > maxRight {
			dp2[i+1] = max - prices[i]
			maxRight = dp2[i+1]
		} else {
			dp2[i+1] = maxRight
		}

		if prices[i] > max {
			max = prices[i]
		}
	}

	max = dp1[len(dp1)-1]
	for i := 0; i < len(dp1)-1; i++ {
		if dp1[i]+dp2[i+1] > max {
			max = dp1[i] + dp2[i+1]
		}
	}

	return max
}
