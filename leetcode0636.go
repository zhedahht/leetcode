/*
LeetCode 636: https://leetcode.com/problems/exclusive-time-of-functions/description/
*/

package leetcode

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	times := make([]int, n)
	threads := make([][3]int, 0) // thread Id, start time, duration
	prevEnd := -1

	for _, log := range logs {
		segs := strings.Split(log, ":")
		id, _ := strconv.Atoi(segs[0])
		timestamp, _ := strconv.Atoi(segs[2])
		if segs[1] == "start" {
			if len(threads) > 0 {
				if prevEnd < 0 {
					threads[len(threads)-1][2] = timestamp - threads[len(threads)-1][1]
				} else {
					threads[len(threads)-1][2] += timestamp - prevEnd
				}
			}
			threads = append(threads, [3]int{id, timestamp, -1})
			prevEnd = -1
		} else {
			prevThread := threads[len(threads)-1]
			if prevEnd < 0 {
				times[id] += timestamp + 1 - prevThread[1]
			} else {
				times[id] += timestamp + 1 - prevEnd + prevThread[2]
			}

			threads = threads[:len(threads)-1]
			prevEnd = timestamp + 1
		}
	}

	return times
}
