/*
LeetCode 32: https://leetcode.com/problems/longest-valid-parentheses/
*/

package leetcode

import (
	"math"
)

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	result := 0
	for i, ch := range []byte(s) {
		if ch == ')' {
			var len int
			if i > 0 && s[i-1] == '(' {
				len = 2
			} else {
				if i > 0 {
					prev := i - dp[i-1] - 1
					if prev >= 0 && s[prev] == '(' {
						len = dp[i-1] + 2
					}
				}
			}

			if i-len >= 0 {
				len += dp[i-len]
			}

			dp[i] = len
			result = int(math.Max(float64(result), float64(len)))
		}
	}

	return result
}
