/*
LeetCode 416: https://leetcode.com/problems/partition-equal-subset-sum/
*/

package leetcode

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	sum = sum / 2
	dp := make([]bool, sum+1)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := sum; j >= 0; j-- {
			dp[j] = dp[j] || (j >= nums[i] && dp[j-nums[i]])
		}

		dp[0] = false
	}

	return dp[sum]
}
