/*
LeetCode 538: https://leetcode.com/problems/minimum-time-difference/
*/

package leetcode

import "strconv"

func findMinDifference(timePoints []string) int {
	getInt := func(time string) int {
		hour, _ := strconv.Atoi(time[:2])
		minute, _ := strconv.Atoi(time[3:])
		return hour*60 + minute
	}

	counts := make([]int, 1440)
	for _, t := range timePoints {
		counts[getInt(t)]++
	}

	minDiff := 1440
	first, last, prev := -1, -1, -1
	for t, count := range counts {
		if count > 1 {
			return 0
		} else if count == 1 {
			if first < 0 {
				first = t
			} else {
				last = t
				diff := last - prev
				if minDiff > diff {
					minDiff = diff
				}
			}
			prev = t
		}
	}

	diff := first + 1440 - last
	if minDiff > diff {
		minDiff = diff
	}
	return minDiff
}
