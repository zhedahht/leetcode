/*
LeetCode 891: https://leetcode.com/problems/sum-of-subsequence-widths/
*/

package leetcode

import "sort"

func sumSubseqWidths(A []int) int {
	const modBase = 1000000007
	pow2 := make([]int, len(A)+1)
	pow2[0] = int(1)
	for i := 1; i < len(pow2); i++ {
		pow2[i] = (pow2[i-1] << 1) % modBase
	}

	result := 0
	sort.Ints(A)
	for i, num := range A {
		result += num * (pow2[i] - 1)
		result -= num * (pow2[len(A)-i-1] - 1)
		result %= modBase
	}

	return result
}
