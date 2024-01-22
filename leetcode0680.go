/*
LeetCode 680: https://leetcode.com/problems/valid-palindrome-ii/description/
*/

package leetcode

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			break
		}

		i, j = i+1, j-1
	}

	if i >= j {
		return true
	}

	return palindromeWithoutDeleting(s, i+1, j) || palindromeWithoutDeleting(s, i, j-1)
}

func palindromeWithoutDeleting(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			break
		}

		i, j = i+1, j-1
	}

	return i >= j
}
