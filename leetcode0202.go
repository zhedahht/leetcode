/*
LeetCode 202: https://leetcode.com/problems/happy-number/
*/

package leetcode

import "strconv"

func isHappy(n int) bool {
	visited := make(map[int]bool)
	_, ok := visited[n]
	for n != 1 && !ok {
		visited[n] = true

		str := strconv.Itoa(n)
		n = 0
		for _, ch := range str {
			n += int((ch - '0') * (ch - '0'))
		}

		_, ok = visited[n]
	}

	return n == 1
}
