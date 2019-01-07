/*
LeetCode 468: https://leetcode.com/problems/validate-ip-address/
*/

package leetcode

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if strings.Contains(IP, ".") && isIPv4(IP) {
		return "IPv4"
	}

	if isIPv6(strings.ToLower(IP)) {
		return "IPv6"
	}

	return "Neither"
}

func isIPv4(IP string) bool {
	segs := strings.Split(IP, ".")
	if len(segs) != 4 {
		return false
	}

	for _, seg := range segs {
		if len(seg) > 1 && seg[0] == '0' {
			return false
		}

		for _, ch := range seg {
			if ch < '0' || ch > '9' {
				return false
			}
		}

		val, err := strconv.Atoi(seg)
		if err != nil || val < 0 || val > 255 {
			return false
		}
	}

	return true
}

func isIPv6(IP string) bool {
	segs := strings.Split(IP, ":")
	if len(segs) != 8 {
		return false
	}

	for _, seg := range segs {
		if len(seg) > 4 || len(seg) == 0 {
			return false
		}

		for _, ch := range seg {
			if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f')) {
				return false
			}
		}
	}

	return true
}
