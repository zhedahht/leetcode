/*
LeetCode 949: https://leetcode.com/problems/largest-time-for-given-digits/
*/

package leetcode

import "strings"

func largestTimeFromDigits(A []int) string {
	return permuteInternal(A, 0, "")
}

func permuteInternal(A []int, i int, max string) string {
	if i == len(A) {
		newTime, ok := isValidTime(A)
		if ok {
			if strings.Compare(newTime, max) > 0 {
				max = newTime
			}
		}
	} else {
		for j := i; j < len(A); j++ {
			A[i], A[j] = A[j], A[i]
			max = permuteInternal(A, i+1, max)
			A[i], A[j] = A[j], A[i]
		}
	}

	return max
}

func isValidTime(A []int) (string, bool) {
	if A[0] > 2 || (A[0] == 2 && A[1] > 3) || A[2] > 5 {
		return "", false
	}

	digits := make([]byte, 0)
	for _, num := range A {
		digits = append(digits, byte(num+'0'))
	}

	time := string(digits)
	return time[:2] + ":" + time[2:], true
}
