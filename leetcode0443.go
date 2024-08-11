/*
LeetCode 443: https://leetcode.com/problems/string-compression/description/
*/

package leetcode

import "strconv"

func compress(chars []byte) int {
	prev := byte(0)
	result := 0
	count := 0
	for _, ch := range chars {
		if ch != prev {
			result += getLength443(chars, prev, count, result)
			count = 1
			prev = ch
		} else {
			count++
		}
	}

	result += getLength443(chars, prev, count, result)
	return result
}

func getLength443(chars []byte, prev byte, count int, result int) int {
	if count == 0 {
		return 0
	}

	chars[result] = prev
	if count == 1 {
		return 1
	}

	countStr := strconv.Itoa(count)
	for i, ch := range []byte(countStr) {
		chars[result+i+1] = ch
	}

	return len(countStr) + 1
}
