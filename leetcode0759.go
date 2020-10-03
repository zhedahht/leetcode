/*
LeetCode 759: https://leetcode.com/problems/employee-free-time/
*/

package leetcode

/**
 * Definition for an Interval.
 * type Interval struct {
 *     Start int
 *     End   int
 * }
 */

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	union := func(s1, s2 []*Interval) []*Interval {
		if len(s1) == 0 {
			return s2
		}

		if len(s2) == 0 {
			return s1
		}

		i, j := 0, 0
		result := make([]*Interval, 0)
		var prev, next *Interval
		for i < len(s1) || j < len(s2) {
			if i == len(s1) || (j < len(s2) && s1[i].Start > s2[j].Start) {
				next = s2[j]
				j++
			} else {
				next = s1[i]
				i++
			}

			if prev == nil {
				prev = next
			} else if prev.End < next.Start {
				result = append(result, prev)
				prev = next
			} else if prev.End < next.End {
				prev.End = next.End
			}
		}

		if prev != nil {
			result = append(result, prev)
		}

		return result
	}

	occupied := make([]*Interval, 0)
	for _, s := range schedule {
		occupied = union(occupied, s)
	}

	result := make([]*Interval, 0)
	for i := 0; i < len(occupied)-1; i++ {
		free := &Interval{
			Start: occupied[i].End,
			End:   occupied[i+1].Start,
		}

		result = append(result, free)
	}

	return result
}
