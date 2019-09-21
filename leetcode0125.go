/*
LeetCode 125: https://leetcode.com/problems/valid-palindrome/
*/

package leetcode

import "unicode"

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < len(s) && !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
		}

		for j >= 0 && !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j--
		}

		if i >= j {
			break
		}

		chI, chJ := unicode.ToUpper(rune(s[i])), unicode.ToUpper(rune(s[j]))
		if chI != chJ {
			return false
		}
	}

	return true
}
