/*
LeetCode 639: https://leetcode.com/problems/decode-ways-ii/description/
*/

package leetcode

import "strconv"

// The function name should be numDecodings. It's renamed to avoid conflicts.
func numDecodings639(s string) int {
	rs := []rune(s)
	oneDigit := numDecodingsOneDigit(rs[0])
	if len(s) == 1 {
		return oneDigit
	}

	twoDigits := oneDigit*numDecodingsOneDigit(rs[1]) + numDecodingsTwoDigits(rs[:2])
	if len(s) == 2 {
		return twoDigits
	}

	dp := [2]int{oneDigit, twoDigits}
	for i := 2; i < len(rs); i++ {
		oneDigit = dp[(i-1)%2] * numDecodingsOneDigit(rs[i])
		twoDigits = dp[(i-2)%2] * numDecodingsTwoDigits([]rune{rs[i-1], rs[i]})
		dp[i%2] = (oneDigit + twoDigits) % 1000000007
	}

	return dp[(len(rs)-1)%2]
}

func numDecodingsOneDigit(r rune) int {
	switch r {
	case '0':
		return 0
	case '*':
		return 9
	default:
		return 1
	}
}

func numDecodingsTwoDigits(r []rune) int {
	if r[0] == '*' && r[1] == '*' {
		// 11-26, excluding 20
		return 15
	}

	if r[0] == '*' {
		count := 0
		for _, n := range []rune{'1', '2'} {
			count += numDecodingsTwoDigits([]rune{n, r[1]})
		}

		return count
	}

	if r[1] == '*' {
		count := 0
		for _, n := range []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'} {
			count += numDecodingsTwoDigits([]rune{r[0], n})
		}

		return count
	}

	val, _ := strconv.Atoi(string(r))
	//fmt.Println(val)
	if val >= 10 && val <= 26 {
		return 1
	}

	return 0
}
