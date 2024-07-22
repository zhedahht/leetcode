/*
LeetCode 409: https://leetcode.com/problems/longest-palindrome/description/
*/

package leetcode

func longestPalindrome(s string) int {
	charCount := make(map[rune]int)
	for _, ch := range s {
		charCount[ch]++
	}

	hasOdd := false
	length := 0
	for _, count := range charCount {
		if count&1 == 1 {
			hasOdd = true
			length += count - 1
		} else {
			length += count
		}
	}

	if hasOdd {
		length++
	}

	return length
}
