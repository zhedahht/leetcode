/*
LeetCode 665: https://leetcode.com/problems/non-decreasing-array/description/
*/

package leetcode

func checkPossibility(nums []int) bool {
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] <= nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	maxLength := 0
	for _, v := range dp {
		maxLength = max(maxLength, v)
	}

	return maxLength >= len(nums)-1
}
