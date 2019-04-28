/*
LeetCode 56: https://leetcode.com/problems/merge-intervals/
*/

package leetcode

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	result := make([][]int, 0)
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > cur[1] {
			result = append(result, cur)
			cur = intervals[i]
		} else {
			cur[1] = int(math.Max(float64(cur[1]), float64(intervals[i][1])))
		}
	}

	result = append(result, cur)
	return result
}
