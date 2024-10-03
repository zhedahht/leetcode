/*
LeetCode 252: https://leetcode.com/problems/meeting-rooms/description/
*/

package leetcode

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] > intervals[i][0] {
			return false
		}
	}

	return true
}
