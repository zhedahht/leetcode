/*
LeetCode 228: https://leetcode.com/problems/summary-ranges/
*/

package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	start := -1
	end := -1
	for i, num := range nums {
		if i > 0 && num == nums[i-1]+1 {
			end = i
		} else {
			if i > 0 {
				result = appendRanges(result, nums, start, end)
			}

			start = i
			end = i
		}
	}

	if end >= 0 {
		result = appendRanges(result, nums, start, end)
	}

	return result
}

func appendRanges(ranges []string, nums []int, start int, end int) []string {
	if end > start {
		ranges = append(ranges, fmt.Sprintf("%d->%d", nums[start], nums[end]))
	} else {
		ranges = append(ranges, fmt.Sprintf("%d", nums[start]))
	}

	return ranges
}
