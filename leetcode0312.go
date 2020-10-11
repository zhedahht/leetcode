/*
LeetCode 312: https://leetcode.com/problems/burst-balloons/
*/

package leetcode

func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}

	for le := 1; le <= len(nums); le++ {
		for i := 0; i <= len(nums)-le; i++ {
			j := i + le - 1
			for k := i; k <= j; k++ {
				left, right := 1, 1
				if i > 0 {
					left = nums[i-1]
				}

				if j < len(nums)-1 {
					right = nums[j+1]
				}

				coins := left * nums[k] * right
				if k != i {
					coins += dp[i][k-1]
				}

				if k != j {
					coins += dp[k+1][j]
				}

				if coins > dp[i][j] {
					dp[i][j] = coins
				}
			}
		}
	}

	return dp[0][len(nums)-1]
}
