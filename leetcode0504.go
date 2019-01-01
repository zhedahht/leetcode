/*
LeetCode 504: https://leetcode.com/problems/base-7/
*/

package leetcode

import "fmt"

func convertToBase7(num int) string {
	result := ""
	negative := false
	if num < 0 {
		negative = true
		num = -num
	}

	for num > 0 {
		result = fmt.Sprintf("%d%s", num%7, result)
		num = num / 7
	}

	if len(result) == 0 {
		result = "0"
	}

	if negative {
		result = "-" + result
	}

	return result
}
