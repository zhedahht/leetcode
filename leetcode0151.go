/*
LeetCode 151: https://leetcode.com/problems/reverse-words-in-a-string/description/
*/

package leetcode

import "strings"

// The name should be reverseWords. It's renamed to avoid conflicts.
func reverseWords151(s string) string {
	reversed := make([]byte, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		reversed[len(s)-1-i] = s[i]
	}

	s = string(reversed)
	words := strings.Split(s, " ")
	result := make([]string, 0)
	for _, word := range words {
		if len(word) == 0 {
			continue
		}

		reversed = make([]byte, len(word))
		for i := len(word) - 1; i >= 0; i-- {
			reversed[len(word)-1-i] = word[i]
		}

		result = append(result, string(reversed))
	}

	return strings.Join(result, " ")
}
