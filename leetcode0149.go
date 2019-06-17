/*
LeetCode 149: https://leetcode.com/problems/max-points-on-a-line/
*/

package leetcode

import "math"

func maxPoints(points [][]int) int {
	result := 0
	for i, p1 := range points {
		max, sameP, sameX := 0, 1, 0
		slopeCounts := make(map[float64]int)
		for j := 0; j < i; j++ {
			p2 := points[j]
			if p1[0] == p2[0] && p1[1] == p2[1] {
				sameP++
			} else if p1[0] == p2[0] {
				sameX++
			} else {
				slope := float64(p1[1]-p2[1]) / float64(p1[0]-p2[0])
				if _, exists := slopeCounts[slope]; exists {
					slopeCounts[slope]++
				} else {
					slopeCounts[slope] = 1
				}
			}
		}

		max = sameX
		for _, count := range slopeCounts {
			max = int(math.Max(float64(max), float64(count)))
		}

		max += sameP
		result = int(math.Max(float64(result), float64(max)))
	}

	return result
}
