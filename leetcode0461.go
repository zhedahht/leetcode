/*
LeetCode 461: https://leetcode.com/problems/hamming-distance/description/
*/

package leetcode

func hammingDistance(x int, y int) int {
	xor := x ^ y
	result := 0
	for xor != 0 {
		result++
		xor = xor & (xor - 1)
	}

	return result
}
