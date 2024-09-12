/*
LeetCode 436: https://leetcode.com/problems/find-right-interval/description/
*/

package leetcode

import "sort"

func findRightInterval(intervals [][]int) []int {
	startAndIndex := make([][2]int, len(intervals))
	for i, interval := range intervals {
		startAndIndex[i] = [2]int{interval[0], i}
	}

	sort.Slice(startAndIndex, func(i, j int) bool {
		return startAndIndex[i][0] < startAndIndex[j][0]
	})

	result := make([]int, len(intervals))
	for i, interval := range intervals {
		start, end := 0, len(startAndIndex)-1
		target := interval[1]
		for start <= end {
			mid := (start + end) / 2
			midStart, midIndex := startAndIndex[mid][0], startAndIndex[mid][1]
			if midStart < target {
				start = mid + 1
			} else {
				if mid == 0 || startAndIndex[mid-1][0] < target {
					result[i] = midIndex
					break
				} else {
					end = mid - 1
				}
			}
		}

		if start > end {
			result[i] = -1
		}
	}

	return result
}
