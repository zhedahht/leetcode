/*
LeetCode 95: https://leetcode.com/problems/valid-number/
*/

package leetcode

import "strings"

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	segs := strings.Split(s, "e")
	if len(segs) > 2 {
		return false
	}

	num := removeSign(segs[0])
	if len(num) == 0 {
		return false
	}

	if len(segs) == 2 {
		exponent := removeSign(segs[1])
		if len(exponent) == 0 || !isUnsignedDigits(exponent) {
			return false
		}
	}

	segs = strings.Split(num, ".")
	if len(segs) > 2 {
		return false
	}

	integer := segs[0]
	if len(integer) > 0 && !isUnsignedDigits(integer) {
		return false
	}

	if len(segs) == 2 {
		float := segs[1]
		if len(integer) == 0 && len(float) == 0 {
			return false
		}

		if len(float) > 0 && !isUnsignedDigits(float) {
			return false
		}
	}

	return true
}

func removeSign(s string) string {
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		s = s[1:]
	}

	return s
}

func isUnsignedDigits(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	return true
}
