/*
LeetCode 253: https://leetcode.com/problems/meeting-rooms-ii/description/
*/

package leetcode

import (
	"container/heap"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)
	for _, interval := range intervals {
		if h.Len() == 0 {
			heap.Push(h, interval[1])
		} else {
			top := (*h)[0]
			if top <= interval[0] {
				heap.Pop(h)
				heap.Push(h, interval[1])
			} else {
				heap.Push(h, interval[1])
			}
		}
	}

	return h.Len()
}
