/*
LeetCode 977: https://leetcode.com/problems/squares-of-a-sorted-array/
*/

package leetcode

import (
	"math"
)

func sortedSquares(A []int) []int {
	i := 1
	for i < len(A) && math.Abs(float64(A[i])) <= math.Abs(float64(A[i-1])) {
		i++
	}

	j := i - 1
	result := make([]int, 0)
	for j >= 0 && i < len(A) {
		s1, s2 := A[j]*A[j], A[i]*A[i]
		if s1 <= s2 {
			result = append(result, s1)
			j--
		} else {
			result = append(result, s2)
			i++
		}
	}

	for j >= 0 {
		result = append(result, A[j]*A[j])
		j--
	}

	for i < len(A) {
		result = append(result, A[i]*A[i])
		i++
	}

	return result
}
