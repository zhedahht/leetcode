/*
LeetCode 166: https://leetcode.com/problems/fraction-to-recurring-decimal/
*/

package leetcode

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	negative := (numerator > 0 && denominator < 0) || (numerator < 0 && denominator > 0)
	if numerator < 0 {
		numerator = -numerator
	}

	if denominator < 0 {
		denominator = -denominator
	}

	result := ""
	if numerator%denominator == 0 {
		result = strconv.Itoa(numerator / denominator)
	} else {
		result = strconv.Itoa(numerator/denominator) + "." + fractionToDecimalCore(numerator%denominator, denominator)
	}

	if negative {
		return "-" + result
	}

	return result
}

func fractionToDecimalCore(numerator int, denominator int) string {
	digits := make([]byte, 0)
	seenNumerators := make(map[int]int)
	index := 0
	prevIndex, seen := seenNumerators[numerator]
	for numerator != 0 && !seen {
		seenNumerators[numerator] = index
		index++

		val := numerator * 10
		digits = append(digits, byte(val/denominator+'0'))
		numerator = val % denominator

		prevIndex, seen = seenNumerators[numerator]
	}

	if numerator == 0 {
		return string(digits)
	}

	return string(digits[:prevIndex]) + "(" + string(digits[prevIndex:]) + ")"
}
