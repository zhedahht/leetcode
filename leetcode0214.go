/*
LeetCode 214: https://leetcode.com/problems/shortest-palindrome/
*/

package leetcode

func shortestPalindrome(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	r := string(runes)
	i := len(s) - 1
	for ; i >= 0 && string(s[:i+1]) != string(r[len(s)-1-i:]); i-- {
	}

	return string(r[:len(s)-i-1]) + s
}
