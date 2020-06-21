/*
LeetCode 908: https://leetcode.com/problems/smallest-range-i/
*/

package leetcode

import "math"

func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]
	for _, v := range A {
		if min > v {
			min = v
		}

		if max < v {
			max = v
		}
	}

	return int(math.Max(0, float64(max-min-2*K)))
}
