/*
LeetCode 917: https://leetcode.com/problems/reverse-only-letters/
*/

package leetcode

func reverseOnlyLetters(S string) string {
	chArray := []byte(S)
	i, j := 0, len(S)-1
	for i < j {
		if !isLetter(chArray[i]) {
			i++
		} else if !isLetter(chArray[j]) {
			j--
		} else {
			chArray[i], chArray[j] = chArray[j], chArray[i]
			i++
			j--
		}
	}

	return string(chArray)
}

func isLetter(ch byte) bool {
	if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
		return true
	}

	return false
}
