/*
LeetCode 643: https://leetcode.com/problems/maximum-average-subarray-i/description/
*/

package leetcode

import "math"

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	maxSum := math.MinInt
	for i := 0; i < k-1; i++ {
		sum += nums[i]
	}

	for i := k - 1; i < len(nums); i++ {
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}

		sum -= nums[i-k+1]
	}

	return float64(maxSum) / float64(k)
}
