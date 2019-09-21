/*
LeetCode 198: https://leetcode.com/problems/house-robber/
*/

package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	max := func(m, n int) int {
		if m > n {
			return m
		}

		return n
	}

	dp := []int{nums[0], max(nums[0], nums[1])}
	for i := 2; i < len(nums); i++ {
		dp[i%2] = max(dp[(i-1)%2], nums[i]+dp[i%2])
	}

	return dp[(len(nums)-1)%2]
}
