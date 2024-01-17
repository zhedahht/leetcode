/*
LeetCode 647: https://leetcode.com/problems/palindromic-substrings/description/
*/

package leetcode

func countSubstrings(s string) int {
	count := 0
	for i := range s {
		start, end := i, i
		for start >= 0 && end < len(s) {
			if s[start] != s[end] {
				break
			}

			count++
			start--
			end++
		}

		start, end = i, i+1
		for start >= 0 && end < len(s) {
			if s[start] != s[end] {
				break
			}

			count++
			start--
			end++
		}
	}

	return count
}
