/*
LeetCode 172: https://leetcode.com/problems/factorial-trailing-zeroes/
*/

package leetcode

func trailingZeroes(n int) int {
	result := 0
	for n > 0 {
		n = n / 5
		result += n
	}

	return result
}
