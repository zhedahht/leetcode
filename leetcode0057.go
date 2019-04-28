/*
LeetCode 57: https://leetcode.com/problems/insert-interval/
*/

package leetcode

import (
	"math"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	var overlapping []int
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		if overlapping == nil {
			overlapping = make([]int, 2)
			overlapping[0] = intervals[i][0]
		}

		overlapping[1] = intervals[i][1]
		i++
	}

	if overlapping == nil {
		result = append(result, newInterval)
	} else {
		p0 := int(math.Min(float64(overlapping[0]), float64(newInterval[0])))
		p1 := int(math.Max(float64(overlapping[1]), float64(newInterval[1])))
		result = append(result, []int{p0, p1})
	}

	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}
