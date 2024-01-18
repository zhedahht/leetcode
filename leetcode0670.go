/*
LeetCode 670: https://leetcode.com/problems/maximum-swap/description/
*/

package leetcode

import "sort"

func maximumSwap(num int) int {
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}

	if len(digits) == 0 {
		digits = append(digits, 0)
	}

	for i, j := 0, len(digits)-1; i < j; {
		digits[i], digits[j] = digits[j], digits[i]
		i++
		j--
	}

	sortedDigits := make([]int, len(digits))
	copy(sortedDigits, digits)
	sort.Slice(sortedDigits, func(i, j int) bool {
		return sortedDigits[i] > sortedDigits[j]
	})

	i := 0
	for ; i < len(digits); i++ {
		if digits[i] != sortedDigits[i] {
			break
		}
	}

	for j := len(digits) - 1; j > i; j-- {
		if digits[j] == sortedDigits[i] {
			digits[i], digits[j] = digits[j], digits[i]
		}
	}

	result := 0
	for i := 0; i < len(digits); i++ {
		result = result*10 + digits[i]
	}

	return result
}
