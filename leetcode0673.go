/*
LeetCode 673: https://leetcode.com/problems/number-of-longest-increasing-subsequence/description/
*/

package leetcode

func findNumberOfLIS(nums []int) int {
	lengths := make([]int, len(nums))
	counts := make([]int, len(nums))
	for i := range nums {
		lengths[i] = 1
		counts[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if lengths[j]+1 > lengths[i] {
					lengths[i] = lengths[j] + 1
					counts[i] = counts[j]
				} else if lengths[j]+1 == lengths[i] {
					counts[i] += counts[j]
				}
			}
		}
	}

	maxLength := 0
	count := 0
	for i, length := range lengths {
		if length > maxLength {
			maxLength = length
			count = counts[i]
		} else if length == maxLength {
			count += counts[i]
		}
	}

	return count
}
