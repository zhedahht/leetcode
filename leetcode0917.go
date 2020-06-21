/*
LeetCode 917: https://leetcode.com/problems/reverse-only-letters/
*/

package leetcode

import "unicode"

func reverseOnlyLetters(S string) string {
	chArray := []byte(S)
	i, j := 0, len(S)-1
	for i < j {
		if !unicode.IsLetter(rune(chArray[i])) {
			i++
		} else if !unicode.IsLetter(rune(chArray[j])) {
			j--
		} else {
			chArray[i], chArray[j] = chArray[j], chArray[i]
			i++
			j--
		}
	}

	return string(chArray)
}
