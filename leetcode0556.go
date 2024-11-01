/*
LeetCode 556: https://leetcode.com/problems/next-greater-element-iii/description/
*/

package leetcode

import (
	"fmt"
	"math"
	"sort"
)

// The name should be nextGreaterElement. It's renamed to avoid conflicts.
func nextGreaterElement556(n int) int {
	nStr := fmt.Sprintf("%d", n)
	digits := []byte(nStr)
	i := len(digits) - 2
	for ; i >= 0; i-- {
		if digits[i] < digits[i+1] {
			break
		}
	}

	if i < 0 {
		return -1
	}

	digit := digits[i]
	minDigit := byte('9')
	index := 0
	for j := i + 1; j < len(digits); j++ {
		if digits[j] > digit && digits[j] <= minDigit {
			minDigit = digits[j]
			index = j
		}
	}

	digits[i], digits[index] = digits[index], digits[i]
	suffix := digits[i+1:]
	sort.Slice(suffix, func(i, j int) bool {
		return suffix[i] < suffix[j]
	})

	nextDigits := digits[:i+1]
	nextDigits = append(nextDigits, suffix...)

	result := int64(0)
	for _, d := range nextDigits {
		result = result*10 + int64(d-'0')
	}

	if result > math.MaxInt32 {
		return -1
	}

	return int(result)
}
