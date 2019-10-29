/*
LeetCode 868: https://leetcode.com/problems/binary-gap/
*/

package leetcode

import "math"

func binaryGap(N int) int {
	result, gap := 0, 0
	found1 := false
	for N > 0 {
		digit := N & 1
		N = N >> 1

		if found1 {
			if digit == 1 {
				result = int(math.Max(float64(result), float64(gap+1)))
				gap = 0
			} else {
				gap++
			}
		} else if digit == 1 {
			found1 = true
		}
	}

	return result
}
