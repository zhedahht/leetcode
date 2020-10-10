/*
LeetCode 494: https://leetcode.com/problems/target-sum/
*/

package leetcode

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if (sum+S)%2 == 1 || S > sum {
		return 0
	}

	sum = (sum + S) / 2
	dp := make([]int, sum+1)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := sum; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}

	return dp[sum]
}
