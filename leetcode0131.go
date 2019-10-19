/*
LeetCode 131: https://leetcode.com/problems/palindrome-partitioning/
*/

package leetcode

// NOTE: The function name should be partition. Rename it to avoid conflicts.
func partition131(s string) [][]string {
	return helper131(s, 0, make([]string, 0), make([][]string, 0))
}

func helper131(s string, i int, segs []string, result [][]string) [][]string {
	isPalindrome := func(s string, i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}

			i, j = i+1, j-1
		}

		return true
	}

	if i == len(s) {
		clone := make([]string, len(segs))
		copy(clone, segs)
		result = append(result, clone)
		return result
	}

	for j := i; j < len(s); j++ {
		if isPalindrome(s, i, j) {
			seg := s[i : j+1]
			segs = append(segs, seg)
			result = helper131(s, j+1, segs, result)
			segs = segs[:len(segs)-1]
		}
	}

	return result
}
