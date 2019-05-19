/*
LeetCode 29: https://leetcode.com/problems/divide-two-integers/
*/

package leetcode

import (
	"math"
)

func divide(dividend int, divisor int) int {
	isNegative := false
	if (dividend >= 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		isNegative = true
	}

	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	result := helper29(dividend, divisor)
	if isNegative {
		result = -result
	}

	return int(math.Min(float64(result), float64(0x7FFFFFFF)))
}

func helper29(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	}

	result, temp := 1, divisor
	for dividend >= temp*2 {
		result, temp = result<<1, temp<<1
	}

	return result + helper29(dividend-temp, divisor)
}
