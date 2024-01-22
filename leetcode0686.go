/*
LeetCode 686: https://leetcode.com/problems/repeated-string-match/description/
*/

package leetcode

import "strings"

func repeatString(a string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result = result + a
	}

	return result
}

func repeatedStringMatch(a string, b string) int {
	lenA, lenB := len(a), len(b)
	times := lenB / lenA

	newA := repeatString(a, times)
	for _, t := range []int{times, times + 1, times + 2} {
		if strings.Contains(newA, b) {
			return t
		}

		newA = newA + a
	}

	return -1
}
