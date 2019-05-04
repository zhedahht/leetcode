/*
LeetCode 91: https://leetcode.com/problems/decode-ways/
*/

package leetcode

import (
	"strconv"
)

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	var dp [2]int
	dp[0] = 1
	if isOneDigitCode(s[0]) {
		dp[1] = 1
	} else {
		dp[1] = 0
	}

	bytes := []byte(s)
	for i := 1; i < len(s); i++ {
		count := 0
		if isOneDigitCode(bytes[i]) {
			count += dp[i%2]
		}

		if isTwoDigitCode(bytes[i-1 : i+1]) {
			count += dp[(i-1)%2]
		}

		dp[(i+1)%2] = count
	}

	return dp[len(s)%2]
}

func isOneDigitCode(b byte) bool {
	return b >= '1' && b <= '9'
}

func isTwoDigitCode(b []byte) bool {
	val, err := strconv.Atoi(string(b))
	return err == nil && val >= 10 && val <= 26
}
