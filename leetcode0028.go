/*
LeetCode 28: https://leetcode.com/problems/implement-strstr/
*/

package leetcode

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if string([]byte(haystack)[i:i+len(needle)]) == needle {
			return i
		}
	}

	return -1
}
