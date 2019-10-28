/*
LeetCode 873: https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/
*/

package leetcode

import "math"

func lenLongestFibSubseq(A []int) int {
	result := 0
	prevToLens := make([]map[int]int, len(A))
	for i, num := range A {
		prevToLens[i] = make(map[int]int)
		for j := 0; j <= i-1; j++ {
			prev, lenJToI := num-A[j], 2
			if A[j] > prev {
				if len, exists := prevToLens[j][prev]; exists {
					lenJToI = len + 1
				}
			}

			prevToLens[i][A[j]] = lenJToI
			result = int(math.Max(float64(result), float64(lenJToI)))
		}
	}

	if result < 3 {
		return 0
	}

	return result
}
