/*
LeetCode 186: https://leetcode.com/problems/reverse-words-in-a-string-ii/
*/

package leetcode

func reverseWords(s []byte) {
	reverse := func(s []byte, start, end int) {
		for start < end {
			s[start], s[end] = s[end], s[start]
			start, end = start+1, end-1
		}
	}

	i := 0
	for i < len(s) {
		if s[i] == ' ' {
			i++
		} else {
			j := i
			for j < len(s) && s[j] != ' ' {
				j++
			}

			reverse(s, i, j-1)
			i = j
		}
	}

	reverse(s, 0, len(s)-1)
}
