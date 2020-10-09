/*
LeetCode 541: https://leetcode.com/problems/reverse-string-ii/
*/

package leetcode

func reverseStr(s string, k int) string {
	min := func(m, n int) int {
		if m < n {
			return m
		}

		return n
	}

	result := make([]byte, len(s))
	for i := 0; i < len(s); i += 2 * k {
		j := i
		for l := min(len(s)-1, i+k-1); l >= i; l-- {
			result[j] = s[l]
			j++
		}

		for l := i + k; l < min(i+2*k, len(s)); l++ {
			result[j] = s[l]
			j++
		}
	}

	return string(result)
}
