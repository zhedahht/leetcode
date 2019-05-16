/*
LeetCode 53: https://leetcode.com/problems/maximum-subarray/
*/

package leetcode

func maxSubArray(nums []int) int {
	sum, max := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if max < sum {
			max = sum
		}
	}

	return max
}
