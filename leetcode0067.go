/*
LeetCode 67: https://leetcode.com/problems/add-binary/
*/

package leetcode

import (
	"math"
)

func addBinary(a string, b string) string {
	length := int(math.Max(float64(len(a)), float64(len(b)))) + 1
	result := make([]byte, length)
	i, j, k := len(a)-1, len(b)-1, length-1
	carry := 0
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		if sum > 1 {
			sum = sum - 2
			carry = 1
		} else {
			carry = 0
		}

		result[k] = byte(sum + '0')
		k--
	}

	if carry > 0 {
		result[0] = '1'
	}

	if result[0] != '1' {
		result = result[1:len(result)]
	}

	return string(result)
}
