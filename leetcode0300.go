/*
LeetCode 300: https://leetcode.com/problems/longest-increasing-subsequence/submissions/
*/

package leetcode

import "math"

func lengthOfLIS(nums []int) int {
	lengths := make([]int, len(nums))
	max := 0
	for i := 0; i < len(nums); i++ {
		lengths[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				lengths[i] = int(math.Max(float64(lengths[i]), float64(lengths[j]+1)))
			}
		}

		max = int(math.Max(float64(max), float64(lengths[i])))
	}

	return max
}
