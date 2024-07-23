/*
LeetCode 482: https://leetcode.com/problems/license-key-formatting/description/
*/

package leetcode

import "strings"

func licenseKeyFormatting(s string, k int) string {
	chars := make([]byte, 0)
	bytes := []byte(s)
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if bytes[i] != '-' {
			chars = append(chars, bytes[i])
			count++
			if count == k {
				chars = append(chars, '-')
				count = 0
			}
		}
	}

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	if len(chars) > 0 && chars[0] == '-' {
		chars = chars[1:]
	}

	return strings.ToUpper(string(chars))
}
