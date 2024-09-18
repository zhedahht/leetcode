/*
LeetCode 405: https://leetcode.com/problems/convert-a-number-to-hexadecimal/description/
*/

package leetcode

import "strings"

func toHex(num int) string {
	bytes := make([]int, 0)
	mask := 0xf
	shift := 0
	for i := 0; i < 8; i++ {
		bt := (num & mask) >> shift
		bytes = append(bytes, bt)

		mask = mask << 4
		shift += 4
	}

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	i := 0
	for ; i < len(bytes); i++ {
		if bytes[i] > 0 {
			break
		}
	}

	if i == len(bytes) {
		i = len(bytes) - 1
	}

	bytes = bytes[i:]

	var sb strings.Builder
	for _, bt := range bytes {
		if bt < 10 {
			sb.WriteRune(rune(bt + '0'))
		} else {
			sb.WriteRune(rune(bt - 10 + 'a'))
		}
	}

	return sb.String()
}
