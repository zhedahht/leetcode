/*
LeetCode 674: https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/
*/

package leetcode

func findLengthOfLCIS(nums []int) int {
	length := 0
	max := 0
	for i := range nums {
		if i > 0 && nums[i] > nums[i-1] {
			length++
		} else {
			length = 1
		}

		if length > max {
			max = length
		}
	}

	return max
}
