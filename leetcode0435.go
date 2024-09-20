/*
LeetCode 435: https://leetcode.com/problems/non-overlapping-intervals/description/
*/

package leetcode

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	prev := 0
	count := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[prev][1] {
			count++
			prev = i
		}
	}

	return len(intervals) - count
}
