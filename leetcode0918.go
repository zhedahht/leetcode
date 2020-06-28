/*
LeetCode 918: https://leetcode.com/problems/maximum-sum-circular-subarray/
*/

package leetcode

import "math"

func maxSubarraySumCircular(A []int) int {
	total, sumMax, sumMin := 0, 0, 0
	max, min := A[0], A[0]
	for _, num := range A {
		total += num

		if sumMax <= 0 {
			sumMax = num
		} else {
			sumMax += num
		}

		if max < sumMax {
			max = sumMax
		}

		if sumMin >= 0 {
			sumMin = num
		} else {
			sumMin += num
		}

		if min > sumMin {
			min = sumMin
		}
	}

	if max < 0 {
		return max
	}

	return int(math.Max(float64(max), float64(total-min)))
}
