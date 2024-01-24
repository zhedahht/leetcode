/*
LeetCode 697: https://leetcode.com/problems/degree-of-an-array/
*/

package leetcode

import "math"

func findShortestSubArray(nums []int) int {
	counts := make(map[int]int)
	firsts, lasts := make(map[int]int), make(map[int]int)
	for i, num := range nums {
		counts[num]++

		lasts[num] = i
		if _, exists := firsts[num]; !exists {
			firsts[num] = i
		}
	}

	degree := 0
	minLength := math.MaxInt
	for num, count := range counts {
		if count > degree {
			degree = count
			length := lasts[num] - firsts[num] + 1
			minLength = length
		} else if count == degree {
			length := lasts[num] - firsts[num] + 1
			minLength = min(minLength, length)
		}
	}

	return minLength
}
