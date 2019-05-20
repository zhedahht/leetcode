/*
LeetCode 93: https://leetcode.com/problems/restore-ip-addresses/
*/

package leetcode

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	return helper93(s, 0, "", make([]string, 0), make([]string, 0))
}

func helper93(s string, index int, num string, ip, result []string) []string {
	if index == len(s) && len(num) > 0 && len(ip) == 3 {
		ip = append(ip, num)
		result = append(result, fmt.Sprintf("%s.%s.%s.%s", ip[0], ip[1], ip[2], ip[3]))
	} else if index < len(s) && len(ip) <= 4 {
		if num == "" || num[0] != '0' {
			newNum := fmt.Sprintf("%s%c", num, s[index])
			if val, err := strconv.Atoi(newNum); val <= 255 && err == nil {
				result = helper93(s, index+1, newNum, ip, result)
			}
		}

		if len(ip) < 4 && len(num) > 0 {
			ip = append(ip, num)
			num = fmt.Sprintf("%c", s[index])
			result = helper93(s, index+1, num, ip, result)
		}
	}

	return result
}
